package main

import (
	"receipt-processor-challenge/server"
)

func main() {
	srv := server.New()
	srv.ServeHTTP()
}
