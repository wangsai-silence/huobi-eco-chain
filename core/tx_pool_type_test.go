package core

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	str := `{"code":0,"list":{"typeInfo":{"items":{"0":60,"1":40}},"froms":{"0x3aa2c5da4bdaca180bf8e92eb9f18b22f9dcc006":1,"0x3d0090fb5f20e0559632c95b877bb1478c2f1111":1},"tos":{"0xe959f432f9244dc5d421455669ca4be591277fd4":1,"0xed7d5f38c79115ca12fe6c0041abb22f0a06c300":1}}}`
	res := PoolTypesResult{}
	err := json.Unmarshal([]byte(str), &res)
	require.NoError(t, err)
	addr := common.HexToAddress("0x3aa2c5da4bdaca180bf8e92eb9f18b22f9dcc006")
	require.EqualValues(t, 1, res.List.Froms[addr])
}
