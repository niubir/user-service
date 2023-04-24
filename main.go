package main

import (
	"sync"

	"github.com/niubir/user-service/config"
	"github.com/niubir/user-service/grpc"
	"github.com/niubir/user-service/http"
	"github.com/niubir/user-service/logs"
	"github.com/niubir/user-service/services"
)

func main() {
	Init()

	var wg sync.WaitGroup

	wg.Add(1)
	go http.Init(&wg)

	wg.Add(1)
	go grpc.Init(&wg)

	wg.Wait()
}

func Init() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := logs.Init(); err != nil {
		panic(err)
	}
	if err := services.Init(); err != nil {
		panic(err)
	}
}
