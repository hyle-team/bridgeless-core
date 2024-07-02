package keeper

import (
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
	amount          int64
	timestamp       time.Time
	unlockTimestamp time.Time
}

func NewNft(owner string, uri string, amount int64, unlockTimestamp time.Time) NFT {
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
