package pos

import (
	"sync"
	"encoding/hex"
	"crypto/sha256"
	"time"
	"net"
	"io"
	"bufio"
	"strconv"
	"log"
	"fmt"
	"encoding/json"
	"math/rand"
	"github.com/davecgh/go-spew/spew"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
	Validator string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block
var tempBlocks []Block

// CandidateBlocks handles incoming blocks for validation
var candidateBlocks = make(chan Block)

// announcements broadcasts winning validator to all nodes
var announcements = make(chan string)

var mutex = &sync.Mutex{}

// validators keeps track of open validator and balances
var validators = make(map[string]int)

// sha256 hashing

func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func calculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash

	return calculateHash(record)
}

func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.BPM = BPM
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()

	// validator address
	var address string

	// allow user to allocate number of tokens to stake
	// the greater the number of tokens, the greater chance to forging a new block
	io.WriteString(conn, "Enter token balance:")
	scanBalance := bufio.NewScanner(conn)

	// 根据读取到的balance数量，来构建validators的数量
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
		}

		t := time.Now()
		address = calculateHash(t.String())
		validators[address] = balance

		fmt.Println(validators)
		break
	}

	io.WriteString(conn, "\nEnter a new BPM:")
	scanBPM := bufio.NewScanner(conn)

	go func() {
		for {
			// take in BPM from stdin and add it to blockchain after conducting necessary validation
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())

				// if malicious party
				if err != nil {
					log.Printf("%v not a number: %v", scanBPM.Text(), err)
					delete(validators, address)
					conn.Close()
				}

				mutex.Lock()
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()

				// create newBlock for consideration to be forged
				newBlock, err := generateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}

				spew.Dump(newBlock, oldLastIndex)

				if isBlockValid(newBlock, oldLastIndex) {
					candidateBlocks <- newBlock
				}

				io.WriteString(conn, "\nEnter a new BPM:")
			}

		}
	}()

	// simulate receiving broadcast
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		output, err := json.Marshal(Blockchain)

		mutex.Unlock()
		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(conn, string(output)+"\n")
	}
}

// pickWinner creates a lottery pool of validators and chooses the validator who gets to forge a block
// to the blockchain
// by random selecting from the pool, weighted by amount of tokens staked
func PickWinner() {
	time.Sleep(10 * time.Second)
	mutex.Lock()

	temp := tempBlocks
	mutex.Unlock()

	lotteryPool := []string{}

	if len(temp) > 0 {
		// slightly modified traditional proof of stake algorithm
		// from all validators who submitted a block, weight them by the
		// number of staked tokens
		// in traditional proof of stake, validators can participate without submitting a block to be forged.
	OUTER:
		for _, block := range temp {
			// if already in lottery pool, skip
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			// lock list of validators to prevent data race
			mutex.Lock()
			setValidators := validators
			mutex.Unlock()

			k, ok := setValidators[block.Validator]
			if ok {
				for i := 0; i < k; i++ {
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}
		}

		// randomly pick winner from lottery pool
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		// add block of winner to blockchain and let all the other nodes known
		for _, block := range temp {
			if block.Validator == lotteryWinner {
				mutex.Lock()
				Blockchain = append(Blockchain, block)
				mutex.Unlock()

				for _ = range validators {
					announcements <- "\nwinning validator: " + lotteryWinner + "\n"
				}
				break
			}
		}
	}

	mutex.Lock()
	tempBlocks = []Block{}
	mutex.Unlock()
}

//
func Run() {

	// create a genesis block
	t := time.Now()
	genesisBlock := Block{}

	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)

	Blockchain = append(Blockchain, genesisBlock)

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			PickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}
