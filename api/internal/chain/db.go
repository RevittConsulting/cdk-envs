package chain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RevittConsulting/cdk-envs/internal/jsonrpc"
	"github.com/boltdb/bolt"
)

type Db struct {
	db *bolt.DB
}

func NewDb(db *bolt.DB) *Db {
	return &Db{
		db: db,
	}
}

func (b *Db) getHighestBlock(ctx context.Context) (*jsonrpc.Block, error) {
	var highestBlock *jsonrpc.Block
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Blocks"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		highestBlockNumber := int64(-1)
		err := bucket.ForEach(func(k, v []byte) error {
			var block *jsonrpc.Block
			if err := json.Unmarshal(v, &block); err != nil {
				return err
			}

			if block.Number > highestBlockNumber {
				highestBlockNumber = block.Number
				highestBlock = block
			}

			return nil
		})

		return err
	})

	if err != nil {
		return nil, err
	}

	if highestBlock == nil {
		return nil, fmt.Errorf("no blocks found")
	}

	return highestBlock, nil
}
