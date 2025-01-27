package v2

import (
	"strings"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func UpdateParams(ctx sdk.Context, params collections.Item[types.Params]) error {
	paramsMap := map[string]types.Params{
		"zenrock": { // local
			EthBatcherAddr:      "0x912D79F8d489d0d007aBE0E26fD5d2f06BA4A2AA",
			DepositKeyringAddr:  "keyring1hpyh7xqr2w7h4eas5y8twnsg",
			WithdrawerKeyID:     1,
			MinterKeyID:         2,
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
			MinterKeyID:         2,
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
			MinterKeyID:         4,
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
			MinterKeyID:         17,
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
