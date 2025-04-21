package main

import (
	"client/controller"
	"client/infrastructure"
	"client/service"
	"sync"
)

func init() {
	infrastructure.InitDb()
}

func main() {
	var wg sync.WaitGroup
	// fetch metrics periodic
	wg.Add(2)
	go service.PeriodicFetchMetrics(&wg) // pass waitgroup
	// both handler and periodic fetch are running in concurrent manner
	go controller.Handler(&wg)
	wg.Wait()
}
