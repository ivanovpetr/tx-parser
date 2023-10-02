package parser

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ivanovpetr/tx-parser/pkg/ethclient"
	"github.com/ivanovpetr/tx-parser/pkg/ethutil"
)

// Storage interface for Parser
type Storage interface {
	AddSubscriber(subscriber ethutil.Address) error
	GetAllTransactionsBySubscriber(subscriber ethutil.Address) ([]*ethclient.Transaction, error)
	GetLastParsedBlock() (int64, error)
	AddBlock([]ethclient.Transaction) error
}

// Parser parses ethereum transactions and stores transactions matched by subscribers
type Parser struct {
	lookupInterval int64
	client         ethclient.EthClient
	storage        Storage
	lookupLock     chan struct{}
}

// New creates new parser
func New(client ethclient.EthClient, storage Storage, options ...Option) *Parser {
	parser := &Parser{client: client, storage: storage, lookupLock: make(chan struct{}, 1)}
	for _, apply := range options {
		apply(parser)
	}
	return parser
}

type Option func(parser *Parser)

// WithLookupInterval sets lookupInterval for Parser
func WithLookupInterval(interval int64) Option {
	return func(parser *Parser) {
		parser.lookupInterval = interval
	}
}

// Run runs Parser
func (p *Parser) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(p.lookupInterval))
	wg := &sync.WaitGroup{}
	errChan := make(chan error)
	go p.logger(errChan)
	wg.Add(1)
	go func() {
	Loop:
		for {
			select {
			case <-ticker.C:
				wg.Add(1)
				go p.lookup(ctx, wg, errChan)

			case <-ctx.Done():
				fmt.Println("Shutting down parser")
				break Loop
			}
		}
		wg.Done()
	}()
	wg.Wait()
	close(errChan)
}

// lookup requests last mined from ethereum and queries all absent blocks
func (p *Parser) lookup(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	select {
	case p.lookupLock <- struct{}{}:
	default:
		return
	}
	defer func() { <-p.lookupLock }()
	lastParsedBlock, err := p.storage.GetLastParsedBlock()
	if err != nil {
		return
	}

	lastMinedBlock, err := p.client.GetLatestBlockNumber(ctx)
	if lastMinedBlock > lastParsedBlock {
		for i := 1; int64(i) <= lastMinedBlock-lastParsedBlock; i++ {
			txs, err := p.client.GetTransactionsByBlockNumber(ctx, int64(i)+lastParsedBlock)
			if err != nil {
				errChan <- err
				return
			}
			err = p.storage.AddBlock(txs)
			if err != nil {
				errChan <- err
				return
			}
		}
	}
}

// logger logs Parser errors
func (p Parser) logger(errChan <-chan error) {
	for err := range errChan {
		log.Printf("parsing error occured: %w", err)
	}
}
