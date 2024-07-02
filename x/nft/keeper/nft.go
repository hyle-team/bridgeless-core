package keeper

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type NFT interface {
	Owner() string
	ID() uuid.UUID
	URI() string
	Delegate(address string)
}

type BaseNFT struct {
	id              uuid.UUID
	uri             string
	owner           string
	delegator       string
	amount          uint32
	timestamp       time.Time
	unlockTimestamp time.Time
	rewards         uint32
}

func NewNft(owner string, uri string, amount uint32, unlockTimestamp time.Time) NFT {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil
	}

	return &BaseNFT{
		id:              id,
		owner:           owner,
		uri:             uri,
		amount:          amount,
		unlockTimestamp: unlockTimestamp,
	}
}

func (n BaseNFT) ID() uuid.UUID {
	return n.id
}

func (n BaseNFT) Owner() string {
	return n.owner
}

func (n BaseNFT) URI() string {
	return n.uri
}

func (n BaseNFT) Delegate(address string) {
	n.delegator = address
}

func (n BaseNFT) CheckIsDelegated() bool {
	return n.delegator != ""
}

func (n BaseNFT) UnlockRewards() error {
	if n.CheckIsDelegated() {
		return fmt.Errorf("nft is delegated to %s", n.delegator)
	}

	if n.rewards == 0 {
		return fmt.Errorf("nft has no rewards to unlock")
	}

	// TODO make call to accumulator

	return nil
}
