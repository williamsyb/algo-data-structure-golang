package main

import (
	"flag"
	mgrpc "github.com/csunny/argo/examples/grpc/grpc_example"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := mgrpc.Proxy(); err != nil {
		glog.Fatal(err)
	}
}
