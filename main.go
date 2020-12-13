package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

func main() {

	var name = flag.String("name", "", "Give a name please")
	flag.Parse()

	//create an etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer cli.Close()

	//create an etcd session
	s, _ = concurrency.NewSession(cli)

	defer s.Close()

	lck := concurrency.NewMutex(s, "/distributed-lock/")

	ctx := context.Background()

	//locking
	if err := lck.Lock(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locked for ", *name)
	fmt.Println("Doing some work for", *name)
	time.Sleep(5 * time.Second)

	//unlocking
	if err := lck.Unlock(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unlocked for ", *name)
}
