package clients

import (
	"github.com/brenordv/go-monitor-internet-connection/internal/utils"
	"github.com/dgraph-io/badger/v3"
	"time"
)

func Persist(key []byte, content []byte, ttl int) error {
	var dbDir string
	var err error
	var db *badger.DB

	dbDir, err = utils.GetDbDir()
	if err != nil {
		return err
	}

	opts := badger.DefaultOptions(dbDir)
	opts.CompactL0OnClose = true
	opts.Logger = nil


	db, err = badger.Open(opts)
	if err != nil {
		return err
	}

	defer func(db *badger.DB) {
		_ = db.Close()
	}(db)

	err = db.Update(func(txn *badger.Txn) error {
		if ttl > 0 {
			entry := badger.NewEntry(key, content).WithTTL(time.Duration(ttl) * time.Hour)
			return txn.SetEntry(entry)
		}
		return txn.Set(key, content)
	})

	return err
}
