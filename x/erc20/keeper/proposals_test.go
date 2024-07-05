package keeper_test

import (
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/mock"

	"github.com/ethereum/go-ethereum/common"

	utiltx "github.com/hyle-team/bridgeless-core/testutil/tx"
	evmtypes "github.com/hyle-team/bridgeless-core/x/evm/types"

	"github.com/hyle-team/bridgeless-core/x/erc20/keeper"
	"github.com/hyle-team/bridgeless-core/x/erc20/types"
	inflationtypes "github.com/hyle-team/bridgeless-core/x/inflation/types"
)

const (
	contractMinterBurner = iota + 1
	contractDirectBalanceManipulation
	contractMaliciousDelayed
)

const (
	erc20Name          = "Coin Token"
	erc20Symbol        = "CTKN"
	erc20Decimals      = uint8(18)
	cosmosTokenBase    = "acoin"
	cosmosTokenDisplay = "coin"
	cosmosDecimals     = uint8(6)
	defaultExponent    = uint32(18)
	zeroExponent       = uint32(0)
	ibcBase            = "ibc/7B2A4F6E798182988D77B6B884919AF617A73503FDAC27C916CD7A69A69013CF"
)

var (
	metadataCoin = banktypes.Metadata{
		Description: "description of the token",
		Base:        cosmosTokenBase,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    cosmosTokenBase,
				Exponent: 0,
			},
			{
				Denom:    cosmosTokenDisplay,
				Exponent: defaultExponent,
			},
		},
		Name:    cosmosTokenBase,
		Symbol:  erc20Symbol,
		Display: cosmosTokenBase,
	}

	metadataIbc = banktypes.Metadata{
		Description: "ATOM IBC voucher (channel 14)",
		Base:        ibcBase,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    ibcBase,
				Exponent: 0,
			},
		},
		Name:    "ATOM channel-14",
		Symbol:  "ibcATOM-14",
		Display: ibcBase,
	}
)

func (suite *KeeperTestSuite) setupRegisterERC20Pair(contractType int) common.Address {
	var (
		contract common.Address
		err      error
	)
	// Deploy contract
	switch contractType {
	case contractDirectBalanceManipulation:
		contract, err = suite.DeployContractDirectBalanceManipulation()
	case contractMaliciousDelayed:
		contract, err = suite.DeployContractMaliciousDelayed()
	default:
		contract, err = suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
	}
	suite.Require().NoError(err)
	suite.Commit()

	_, err = suite.app.Erc20Keeper.RegisterERC20(suite.ctx, contract)
	suite.Require().NoError(err)
	return contract
}

