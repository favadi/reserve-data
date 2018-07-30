package http

import (
	ethereum "github.com/ethereum/go-ethereum/common"
)

type HTTPBlockchain interface {
	LoadAndSetTokenIndices([]ethereum.Address) error
	CheckTokenIndices(ethereum.Address) error
}
