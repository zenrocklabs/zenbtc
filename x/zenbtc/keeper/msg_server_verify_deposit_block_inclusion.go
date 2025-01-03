package keeper

import (
	"context"
	"errors"
	"net/rpc"

	"github.com/zenrocklabs/zenbtc/utils"

	"github.com/btcsuite/btcd/chaincfg/chainhash"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Zenrock-Foundation/zrchain/v5/bitcoin"
	"github.com/Zenrock-Foundation/zrchain/v5/sidecar/proto/api"
	treasurytypes "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k msgServer) VerifyDepositBlockInclusion(goCtx context.Context, msg *types.MsgVerifyDepositBlockInclusion) (*types.MsgVerifyDepositBlockInclusionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	blockHeader, err := k.validationKeeper.BtcBlockHeaders.Get(ctx, msg.BlockHeight)

	//CSM For Debugging Only
	//try and get missing Blockheader over RPC - WARNING for debugging only!!!!
	//if err != nil {
	//	bh, _ := debugRetrieveBlockHeaderViaRPC(msg.ChainName, msg.BlockHeight)
	//	if bh != nil {
	//		err = nil
	//		blockHeader = *bh
	//	}
	//}
	// END of debug code

	if err != nil {
		return nil, err
	}

	found := false

	//List of addresses to ignore - we don't want to cause a mint for change
	ignoreAddresses, err := k.ZenBTCChangeAddresses(ctx, msg.ChainName)
	if err != nil {
		return nil, errors.New("Error Retrieving the Change Addresses")
	}

	//Verify the blockheader is valid and the proof, return a list of outputs in the transaction
	outputs, _, err := bitcoin.VerifyBTCLockTransaction(msg.RawTx, msg.ChainName, int(msg.Index), msg.Proof, &blockHeader, ignoreAddresses)
	if err != nil {
		return nil, err
	}

	//Check the address & amount specified is actually in the supplied (proven) BTC Transaction
	for _, output := range outputs {
		if output.Address == msg.DepositAddr && output.Amount == msg.Amount {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("zenBTC deposit not found in outputs from provided lock tx")
	}

	q, err := k.treasuryKeeper.KeyByAddress(ctx, &treasurytypes.QueryKeyByAddressRequest{
		Address:     msg.DepositAddr,
		KeyringAddr: k.validationKeeper.GetZenBTCDepositKeyringAddr(ctx),
		KeyType:     treasurytypes.KeyType_KEY_TYPE_BITCOIN_SECP256K1,
		WalletType:  WalletTypeFromChainName(msg),
	})
	if err != nil {
		return nil, err
	}

	if q.Response == nil || q.Response.Wallets == nil {
		return nil, errors.New("zenBTC deposit address does not correspond to correct key (no wallets returned)")
	}

	found = false
	for _, wallet := range q.Response.Wallets {
		if wallet.Address == msg.DepositAddr {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("zenBTC deposit address does not correspond to correct key (no matching wallet)")
	}

	// Deposit/lock txs are stored in zenBTC module so they can't be used to mint zenBTC tokens more than once
	txExists, err := k.LockTransactionStore.Has(ctx, msg.RawTx)
	if err != nil {
		return nil, err
	}
	if txExists {
		return nil, errors.New("lock tx was already previously used to mint zenBTC tokens")
	}

	supply, err := k.validationKeeper.ZenBTCSupply.Get(ctx)
	if err != nil {
		if !errors.Is(err, collections.ErrNotFound) {
			return nil, err
		}
		supply = types.Supply{CustodiedBTC: 0, MintedZenBTC: 0}
	}

	supply.CustodiedBTC += msg.Amount

	if err := k.validationKeeper.ZenBTCSupply.Set(ctx, supply); err != nil {
		return nil, err
	}

	if err := k.LockTransactionStore.Set(ctx, msg.RawTx); err != nil {
		return nil, err
	}

	// Don't mint zenBTC tokens for rewards deposits; return early
	if q.Response.Key.Id == k.validationKeeper.GetZenBTCRewardsDepositKeyID(ctx) {
		return &types.MsgVerifyDepositBlockInclusionResponse{}, nil
	}

	pendingTxs, err := k.validationKeeper.PendingMintTransactions.Get(ctx)
	if err != nil {
		if !errors.Is(err, collections.ErrNotFound) {
			return nil, err
		}
		pendingTxs = treasurytypes.PendingMintTransactions{Txs: []*treasurytypes.PendingMintTransaction{}}
	}

	// Calculate amount of zenBTC to mint based on current exchange rate
	exchangeRate, err := k.validationKeeper.GetZenBTCExchangeRate(ctx)
	if err != nil {
		return nil, err
	}

	// Amount of zenBTC to mint is the BTC amount divided by BTC/zenBTC exchange rate
	amount := float64(msg.Amount) / exchangeRate

	pendingTxs.Txs = append(pendingTxs.Txs, &treasurytypes.PendingMintTransaction{
		ChainId:          q.Response.Key.ZenbtcMetadata.ChainId,
		ChainType:        q.Response.Key.ZenbtcMetadata.ChainType,
		RecipientAddress: q.Response.Key.ZenbtcMetadata.RecipientAddr,
		Amount:           uint64(amount),
		Creator:          msg.Creator,
		KeyId:            k.validationKeeper.GetZenBTCMinterKeyID(ctx),
	})
	if err := k.validationKeeper.PendingMintTransactions.Set(ctx, pendingTxs); err != nil {
		return nil, err
	}

	if err := k.validationKeeper.EthereumNonceRequested.Set(ctx, k.validationKeeper.GetZenBTCMinterKeyID(ctx), true); err != nil {
		return nil, err
	}

	return &types.MsgVerifyDepositBlockInclusionResponse{}, nil
}

func debugRetrieveBlockHeaderViaRPC(chainName string, blockHeight int64) (*api.BTCBlockHeader, error) {
	if chainName == "mainnet" {
		return nil, errors.New("cannot retrieve block header from mainnet")
	}
	type GetBlockHeaderByHeightArgs struct {
		ChainName string
		Height    int64
	}

	type GetBlockHeaderByHeightReply struct {
		BlockHeader *api.BTCBlockHeader
		BlockHash   *chainhash.Hash
		Height      int32
	}

	args := GetBlockHeaderByHeightArgs{
		ChainName: chainName,
		Height:    blockHeight,
	}
	var resp GetBlockHeaderByHeightReply
	client, err := rpc.DialHTTP("tcp", "localhost"+":12345")
	if err != nil {
		return nil, err
	}

	err = client.Call("NeutrinoServer.BlockHeaderByHeight", args, &resp)
	if err != nil {
		return nil, err
	}
	return resp.BlockHeader, nil
}

/*
Get the list of Change KeyID and derive the addresses for the correct Chain
*/
func (k msgServer) ZenBTCChangeAddresses(ctx context.Context, chainName string) ([]string, error) {
	keyIDs := k.validationKeeper.GetZenBTCChangeAddressKeyIDs(ctx)
	chaincfg := utils.ChainFromString(chainName)
	result := make([]string, 0)
	for _, keyID := range keyIDs {
		key, err := k.Keeper.treasuryKeeper.KeyStore.Get(ctx, keyID)
		if err != nil {
			return nil, err
		}
		address, err := treasurytypes.BitcoinP2WPKH(&key, chaincfg)
		if err != nil {
			return nil, err
		}
		result = append(result, address)
	}
	return result, nil
}

func WalletTypeFromChainName(msg *types.MsgVerifyDepositBlockInclusion) treasurytypes.WalletType {
	switch msg.ChainName {
	case "mainnet":
		return treasurytypes.WalletType_WALLET_TYPE_BTC_MAINNET
	case "regtest":
		return treasurytypes.WalletType_WALLET_TYPE_BTC_REGNET
	case "testnet", "testnet3", "testnet4":
		return treasurytypes.WalletType_WALLET_TYPE_BTC_TESTNET
	default:
		return treasurytypes.WalletType_WALLET_TYPE_UNSPECIFIED
	}
}
