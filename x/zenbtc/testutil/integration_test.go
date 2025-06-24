package testutil

import (
	"context"
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/Zenrock-Foundation/zrchain/v6/testutil/sample"
	"github.com/cosmos/cosmos-sdk/testutil"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/keeper"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	zenbtcKeeper  *keeper.Keeper
	ctx           sdk.Context
	msgServer     types.MsgServer
	accountKeeper *MockAccountKeeper
	bankKeeper    *MockBankKeeper
	paramSubspace *MockParamSubspace
	ctrl          *gomock.Controller
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupTest() {
	encCfg := moduletestutil.MakeTestEncodingConfig()
	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	s.ctx = testCtx.Ctx

	ctrl := gomock.NewController(s.T())
	accountKeeper := NewMockAccountKeeper(ctrl)
	bankKeeper := NewMockBankKeeper(ctrl)
	paramSubspace := NewMockParamSubspace(ctrl)

	s.accountKeeper = accountKeeper
	s.bankKeeper = bankKeeper
	s.paramSubspace = paramSubspace
	s.ctrl = ctrl

	s.zenbtcKeeper = keeper.NewKeeper(
		encCfg.Codec,
		storeService,
		testCtx.Ctx.Logger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		nil,
		nil,
	)

	s.Require().Equal(testCtx.Ctx.Logger().With("module", "x/"+types.ModuleName),
		s.zenbtcKeeper.Logger())

	err := s.zenbtcKeeper.Params.Set(s.ctx, *keeper.DefaultParams())
	s.Require().NoError(err)

	s.msgServer = keeper.NewMsgServerImpl(*s.zenbtcKeeper)
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_GetExchangeRate() {
	tests := []struct {
		name         string
		setupSupply  *types.Supply
		expectedRate string
		expectError  bool
	}{
		{
			name:         "Initial exchange rate when no supply exists",
			setupSupply:  nil,
			expectedRate: "1.000000000000000000",
		},
		{
			name: "Exchange rate with supply",
			setupSupply: &types.Supply{
				MintedZenBTC:  1000,
				PendingZenBTC: 500,
				CustodiedBTC:  2000,
			},
			expectedRate: "1.333333333333333333",
		},
		{
			name: "Exchange rate with zero zenBTC",
			setupSupply: &types.Supply{
				MintedZenBTC:  0,
				PendingZenBTC: 0,
				CustodiedBTC:  1000,
			},
			expectedRate: "1.000000000000000000",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Setup supply if provided
			if tt.setupSupply != nil {
				err := s.zenbtcKeeper.SetSupply(s.ctx, *tt.setupSupply)
				s.Require().NoError(err)
			}

			// Get exchange rate
			rate, err := s.zenbtcKeeper.GetExchangeRate(s.ctx)

			if tt.expectError {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tt.expectedRate, rate.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_PendingMintTransactions() {
	pendingTx := types.PendingMintTransaction{
		Id:               1,
		RecipientAddress: sample.AccAddress(),
		Amount:           1000,
		ChainType:        types.WalletType_WALLET_TYPE_EVM,
		Status:           types.MintTransactionStatus_MINT_TRANSACTION_STATUS_DEPOSITED,
		BlockHeight:      100,
		Caip2ChainId:     "eip155:1",
	}

	err := s.zenbtcKeeper.SetPendingMintTransaction(s.ctx, pendingTx)
	s.Require().NoError(err)

	var walkedTx types.PendingMintTransaction
	err = s.zenbtcKeeper.WalkPendingMintTransactions(s.ctx, func(id uint64, tx types.PendingMintTransaction) (stop bool, err error) {
		walkedTx = tx
		return true, nil
	})
	s.Require().NoError(err)
	s.Require().Equal(pendingTx.Id, walkedTx.Id)
	s.Require().Equal(pendingTx.RecipientAddress, walkedTx.RecipientAddress)
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_Redemptions() {
	redemption := types.Redemption{
		Data: types.RedemptionData{
			Id:                 1,
			Amount:             1000,
			DestinationAddress: []byte(sample.AccAddress()),
		},
		Status: types.RedemptionStatus_INITIATED,
	}

	err := s.zenbtcKeeper.SetRedemption(s.ctx, redemption.Data.Id, redemption)
	s.Require().NoError(err)

	exists, err := s.zenbtcKeeper.HasRedemption(s.ctx, redemption.Data.Id)
	s.Require().NoError(err)
	s.Require().True(exists)

	var walkedRedemption types.Redemption
	err = s.zenbtcKeeper.WalkRedemptions(s.ctx, func(id uint64, r types.Redemption) (stop bool, err error) {
		walkedRedemption = r
		return true, nil
	})
	s.Require().NoError(err)
	s.Require().Equal(redemption.Data.Id, walkedRedemption.Data.Id)
	s.Require().Equal(redemption.Data.Amount, walkedRedemption.Data.Amount)
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_Supply() {
	supply := types.Supply{
		MintedZenBTC:  1000,
		PendingZenBTC: 500,
		CustodiedBTC:  2000,
	}

	err := s.zenbtcKeeper.SetSupply(s.ctx, supply)
	s.Require().NoError(err)

	retrievedSupply, err := s.zenbtcKeeper.GetSupply(s.ctx)
	s.Require().NoError(err)
	s.Require().Equal(supply.MintedZenBTC, retrievedSupply.MintedZenBTC)
	s.Require().Equal(supply.PendingZenBTC, retrievedSupply.PendingZenBTC)
	s.Require().Equal(supply.CustodiedBTC, retrievedSupply.CustodiedBTC)
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_BurnEvents() {
	burnEvent := types.BurnEvent{
		Id:              1,
		Amount:          1000,
		DestinationAddr: []byte(sample.AccAddress()),
		Status:          types.BurnStatus_BURN_STATUS_BURNED,
		ChainID:         "eip155:1",
		TxID:            "0x1234567890abcdef",
	}

	err := s.zenbtcKeeper.SetBurnEvent(s.ctx, burnEvent.Id, burnEvent)
	s.Require().NoError(err)

	retrievedBurnEvent, err := s.zenbtcKeeper.GetBurnEvent(s.ctx, burnEvent.Id)
	s.Require().NoError(err)
	s.Require().Equal(burnEvent.Id, retrievedBurnEvent.Id)
	s.Require().Equal(burnEvent.Amount, retrievedBurnEvent.Amount)

	var walkedBurnEvent types.BurnEvent
	err = s.zenbtcKeeper.WalkBurnEvents(s.ctx, func(id uint64, be types.BurnEvent) (stop bool, err error) {
		walkedBurnEvent = be
		return true, nil
	})
	s.Require().NoError(err)
	s.Require().Equal(burnEvent.Id, walkedBurnEvent.Id)
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_PendingTransactionIndices() {
	testCases := []struct {
		name     string
		setValue uint64
		getFunc  func(context.Context) (uint64, error)
		setFunc  func(context.Context, uint64) error
	}{
		{
			name:     "FirstPendingEthMintTransaction",
			setValue: 100,
			getFunc:  s.zenbtcKeeper.GetFirstPendingEthMintTransaction,
			setFunc:  s.zenbtcKeeper.SetFirstPendingEthMintTransaction,
		},
		{
			name:     "FirstPendingSolMintTransaction",
			setValue: 200,
			getFunc:  s.zenbtcKeeper.GetFirstPendingSolMintTransaction,
			setFunc:  s.zenbtcKeeper.SetFirstPendingSolMintTransaction,
		},
		{
			name:     "FirstPendingBurnEvent",
			setValue: 300,
			getFunc:  s.zenbtcKeeper.GetFirstPendingBurnEvent,
			setFunc:  s.zenbtcKeeper.SetFirstPendingBurnEvent,
		},
		{
			name:     "FirstPendingRedemption",
			setValue: 400,
			getFunc:  s.zenbtcKeeper.GetFirstPendingRedemption,
			setFunc:  s.zenbtcKeeper.SetFirstPendingRedemption,
		},
		{
			name:     "FirstPendingStakeTransaction",
			setValue: 500,
			getFunc:  s.zenbtcKeeper.GetFirstPendingStakeTransaction,
			setFunc:  s.zenbtcKeeper.SetFirstPendingStakeTransaction,
		},
		{
			name:     "FirstRedemptionAwaitingSign",
			setValue: 600,
			getFunc:  s.zenbtcKeeper.GetFirstRedemptionAwaitingSign,
			setFunc:  s.zenbtcKeeper.SetFirstRedemptionAwaitingSign,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := tc.setFunc(s.ctx, tc.setValue)
			s.Require().NoError(err)

			retrievedValue, err := tc.getFunc(s.ctx)
			s.Require().NoError(err)
			s.Require().Equal(tc.setValue, retrievedValue)
		})
	}
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_ParameterGetters() {
	testCases := []struct {
		name          string
		getterFunc    func(context.Context) string
		expectedValue string
	}{
		{
			name:          "GetControllerAddr",
			getterFunc:    s.zenbtcKeeper.GetControllerAddr,
			expectedValue: keeper.DefaultControllerAddr,
		},
		{
			name:          "GetEthTokenAddr",
			getterFunc:    s.zenbtcKeeper.GetEthTokenAddr,
			expectedValue: keeper.DefaultEthTokenAddr,
		},
		{
			name:          "GetDepositKeyringAddr",
			getterFunc:    s.zenbtcKeeper.GetDepositKeyringAddr,
			expectedValue: keeper.DefaultDepositKeyringAddr,
		},
		{
			name:          "GetBitcoinProxyAddress",
			getterFunc:    s.zenbtcKeeper.GetBitcoinProxyAddress,
			expectedValue: keeper.DefaultTestnetBitcoinProxyAddress,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			value := tc.getterFunc(s.ctx)
			s.Require().Equal(tc.expectedValue, value)
		})
	}

	uint64TestCases := []struct {
		name          string
		getterFunc    func(context.Context) uint64
		expectedValue uint64
	}{
		{
			name:          "GetStakerKeyID",
			getterFunc:    s.zenbtcKeeper.GetStakerKeyID,
			expectedValue: keeper.DefaultStakerKeyID,
		},
		{
			name:          "GetEthMinterKeyID",
			getterFunc:    s.zenbtcKeeper.GetEthMinterKeyID,
			expectedValue: keeper.DefaultEthMinterKeyID,
		},
		{
			name:          "GetUnstakerKeyID",
			getterFunc:    s.zenbtcKeeper.GetUnstakerKeyID,
			expectedValue: keeper.DefaultUnstakerKeyID,
		},
		{
			name:          "GetCompleterKeyID",
			getterFunc:    s.zenbtcKeeper.GetCompleterKeyID,
			expectedValue: keeper.DefaultCompleterKeyID,
		},
		{
			name:          "GetRewardsDepositKeyID",
			getterFunc:    s.zenbtcKeeper.GetRewardsDepositKeyID,
			expectedValue: keeper.DefaultRewardsDepositKeyID,
		},
	}

	for _, tc := range uint64TestCases {
		s.Run(tc.name, func() {
			value := tc.getterFunc(s.ctx)
			s.Require().Equal(tc.expectedValue, value)
		})
	}

	s.Run("GetChangeAddressKeyIDs", func() {
		value := s.zenbtcKeeper.GetChangeAddressKeyIDs(s.ctx)
		s.Require().Equal(keeper.DefaultChangeAddressKeyIDs, value)
	})

	s.Run("GetSolanaParams", func() {
		solana := s.zenbtcKeeper.GetSolanaParams(s.ctx)
		s.Require().NotNil(solana)
		s.Require().Equal(keeper.DefaultSolana.SignerKeyId, solana.SignerKeyId)
		s.Require().Equal(keeper.DefaultSolana.ProgramId, solana.ProgramId)
	})
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_Authority() {
	authority := s.zenbtcKeeper.GetAuthority()
	s.Require().NotEmpty(authority)
	s.Require().Equal("zen10d07y265gmmuvt4z0w9aw880jnsr700jq0py6z", authority)
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_GetParams() {
	params, err := s.zenbtcKeeper.GetParams(s.ctx)
	s.Require().NoError(err)
	s.Require().NotNil(params)
	s.Require().Equal(keeper.DefaultDepositKeyringAddr, params.DepositKeyringAddr)
	s.Require().Equal(keeper.DefaultStakerKeyID, params.StakerKeyID)
	s.Require().Equal(keeper.DefaultEthMinterKeyID, params.EthMinterKeyID)
}

func (s *IntegrationTestSuite) Test_ZenbtcKeeper_StoreGetters() {
	pendingMintStore := s.zenbtcKeeper.GetPendingMintTransactionsStore()
	s.Require().NotNil(pendingMintStore)

	burnEventsStore := s.zenbtcKeeper.GetBurnEventsStore()
	s.Require().NotNil(burnEventsStore)

	redemptionsStore := s.zenbtcKeeper.GetRedemptionsStore()
	s.Require().NotNil(redemptionsStore)
}
