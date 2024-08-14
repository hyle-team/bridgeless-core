package types

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ModuleName defines the module name
	ModuleName = "bridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bridge"

	// ----- Param Keys -----
	ParamModuleAdminKey = "ModuleAdmin"
	ParamEvmAdminKey    = "EvmAdmin"

	// ---- Store Prefixes ------
	StoreTokenPrefix       = "token/"
	StoreTokenPairsPrefix  = "token-pairs/"
	StoreChainPrefix       = "chain/"
	StoreTransactionPrefix = "transaction/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeyToken(id uint64) []byte {
	return []byte(strconv.FormatInt(int64(id), 10))
}

func TokenPairPrexif(srcChain, srcAddr string) []byte {
	return []byte(fmt.Sprintf("%s-%s", srcChain, strings.ToLower(srcAddr)))
}

func KeyTokenPair(dstChain string) []byte {
	return []byte(dstChain)
}

func KeyChain(chain string) []byte {
	return []byte(chain)
}

func KeyTransaction(id string) []byte {
	return []byte(id)
}
