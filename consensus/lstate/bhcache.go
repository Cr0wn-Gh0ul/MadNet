package lstate

import (
	"strconv"
	"sync"

	"github.com/MadBase/MadNet/consensus/objs"
	"github.com/MadBase/MadNet/utils"
)

type bHCache struct {
	sync.Mutex
	cache map[uint32]string
}

func (bhc *bHCache) Init() error {
	bhc.cache = make(map[uint32]string)
	return nil
}

func (bhc *bHCache) Add(bh *objs.BlockHeader) error {
	bhc.Lock()
	defer bhc.Unlock()
	bhBytes, err := bh.MarshalBinary()
	if err != nil {
		return err
	}
	bhc.cache[bh.BClaims.Height] = string(bhBytes)
	return nil
}

func (bhc *bHCache) Contains(height uint32) bool {
	bhc.Lock()
	defer bhc.Unlock()
	if _, ok := bhc.getInternal(height); ok {
		return true
	}
	return false
}

func (bhc *bHCache) Get(height uint32) (*objs.BlockHeader, bool) {
	bhc.Lock()
	defer bhc.Unlock()
	return bhc.getInternal(height)
}

// refactor to just use height
func (bhc *bHCache) getInternal(height uint32) (*objs.BlockHeader, bool) {
	bhIf, ok := bhc.cache[height]
	if ok {
		bhString := bhIf
		bhBytes := []byte(bhString)
		bhCopy := utils.CopySlice(bhBytes)
		bh := &objs.BlockHeader{}
		err := bh.UnmarshalBinary(bhCopy)
		if err != nil {
			bhc.delInternal(height)
			return nil, false
		}
		return bh, true
	}
	return nil, false
}

func (bhc *bHCache) Del(height uint32) {
	bhc.Lock()
	defer bhc.Unlock()
	bhc.delInternal(height)
}

func (bhc *bHCache) delInternal(height uint32) {
	delete(bhc.cache, height)
}

func (txc *bHCache) DropBeforeHeight(dropHeight uint32) []string {
	out := []string{}
	if dropHeight-256 > dropHeight {
		return out
	}
	txc.Lock()
	defer txc.Unlock()
	for height := range txc.cache {
		height := height
		if height <= uint32(dropHeight) {
			out = append(out, strconv.Itoa(int(height)))
			txc.delInternal(height)
		}
	}
	return out
}
