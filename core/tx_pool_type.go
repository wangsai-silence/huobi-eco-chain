package core

import (
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
