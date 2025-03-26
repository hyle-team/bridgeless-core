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
	fmt.Println("Bridge end blocker")
	if len(validators) == 0 {
		fmt.Println("No validators were found")
	}
	fmt.Println(validators)
	for _, validator := range validators {
		// Validator`s operator address is a wrapper over creator address bytes
		decodedAddr, err := sdk.Bech32ifyAddressBytes(config.Bech32Prefix, validator.GetOperator().Bytes())
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error decoding validator operator address : %s", err.Error()))
			continue
		}
		owner, err := sdk.AccAddressFromBech32(decodedAddr)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error decoding validator operator address : %s", err.Error()))
			continue
		}

		if isPartyInList(owner.String(), params.Parties) {
			if !isEnoughDelegation(owner.String(), k.stakingKeeper.GetValidatorDelegations(ctx, validator.GetOperator()),
				params.StakeThreshold) {

				if !isPartyInList(owner.String(), params.GoodbyeList) {
					// If party is an active Tss party but his delegation is not enough it is moved to goodbye list
					params.GoodbyeList = append(params.GoodbyeList, &types.Party{Address: owner.String()})
					continue
				}
			}
		}

		if !isPartyInList(owner.String(), params.Parties) {
			// If party already in newbies list than check whether he has enough delegation
			if isPartyInList(owner.String(), params.Newbies) && !isEnoughDelegation(owner.String(),
				k.stakingKeeper.GetValidatorDelegations(ctx, validator.GetOperator()), params.StakeThreshold) {

				// If party was a newbie but his delegation became too small remove him from newbies
				params.Newbies = append(params.Newbies[:partyIndex(owner.String(), params.Newbies)],
					params.Newbies[partyIndex(owner.String(), params.Newbies)+1:]...)
				continue
			}

			// If party is now in goodbye list but his delegation became eligible it is removed from list
			if isPartyInList(owner.String(), params.GoodbyeList) && isEnoughDelegation(owner.String(),
				k.stakingKeeper.GetValidatorDelegations(ctx, validator.GetOperator()), params.StakeThreshold) {

				params.GoodbyeList = append(params.GoodbyeList[:partyIndex(owner.String(), params.GoodbyeList)],
					params.GoodbyeList[partyIndex(owner.String(), params.GoodbyeList)+1:]...)
				continue
			}

			// If party was not processed yet and not in blacklist add it to newbies if it has enough delegation
			if !isPartyInList(owner.String(), params.Blacklist) && isEnoughDelegation(owner.String(),
				k.stakingKeeper.GetValidatorDelegations(ctx, validator.GetOperator()), params.StakeThreshold) &&
				!isPartyInList(owner.String(), params.Newbies) {

				params.Newbies = append(params.Newbies, &types.Party{Address: owner.String()})
			}
		}

	}

	k.SetParams(ctx, params)

	return []abci.ValidatorUpdate{}
}

func isPartyInList(party string, partyList []*types.Party) bool {
	for _, p := range partyList {
		if party == p.Address {
			return true
		}
	}
	return false
}

func partyIndex(party string, partyList []*types.Party) int {
	for i, p := range partyList {
		if party == p.Address {
			return i
		}
	}
	return -1
}

func isEnoughDelegation(partyAddr string, delegations stakingTypes.Delegations, stakingThreshold int64) bool {
	for _, d := range delegations {
		if partyAddr == d.DelegatorAddress {
			if !d.Amount.LT(sdk.NewDec(stakingThreshold)) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
