package keeper

import (
	"fmt"
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

		partyIndex := isPartyInList(owner.String(), params.Parties)
		newbieIndex := isPartyInList(owner.String(), params.Newbies)
		goodbyeIndex := isPartyInList(owner.String(), params.GoodbyeList)

		delegationIsEnough := isEnoughDelegation(owner.String(), k.stakingKeeper.GetValidatorDelegations(ctx,
			validator.GetOperator()), params.StakeThreshold)
		fmt.Println(owner.String(), " index: ", partyIndex)
		if partyIndex != -1 {
			fmt.Println(delegationIsEnough)
			if !delegationIsEnough {
				if goodbyeIndex == -1 {
					// If party is an active Tss party but his delegation is not enough it is moved to goodbye list
					params.GoodbyeList = append(params.GoodbyeList, &types.Party{Address: owner.String()})
					paramsUpdated = true
					continue
				}
			}
			continue
		}

		// If party already in newbies list then check whether he has enough delegation
		if newbieIndex != -1 && !delegationIsEnough {
			// If party was a newbie but his delegation became too small remove him from newbies
			params.Newbies = append(params.Newbies[:newbieIndex], params.Newbies[newbieIndex+1:]...)
			paramsUpdated = true
			continue
		}

		// If party is now in goodbye list but his delegation became eligible it is removed from list
		if goodbyeIndex != -1 && delegationIsEnough {
			params.GoodbyeList = append(params.GoodbyeList[:partyIndex], params.GoodbyeList[partyIndex+1:]...)
			paramsUpdated = true
			continue
		}

		// If party was not processed yet and not in blacklist add it to newbies if it has enough delegation
		if isPartyInList(owner.String(), params.Blacklist) == -1 && delegationIsEnough {
			if newbieIndex == -1 {
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

func isPartyInList(party string, partyList []*types.Party) int {
	for i, p := range partyList {
		if party == p.Address {
			return i
		}
	}
	return -1
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
