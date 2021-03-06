package common

import (
	"github.com/nknorg/nkn/errors"
	"github.com/nknorg/nkn/net/node"
	"github.com/nknorg/nkn/vault"
)

type Serverer interface {
	GetNetNode() (*node.LocalNode, error)
	GetWallet() (vault.Wallet, error)
}

func respPacking(result interface{}, errcode ErrCode) map[string]interface{} {
	resp := map[string]interface{}{
		"result": result,
		"error":  errcode,
	}
	return resp
}

func respPackingDetails(result interface{}, errcode ErrCode, details errors.ErrCode) map[string]interface{} {
	resp := map[string]interface{}{
		"result":  result,
		"error":   errcode,
		"details": details,
	}
	return resp
}

func RespPacking(result interface{}, errcode ErrCode) map[string]interface{} {
	return respPacking(result, errcode)
}

func ResponsePack(errCode ErrCode) map[string]interface{} {
	resp := map[string]interface{}{
		"Action":  "",
		"Result":  "",
		"Error":   errCode,
		"Desc":    "",
		"Version": "1.0.0",
	}
	return resp
}
