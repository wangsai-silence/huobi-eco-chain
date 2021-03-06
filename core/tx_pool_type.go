package core

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	normal types.TxPoolType = 0
)

// PoolTypesResult is the query result for PoolTypes config.
type PoolTypesResult struct {
	List *PoolTypeList `json:"list"`
	Code int           `json:"code"` // 0: normal, 1: no list
}

// PoolTypeList is the config used to seprates transactions to different types
type PoolTypeList struct {
	TypeInfo map[types.TxPoolType]uint           `json:"typeInfo"`
	Froms    map[common.Address]types.TxPoolType `json:"froms"`
	Tos      map[common.Address]types.TxPoolType `json:"tos"`
}

// NewPoolTypeList returns a default PoolTypeList instance
func newPoolTypeList() *PoolTypeList {
	return &PoolTypeList{
		TypeInfo: map[types.TxPoolType]uint{normal: 100},
		Froms:    make(map[common.Address]types.TxPoolType),
		Tos:      make(map[common.Address]types.TxPoolType),
	}
}

func (p *PoolTypeList) check() (err error) {
	if p.TypeInfo == nil {
		return errors.New("Pool type info is nil")
	}

	total := 0
	for t, v := range p.TypeInfo {
		if v > 100 {
			return fmt.Errorf("type:%v percent is %v, greater than 100", t, v)
		}
		total += int(v)
	}
	if total > 100 {
		return errors.New("total percent is greater than 100")
	}

	return
}
