package v2

import (
	"strings"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

// TODO: update with correct params
func UpdateParams(ctx sdk.Context, params collections.Item[types.Params]) error {
	paramsMap := map[string]types.Params{
		"zenrock": { // local
			DepositKeyringAddr:  "keyring1hpyh7xqr2w7h4eas5y8twnsg",
			StakerKeyID:         1,
			EthMinterKeyID:      2,
			UnstakerKeyID:       3,
			CompleterKeyID:      4,
			RewardsDepositKeyID: 5,
			ChangeAddressKeyIDs: []uint64{6},
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			EthTokenAddr:        "0xC8CdeDd20cCb4c06884ac4C2fF952A0B7cC230a3",
			ControllerAddr:      "0x5b9Ea8d5486D388a158F026c337DF950866dA5e9",
		},
		"amber": { // devnet
			DepositKeyringAddr:  "keyring1k6vc6vhp6e6l3rxalue9v4ux",
			StakerKeyID:         1,
			EthMinterKeyID:      2,
			UnstakerKeyID:       3,
			CompleterKeyID:      4,
			RewardsDepositKeyID: 5,
			ChangeAddressKeyIDs: []uint64{6},
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			EthTokenAddr:        "0xC8CdeDd20cCb4c06884ac4C2fF952A0B7cC230a3",
			ControllerAddr:      "0x5b9Ea8d5486D388a158F026c337DF950866dA5e9",
		},
		"gardia": { // testnet
			DepositKeyringAddr:  "keyring1k6vc6vhp6e6l3rxalue9v4ux",
			StakerKeyID:         1,
			EthMinterKeyID:      2,
			UnstakerKeyID:       3,
			CompleterKeyID:      4,
			RewardsDepositKeyID: 5,
			ChangeAddressKeyIDs: []uint64{6},
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			EthTokenAddr:        "0xC8CdeDd20cCb4c06884ac4C2fF952A0B7cC230a3",
			ControllerAddr:      "0x5b9Ea8d5486D388a158F026c337DF950866dA5e9",
		},
		"diamond": { // mainnet
			DepositKeyringAddr:  "keyring1k6vc6vhp6e6l3rxalue9v4ux",
			StakerKeyID:         1,
			EthMinterKeyID:      2,
			UnstakerKeyID:       3,
			CompleterKeyID:      4,
			RewardsDepositKeyID: 5,
			ChangeAddressKeyIDs: []uint64{6},
			BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
			EthTokenAddr:        "0xC8CdeDd20cCb4c06884ac4C2fF952A0B7cC230a3",
			ControllerAddr:      "0x5b9Ea8d5486D388a158F026c337DF950866dA5e9",
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
