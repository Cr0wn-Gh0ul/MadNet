package lstate

import (
	"sync"

	"github.com/MadBase/MadNet/consensus/appmock"
	"github.com/MadBase/MadNet/interfaces"
)

type txCache struct {
	sync.Mutex
	app    appmock.Application
	cache  map[string]string
	rcache map[string]uint32
}

func (txc *txCache) Init() error {
	txc.cache = make(map[string]string)
	txc.rcache = make(map[string]uint32)
	return nil
}

func (txc *txCache) Add(height uint32, tx interfaces.Transaction) error {
	txc.Lock()
	defer txc.Unlock()
	txHash, err := tx.TxHash()
	if err != nil {
		return err
	}
	txb, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	txc.rcache[string(txHash)] = height
	txc.cache[string(txHash)] = string(txb)
	return nil
}

func (txc *txCache) Contains(txHsh []byte) bool {
	txc.Lock()
	defer txc.Unlock()
	if _, ok := txc.getInternal(txHsh); ok {
		return true
	}
	return false
}

func (txc *txCache) Get(txHsh []byte) (interfaces.Transaction, bool) {
	txc.Lock()
	defer txc.Unlock()
	return txc.getInternal(txHsh)
}

func (txc *txCache) getInternal(txHsh []byte) (interfaces.Transaction, bool) {
	txIf, ok := txc.cache[string(txHsh)]
	if ok {
		txb := txIf
		tx, err := txc.app.UnmarshalTx([]byte(txb))
		if err != nil {
			txc.delInternal(txHsh)
			return nil, false
		}
		return tx, ok
	}
	return nil, false
}

func (txc *txCache) GetHeight(height uint32) ([]interfaces.Transaction, [][]byte) {
	txc.Lock()
	defer txc.Unlock()
	out := []interfaces.Transaction{}
	outhsh := [][]byte{}
	for hash, rh := range txc.rcache {
		hash, rh := hash, rh
		if height == rh {
			if txi, ok := txc.getInternal([]byte(hash)); ok {
				out = append(out, txi)
				outhsh = append(outhsh, []byte(hash))
			}
		}
	}
	return out, outhsh
}

/*

func (txc *txCache) GetMany(txHashes [][]byte) ([]interfaces.Transaction, [][]byte) {
	txc.RLock()
	defer txc.RUnlock()
	result := []interfaces.Transaction{}
	missing := [][]byte{}
	for i := 0; i < len(txHashes); i++ {
		txIf, ok := txc.cache[string(txHashes[i])]
		if !ok {
			missing = append(missing, utils.CopySlice(txHashes[i]))
			continue
		}
		txb := txIf
		tx, err := txc.app.UnmarshalTx([]byte(txb))
		if err != nil {
			missing = append(missing, utils.CopySlice(txHashes[i]))
			delete(txc.cache, string(txHashes[i]))
			continue
		}
		result = append(result, tx)
	}
	return result, missing
}
*/

func (txc *txCache) Del(txHsh []byte) {
	txc.Lock()
	defer txc.Unlock()
	txc.delInternal(txHsh)
}

func (txc *txCache) delInternal(txHsh []byte) {
	delete(txc.cache, string(txHsh))
	delete(txc.rcache, string(txHsh))
}

func (txc *txCache) DropBeforeHeight(dropHeight uint32) []string {
	out := []string{}
	if dropHeight-256 > dropHeight {
		return out
	}
	txc.Lock()
	defer txc.Unlock()
	for hash, height := range txc.rcache {
		hash, height := hash, height
		if height <= uint32(dropHeight) {
			txc.delInternal([]byte(hash))
			out = append(out, hash)
		}
	}
	return out
}
