package storage

import (
	"sync"
	"sync/atomic"

	"github.com/ivanovpetr/tx-parser/pkg/ethclient"
	"github.com/ivanovpetr/tx-parser/pkg/ethutil"
)

// Storage storage implementation for Parser
type Storage struct {
	mu              *sync.RWMutex
	LastParsedBlock int64
	Subscribers     map[ethutil.Address][]*ethclient.Transaction
}

// New creates new parser
func New(lastParsedBlock int64) *Storage {
	return &Storage{
		mu:              &sync.RWMutex{},
		LastParsedBlock: lastParsedBlock,
		Subscribers:     make(map[ethutil.Address][]*ethclient.Transaction),
	}
}

// AddSubscriber adds subscriber to the storage
func (s *Storage) AddSubscriber(subscriber ethutil.Address) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Subscribers[subscriber]; !ok {
		s.Subscribers[subscriber] = []*ethclient.Transaction{}
	}
	return nil
}

// RemoveSubscriber removes subscriber from the storage
func (s *Storage) RemoveSubscriber(subscriber ethutil.Address) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Subscribers, subscriber)
	return nil
}

// GetAllTransactionsBySubscriber returns list of transactions for the selected address
func (s *Storage) GetAllTransactionsBySubscriber(subscriber ethutil.Address) ([]*ethclient.Transaction, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.Subscribers[subscriber]; !ok {
		return nil, NotFoundError
	}

	return s.Subscribers[subscriber], nil
}

// AddBlock adds transactions for subscribers and increment block number
func (s *Storage) AddBlock(txs []ethclient.Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, tx := range txs {
		tx := tx
		if _, ok := s.Subscribers[tx.From.Value]; ok {
			s.Subscribers[tx.From.Value] = append(s.Subscribers[tx.From.Value], &tx)
		}
		if _, ok := s.Subscribers[tx.To.Value]; ok {
			s.Subscribers[tx.To.Value] = append(s.Subscribers[tx.To.Value], &tx)
		}
	}
	s.LastParsedBlock++

	return nil
}

// GetLastParsedBlock returns last parsed block
func (s *Storage) GetLastParsedBlock() (int64, error) {
	return atomic.LoadInt64(&s.LastParsedBlock), nil
}
