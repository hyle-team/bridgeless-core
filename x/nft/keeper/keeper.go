package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/evmos/evmos/v12/x/nft/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		nfts          map[string][]NFT
		delegatedNfts map[string][]NFT
		commonAmount  uint64
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		delegatedNfts: make(map[string][]NFT),
		nfts:          make(map[string][]NFT),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) checkIsAdmin(address string) bool {
	// TODO call accumulator module
	return true
}

func (k Keeper) AppendNft(address string, nft NFT) {
	k.nfts[address] = append(k.nfts[address], nft)
}

func (k Keeper) AppendDelegatedNft(address string, nft NFT) {
	k.delegatedNfts[address] = append(k.delegatedNfts[address], nft)
}

func (k Keeper) DelegatedNft(address string) []NFT {
	return k.delegatedNfts[address]
}

func (k Keeper) NFTsByAddress(address string) []NFT {
	return k.nfts[address]
}

func (k Keeper) DelegateNft(recipient, owner string, id uint32) error {
	nfts, ok := k.nfts[owner]
	if !ok {
		return fmt.Errorf("owner not found")
	}

	for _, nft := range nfts {
		if nft.ID().ID() == id {
			nft.Delegate(recipient)
			k.AppendDelegatedNft(recipient, nft)
			return nil
		}
	}

	return fmt.Errorf("ntf with id %d not found", id)
}