func (suite *KeeperTestSuite) setupRegisterCoin(metadata banktypes.Metadata) *types.TokenPair {
	err := suite.app.BankKeeper.MintCoins(suite.ctx, inflationtypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
	suite.Require().NoError(err)

	// pair := types.NewTokenPair(contractAddr, cosmosTokenBase, true, types.OWNER_MODULE)
	pair, err := suite.app.Erc20Keeper.RegisterCoin(suite.ctx, metadata)
	suite.Require().NoError(err)
	suite.Commit()
	return pair
}

func (suite KeeperTestSuite) TestRegisterCoin() { //nolint:govet // we can copy locks here because it is a test
	metadata := banktypes.Metadata{
		Description: "description",
		Base:        cosmosTokenBase,
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    cosmosTokenBase,
				Exponent: 0,
			},
			{
				Denom:    cosmosTokenDisplay,
				Exponent: defaultExponent,
			},
		},
		Name:    cosmosTokenBase,
		Symbol:  erc20Symbol,
		Display: cosmosTokenDisplay,
	}

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"conversion is disabled globally",
			func() {
				params := types.DefaultParams()
				params.EnableErc20 = false
				suite.app.Erc20Keeper.SetParams(suite.ctx, params) //nolint:errcheck
			},
			false,
		},
		{
			"denom already registered",
			func() {
				regPair := types.NewTokenPair(utiltx.GenerateAddress(), metadata.Base, types.OWNER_MODULE)
				suite.app.Erc20Keeper.SetDenomMap(suite.ctx, regPair.Denom, regPair.GetID())
				suite.Commit()
			},
			false,
		},
		{
			"token doesn't have supply",
			func() {
			},
			false,
		},
		{
			"metadata different that stored",
			func() {
				metadata.Base = cosmosTokenBase
				validMetadata := banktypes.Metadata{
					Description: "description",
					Base:        cosmosTokenBase,
					// NOTE: Denom units MUST be increasing
					DenomUnits: []*banktypes.DenomUnit{
						{
							Denom:    cosmosTokenBase,
							Exponent: 0,
						},
						{
							Denom:    cosmosTokenDisplay,
							Exponent: uint32(18),
						},
					},
					Name:    erc20Name,
					Symbol:  erc20Symbol,
					Display: cosmosTokenDisplay,
				}

				err := suite.app.BankKeeper.MintCoins(suite.ctx, inflationtypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(validMetadata.Base, 1)})
				suite.Require().NoError(err)
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, validMetadata)
			},
			false,
		},
		{
			"ok",
			func() {
				metadata.Base = cosmosTokenBase
				err := suite.app.BankKeeper.MintCoins(suite.ctx, inflationtypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"force fail evm",
			func() {
				metadata.Base = cosmosTokenBase
				err := suite.app.BankKeeper.MintCoins(suite.ctx, inflationtypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
				suite.Require().NoError(err)

				mockEVMKeeper := &MockEVMKeeper{}

				suite.app.Erc20Keeper = keeper.NewKeeper(
					suite.app.GetKey("erc20"), suite.app.AppCodec(),
					authtypes.NewModuleAddress(govtypes.ModuleName), suite.app.AccountKeeper,
					suite.app.BankKeeper, mockEVMKeeper, suite.app.StakingKeeper, suite.app.ClaimsKeeper)

				mockEVMKeeper.On("EstimateGas", mock.Anything, mock.Anything).Return(&evmtypes.EstimateGasResponse{Gas: uint64(200)}, nil)
				mockEVMKeeper.On("ApplyMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("forced ApplyMessage error"))
			},
			false,
		},
		{
			"force delete module account evm",
			func() {
				metadata.Base = cosmosTokenBase
				err := suite.app.BankKeeper.MintCoins(suite.ctx, inflationtypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(metadata.Base, 1)})
				suite.Require().NoError(err)

				acc := suite.app.AccountKeeper.GetAccount(suite.ctx, types.ModuleAddress.Bytes())
				suite.app.AccountKeeper.RemoveAccount(suite.ctx, acc)
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			tc.malleate()

			pair, err := suite.app.Erc20Keeper.RegisterCoin(suite.ctx, metadata)
			suite.Commit()

			expPair := &types.TokenPair{
				Erc20Address:  "0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd",
				Denom:         "acoin",
				Enabled:       true,
				ContractOwner: 1,
			}

			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				suite.Require().Equal(pair, expPair)
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}

func (suite KeeperTestSuite) TestRegisterERC20() { //nolint:govet // we can copy locks here because it is a test
	var (
		contractAddr common.Address
		pair         types.TokenPair
	)
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"token ERC20 already registered",
			func() {
				suite.app.Erc20Keeper.SetERC20Map(suite.ctx, pair.GetERC20Contract(), pair.GetID())
			},
			false,
		},
		{
			"denom already registered",
			func() {
				suite.app.Erc20Keeper.SetDenomMap(suite.ctx, pair.Denom, pair.GetID())
			},
			false,
		},
		{
			"meta data already stored",
			func() {
				suite.app.Erc20Keeper.CreateCoinMetadata(suite.ctx, contractAddr) //nolint:errcheck
			},
			false,
		},
		{
			"ok",
			func() {},
			true,
		},
		{
			"force fail evm",
			func() {
				mockEVMKeeper := &MockEVMKeeper{}

				suite.app.Erc20Keeper = keeper.NewKeeper(
					suite.app.GetKey("erc20"), suite.app.AppCodec(),
					authtypes.NewModuleAddress(govtypes.ModuleName), suite.app.AccountKeeper,
					suite.app.BankKeeper, mockEVMKeeper, suite.app.StakingKeeper, suite.app.ClaimsKeeper)

				mockEVMKeeper.On("EstimateGas", mock.Anything, mock.Anything).Return(&evmtypes.EstimateGasResponse{Gas: uint64(200)}, nil)
				mockEVMKeeper.On("ApplyMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("forced ApplyMessage error"))
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			var err error
			suite.SetupTest() // reset

			contractAddr, err = suite.DeployContract(erc20Name, erc20Symbol, cosmosDecimals)
			suite.Require().NoError(err)

			coinName := types.CreateDenom(contractAddr.String())
			pair = types.NewTokenPair(contractAddr, coinName, types.OWNER_EXTERNAL)

			tc.malleate()

			_, err = suite.app.Erc20Keeper.RegisterERC20(suite.ctx, contractAddr)
			metadata, found := suite.app.BankKeeper.GetDenomMetaData(suite.ctx, coinName)
			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				// Metadata variables
				suite.Require().True(found)
				suite.Require().Equal(coinName, metadata.Base)
				suite.Require().Equal(coinName, metadata.Name)
				suite.Require().Equal(types.SanitizeERC20Name(erc20Name), metadata.Display)
				suite.Require().Equal(erc20Symbol, metadata.Symbol)
				// Denom units
				suite.Require().Equal(len(metadata.DenomUnits), 2)
				suite.Require().Equal(coinName, metadata.DenomUnits[0].Denom)
				suite.Require().Equal(zeroExponent, metadata.DenomUnits[0].Exponent)
				suite.Require().Equal(types.SanitizeERC20Name(erc20Name), metadata.DenomUnits[1].Denom)
				// Custom exponent at contract creation matches coin with token
				suite.Require().Equal(metadata.DenomUnits[1].Exponent, uint32(cosmosDecimals))
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}

func (suite KeeperTestSuite) TestToggleConverision() { //nolint:govet // we can copy locks here because it is a test
	var (
		contractAddr common.Address
		id           []byte
		pair         types.TokenPair
	)

	testCases := []struct {
		name              string
		malleate          func()
		expPass           bool
		conversionEnabled bool
	}{
		{
			"token not registered",
			func() {
				contractAddr, err := suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
				suite.Commit()
				pair = types.NewTokenPair(contractAddr, cosmosTokenBase, types.OWNER_MODULE)
			},
			false,
			false,
		},
		{
			"token not registered - pair not found",
			func() {
				contractAddr, err := suite.DeployContract(erc20Name, erc20Symbol, erc20Decimals)
				suite.Require().NoError(err)
				suite.Commit()
				pair = types.NewTokenPair(contractAddr, cosmosTokenBase, types.OWNER_MODULE)
				suite.app.Erc20Keeper.SetERC20Map(suite.ctx, common.HexToAddress(pair.Erc20Address), pair.GetID())
			},
			false,
			false,
		},
		{
			"disable conversion",
			func() {
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				id = suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
				pair, _ = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
			},
			true,
			false,
		},
		{
			"disable and enable conversion",
			func() {
				contractAddr = suite.setupRegisterERC20Pair(contractMinterBurner)
				id = suite.app.Erc20Keeper.GetTokenPairID(suite.ctx, contractAddr.String())
				pair, _ = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
				pair, _ = suite.app.Erc20Keeper.ToggleConversion(suite.ctx, contractAddr.String())
			},
			true,
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			tc.malleate()

			var err error
			pair, err = suite.app.Erc20Keeper.ToggleConversion(suite.ctx, contractAddr.String())
			// Request the pair using the GetPairToken func to make sure that is updated on the db
			pair, _ = suite.app.Erc20Keeper.GetTokenPair(suite.ctx, id)
			if tc.expPass {
				suite.Require().NoError(err, tc.name)
				if tc.conversionEnabled {
					suite.Require().True(pair.Enabled)
				} else {
					suite.Require().False(pair.Enabled)
				}
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}
