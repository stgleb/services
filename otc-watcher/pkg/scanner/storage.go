package scanner

import (
	"sync"
	"time"

	"github.com/skycoin/services/otc/pkg/otc"
)

type Relevant struct {
	sync.RWMutex

	Outputs otc.Outputs `json:"outputs"`
}

type Updated struct {
	Time   int64  `json:"time"`
	Height uint64 `json:"height"`
}

type Storage struct {
	sync.RWMutex

	Filename     string               `json:"filename"`
	Updated      *Updated             `json:"updated"`
	Addresses    map[string]*Relevant `json:"addresses"`
	Transactions map[string]*Relevant `json:"transaction"`
}

func NewStorage(cur otc.Currency) *Storage {
	return &Storage{
		Filename:  string(cur) + ".json",
		Updated:   &Updated{},
		Addresses: make(map[string]*Relevant, 0),
	}
}

func (s *Storage) Register(addr string) {
	s.Lock()
	defer s.Unlock()

	s.Addresses[addr] = &Relevant{Outputs: make(otc.Outputs, 0)}
}

func (s *Storage) Outputs(addr string) otc.Outputs {
	s.RLock()
	defer s.RUnlock()

	// TODO: is this needed?
	s.Addresses[addr].RLock()
	defer s.Addresses[addr].RUnlock()

	return s.Addresses[addr].Outputs
}

func (s *Storage) Update(block *otc.Block) {
	s.Lock()
	defer s.Unlock()

	// iterate all transactions in block
	for hash, tx := range block.Transactions {
		// Iterate over all inputs to keep track if transaction was spent or unspent.
		for _, in := range tx.In {
			for _, addr := range s.Addresses {
				addr.Outputs.UpdateSpent(in.From, in.To)
			}
		}

		// iterate all outputs in transaction
		for index, out := range tx.Out {
			for _, outAddr := range out.Addresses {
				// iterate all registered addresses
				for addr, rel := range s.Addresses {
					// if registered address is in output addresses
					if addr == outAddr {
						rel.Lock()
						rel.Outputs.Update(
							hash,
							index,
							&otc.OutputVerbose{
								Amount:        out.Amount,
								TxHash:        tx.Hash,
								BlockHash:     tx.BlockHash,
								Confirmations: tx.Confirmations,
								Height:        block.Height,
							},
						)
						rel.Unlock()
					}
				}
			}
		}
	}

	// update confirmations number
	for _, rel := range s.Addresses {
		for _, outputs := range rel.Outputs {
			for _, out := range outputs {
				out.Confirmations = (block.Height - out.Height) + 1
			}
		}
	}

	// record that everything was updated to current block and time
	s.Updated.Height = block.Height
	s.Updated.Time = time.Now().UTC().Unix()
}
