package v2

import (
	"errors"
	"fmt"
	"strings"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func UpdateParams(ctx sdk.Context, params collections.Item[types.Params]) error {
	paramsMap := map[string]types.Params{
		"zenrock": { // local
			EthBatcherAddr:      "0x912D79F8d489d0d007aBE0E26fD5d2f06BA4A2AA",
			DepositKeyringAddr:  "keyring1hpyh7xqr2w7h4eas5y8twnsg",
			WithdrawerKeyID:     1,
			EthMinterKeyID:      2,
			ChangeAddressKeyIDs: []uint64{3},
			UnstakerKeyID:       4,
			RewardsDepositKeyID: 5,
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			Authority:           "zen1sd3fwcpw2mdw3pxexmlg34gsd78r0sxrk5weh3",
		},
		"amber": { // devnet
			EthBatcherAddr:      "0x912D79F8d489d0d007aBE0E26fD5d2f06BA4A2AA",
			DepositKeyringAddr:  "keyring1k6vc6vhp6e6l3rxalue9v4ux",
			WithdrawerKeyID:     1,
			EthMinterKeyID:      2,
			ChangeAddressKeyIDs: []uint64{3},
			UnstakerKeyID:       4,
			RewardsDepositKeyID: 5,
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			Authority:           "zen1sd3fwcpw2mdw3pxexmlg34gsd78r0sxrk5weh3",
		},
		"gardia": { // testnet
			EthBatcherAddr:      "0x912D79F8d489d0d007aBE0E26fD5d2f06BA4A2AA",
			DepositKeyringAddr:  "keyring1k6vc6vhp6e6l3rxalue9v4ux",
			WithdrawerKeyID:     3,
			EthMinterKeyID:      4,
			ChangeAddressKeyIDs: []uint64{5},
			UnstakerKeyID:       6,
			RewardsDepositKeyID: 7,
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			Authority:           "zen1sd3fwcpw2mdw3pxexmlg34gsd78r0sxrk5weh3",
		},
		"diamond": { // mainnet
			EthBatcherAddr:      "",
			DepositKeyringAddr:  "keyring1k6vc6vhp6e6l3rxalue9v4ux",
			WithdrawerKeyID:     16,
			EthMinterKeyID:      17,
			ChangeAddressKeyIDs: []uint64{18},
			UnstakerKeyID:       19,
			RewardsDepositKeyID: 20,
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			Authority:           "zen1sd3fwcpw2mdw3pxexmlg34gsd78r0sxrk5weh3",
		},
	}

	chainID := ctx.ChainID()
	if chainID == "" {
		chainID = "zenrock"
	}

	newParams := types.Params{}

	for prefix, paramSet := range paramsMap {
		if strings.HasPrefix(chainID, prefix) {
			newParams = paramSet
			break
		}
	}

	if err := params.Set(ctx, newParams); err != nil {
		return err
	}

	return nil
}

func RemoveBadTestnetState(ctx sdk.Context, pendingMintTxCol collections.Item[types.PendingMintTransactions], codec codec.BinaryCodec) error {
	p, err := pendingMintTxCol.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil
		}
		return err
	}

	pTxs := types.PendingMintTransactions{}

	for _, pendingMintTx := range p.Txs {
		if pendingMintTx.ChainId != 17000 {
			continue
		}

		pTxs.Txs = append(pTxs.Txs, pendingMintTx)
	}

	if err := pendingMintTxCol.Set(ctx, pTxs); err != nil {
		return err
	}

	return nil
}

func ChangePendingMintTxChainIdtoCaip2Id(ctx sdk.Context, pendingMintTxCol collections.Item[types.PendingMintTransactions], createPendingMintTx func(sdk.Context, *types.PendingMintTransaction) error) error {
	p, err := pendingMintTxCol.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil
		}
		return err
	}

	for _, pendingMintTx := range p.Txs {
		switch pendingMintTx.ChainType {
		case types.WalletType_WALLET_TYPE_EVM:
			pendingMintTx.Caip2ChainId = fmt.Sprintf("eip155:%d", pendingMintTx.ChainId)
		}

		if err := createPendingMintTx(ctx, pendingMintTx); err != nil {
			return err
		}
	}

	return nil
}
