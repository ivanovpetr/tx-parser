package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ivanovpetr/tx-parser/handler"
	"github.com/ivanovpetr/tx-parser/internal/config"
	"github.com/ivanovpetr/tx-parser/pkg/ethclient"
	"github.com/ivanovpetr/tx-parser/pkg/signalobserver"
	"github.com/ivanovpetr/tx-parser/pkg/storage"
	"github.com/ivanovpetr/tx-parser/service/parser"
)

func main() {
	var configPath = flag.String("c", "", "specify config path")
	flag.Parse()
	if *configPath == "" {
		log.Fatal("specify config path using -c")
	}
	conf, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to parse config file: %w", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	store := storage.New(conf.Parser.StartingBlock)
	client := ethclient.New(conf.Ethereum.RPCUrl)

	p := parser.New(client, store, parser.WithLookupInterval(conf.Parser.LookupInterval))

	h := handler.New(store)
	mux := http.NewServeMux()
	mux.HandleFunc("/currentBlock", h.GetCurrentBlock)
	mux.HandleFunc("/transactions/", h.Transactions)
	mux.HandleFunc("/subscribe", h.Subscribe)

	srv := &http.Server{
		Addr:    ":" + conf.Http.Port,
		Handler: mux,
	}
	go func() {
		fmt.Printf("Starting http server. Listen on port :%s \n", conf.Http.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go signalobserver.ObserveSignal(cancel)

	// lock on parser
	fmt.Printf("Starting parser after block %d", conf.Parser.StartingBlock)
	p.Run(ctx)

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	fmt.Println("Shutting down http server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
