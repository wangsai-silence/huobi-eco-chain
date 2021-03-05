package core

import (
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

const (
	normal types.TxPoolType = 0
)

type poolTypeItem struct {
	PoolType types.TxPoolType `json:"poolType"`
	Percent  uint             `json:"percent"`
}

type PoolTypeInfo struct {
	Items map[types.TxPoolType]uint
	Lock  *sync.RWMutex
}

func newPoolTypeInfo() *PoolTypeInfo {
	return &PoolTypeInfo{
		Items: map[types.TxPoolType]uint{
			normal: 100,
		},
		Lock: &sync.RWMutex{},
	}
}

func (p *PoolTypeInfo) update(items []*poolTypeItem) {
	log.Info("Update pool type info...")

	if items == nil {
		return
	}

	newItems := make(map[types.TxPoolType]uint, 0)

	var acc uint = 0
	for _, item := range items {
		if _, ok := newItems[item.PoolType]; ok {
			log.Error("Invalid pool type info, repeat type", "pool type", item.PoolType, "percent", item.Percent)
			return
		}

		if item.Percent > 100 {
			log.Error("Invalid pool type info, wrong percent", "pool type", item.PoolType, "percent", item.Percent)
			return
		}

		newItems[item.PoolType] = item.Percent
		acc += item.Percent
	}

	if acc > 100 {
		log.Error("Invalid pool type info, acculate percent greater than 100", "acculate percent ", acc)
		return
	}

	if acc < 100 {
		newItems[normal] = newItems[normal] + 100 - acc
	}
	p.Lock.Lock()
	defer p.Lock.Unlock()

	p.Items = newItems

	return
}
