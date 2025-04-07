package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/hyle-team/bridgeless-core/v12/cmd/config"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k *Keeper) EndBlocker(ctx sdk.Context) []abci.ValidatorUpdate {
	params := k.GetParams(ctx)
	validators := k.stakingKeeper.GetAllValidators(ctx)
	if len(validators) == 0 {
		return []abci.ValidatorUpdate{}
	}
	var paramsUpdated bool
	for _, validator := range validators {
		// Validator`s operator address is a wrapper over creator address bytes
		decodedAddr, err := sdk.Bech32ifyAddressBytes(config.Bech32Prefix, validator.GetOperator().Bytes())
		if err != nil {
			continue
		}
		owner, err := sdk.AccAddressFromBech32(decodedAddr)
		if err != nil {
			continue
		}

		partiesIndex := partyIndex(owner.String(), params.Parties)
		newbieIndex := partyIndex(owner.String(), params.Newbies)
		goodbyeIndex := partyIndex(owner.String(), params.GoodbyeList)

		delegationIsEnough := isEnoughDelegation(owner.String(), k.stakingKeeper.GetValidatorDelegations(ctx,
			validator.GetOperator()), params.StakeThreshold)
		if isPartyInList(partiesIndex) {
			if !delegationIsEnough {
				if !isPartyInList(goodbyeIndex) {
					// If party is an active Tss party but his delegation is not enough it is moved to goodbye list
					params.GoodbyeList = append(params.GoodbyeList, &types.Party{Address: owner.String()})
					paramsUpdated = true
					continue
				}
			}
			continue
		}

		// If party already in newbies list then check whether he has enough delegation
		if isPartyInList(newbieIndex) && !delegationIsEnough {
			// If party was a newbie but his delegation became too small remove him from newbies
			params.Newbies = append(params.Newbies[:newbieIndex], params.Newbies[newbieIndex+1:]...)
			paramsUpdated = true
			continue
		}

		// If party was not processed yet and not in blacklist add it to newbies if it has enough delegation
		if !isPartyInList(partyIndex(owner.String(), params.Blacklist)) && delegationIsEnough {
			if !isPartyInList(newbieIndex) {
				params.Newbies = append(params.Newbies, &types.Party{Address: owner.String()})
				paramsUpdated = true
			}
		}
	}

	if paramsUpdated {
		k.SetParams(ctx, params)
	}

	return []abci.ValidatorUpdate{}
}

func partyIndex(party string, partyList []*types.Party) int {
	for i, p := range partyList {
		if party == p.Address {
			return i
		}
	}
	return -1
}

func isPartyInList(i int) bool {
	return i >= 0
}

func isEnoughDelegation(partyAddr string, delegations stakingTypes.Delegations, stakingThreshold string) bool {
	for _, d := range delegations {
		if partyAddr == d.DelegatorAddress {
			n, _ := sdk.NewDecFromStr(stakingThreshold)
			return !d.Amount.LT(n)
		}
	}
	return false
}
