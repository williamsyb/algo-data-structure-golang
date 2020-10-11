package wantlist

import (
	"sort"
	"sync"
	cid "github.com/ipfs/go-cid"
)

type ThreadSafe struct{
	lk sync.RWMutex
	set map[string]*Entry
}

type Wantlist struct{
	set map[string]*Entry
}

type Entry  struct{
	Cid  *cid.Cid
	Priority int
	SesTrk map[uint64]struct{}
}

// NewRefEntry creates a new reference tracked wantlist entry
func NewRefEntry(c *cid.Cid, p int) *Entry{
	return &Entry{
		Cid: c,
		Priority: p,
		SesTrk: make(map[uint64]struct{}),
	}
}

type entrySlice []*Entry

func (es entrySlice) Len() int   {return len(es)}
func (es entrySlice) Swap(i, j int) {es[i], es[j] = es[j], es[i]}
func (es entrySlice) Less(i, j int) bool {return es[i].Priority > es[j].Priority}

func NewThreadSafe() *ThreadSafe{
	return &ThreadSafe{
		set: make(map[string]*Entry),
	}
}

func New() *Wantlist{
	return &Wantlist{
		set: make(map[string]*Entry),
	}
}

// Add adds the given cid to the wantlist with the specified priority,governed by the session 
// ID 'ses'. if a cid is added under multiple session IDs, then it must be removed by each of 
// those sessions before it is no longer in the wantlist calls to add are indmpotent given the
// same arguments. Subsequent 
// add return true if the cid did not in the wantlist before this call even if it was under a 
// different session

func (w *ThreadSafe) Add(c *cid.Cid, priority int, ses uint64) bool{
	w.lk.Lock()
	defer w.lk.Unlock()

	k := c.KeyString()
	if e, ok := w.set[k]; ok{
		e.SesTrk[ses] = struct{}{}
		return false
	}

	w.set[k] = &Entry{
		Cid: c,
		Priority: priority,
		SesTrk: map[uint64]struct{}{ses: struct{}{}},
	}
	return true
}

func (w *ThreadSafe) AddEntry(e *Entry, ses uint64) bool{
	w.lk.Lock()
	defer w.lk.Unlock()
	k := e.Cid.KeyString()
	if ex, ok := w.set[k]; ok{
		ex.SesTrk[ses] = struct{}{}
		return false
	}
	w.set[k] = e
	e.SesTrk[ses] = struct{}{}
	return true
}

func (w *ThreadSafe) Remove(c *cid.Cid, ses uint64) bool{
	w.lk.Lock()
	defer w.lk.Unlock()

	k := c.KeyString()
	e, ok := w.set[k]
	if !ok{
		return false
	}

	delete(e.SesTrk, ses)
	if len(e.SesTrk) == 0{
		delete(w.set, k)
		return true
	}
	return false
}

func (w *ThreadSafe) Contains(k *cid.Cid) (*Entry, bool){
	w.lk.RLock()
	defer w.lk.RUnlock()
	e, ok := w.set[k.KeyString()]
	return e, ok
}

func (w *ThreadSafe) Entries() []*Entry{
	w.lk.RLock()
	defer w.lk.RUnlock()

	es := make([]*Entry, 0, len(w.set))
	for _, e := range w.set{
		es = append(es, e)
	}
	return es
}

func (w *ThreadSafe) Len()int{
	w.lk.RLock()
	defer w.lk.RUnlock()
	return len(w.set)
}

func (w *Wantlist) Add(c *cid.Cid, priority int) bool{
	k := c.KeyString()
	if _, ok := w.set[k]; ok{
		return false
	}

	w.set[k] = &Entry{
		Cid: c,
		Priority: priority,
	}
	return true
}

func (w *Wantlist) AddEntry (e *Entry) bool{
	k := e.Cid.KeyString()
	if _, ok := w.set[k]; ok{
		return false
	}
	w.set[k] = e
	return true
}

func (w *Wantlist) Remove(c *cid.Cid) bool{
	k := c.KeyString()
	_, ok := w.set[k]
	if !ok {
		return false
	}
	delete(w.set, k)
	return true
}

func (w *Wantlist) Contains(k *cid.Cid) (*Entry, bool){
	e, ok := w.set[k.KeyString()]
	return e, ok
}

func (w *Wantlist) Entries() []*Entry{
	es := make([]*Entry, 0, len(w.set))
	for _, e := range w.set{
		es = append(es, e)
	}

	return es
}

func (w *Wantlist) SortedEntries() []*Entry{
	es := w.Entries()
	sort.Sort(entrySlice(es))
	return es
}