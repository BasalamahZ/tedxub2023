package main

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/tedxub2023/cmd/tedxub2023-api-http/server"
)

func main() {
	godotenv.Load()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		os.Exit(server.Run())
		defer wg.Done()
	}()
	wg.Wait()
}
