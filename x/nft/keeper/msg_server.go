package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/evmos/evmos/v12/x/nft/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) Mint(ctx context.Context, tokens *types.MsgMintTokens) (*types.MsgMontTokensResponse, error) {
	nft := NewNft(
		tokens.Owner,
		tokens.Uri,
		tokens.Amount,
		tokens.UnlockTimestemp.AsTime(),
	)

	m.AppendDelegatedNft(tokens.Owner, nft)

	return nil, nil
}

func (m msgServer) UnlockRewards(ctx context.Context, rewards *types.MsgUnlockRewards) (*types.MsgUnlockRewardsResponse, error) {
	m.DelegatedNft(rewards.String())

	return nil, nil
}

func (m msgServer) Delegate(ctx context.Context, delegate *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	err := m.DelegateNft(delegate.Recipient, delegate.Owner, delegate.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delegate nft")
	}
	return nil, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
