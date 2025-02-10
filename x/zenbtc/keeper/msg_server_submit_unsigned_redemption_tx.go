package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	treasurytypes "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k msgServer) SubmitUnsignedRedemptionTx(goCtx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) (*types.MsgSubmitUnsignedRedemptionTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//For Debugging Only

	//if err := k.validationKeeper.ZenBTCRedemptions.Set(ctx, 1, types.Redemption{
	//	Data: types.RedemptionData{ // populate data below
	//		Id:                 1,
	//		DestinationAddress: []byte("tb1qypwjx7yj5jz0gw0vh76348ypa2ns7tfwsnhlh9"),
	//		Amount:             9000,
	//	},
	//	Completed: false,
	//}); err != nil {
	//	return nil, err
	//}
	//if err := k.validationKeeper.ZenBTCRedemptions.Set(ctx, 2, types.Redemption{
	//	Data: types.RedemptionData{ // populate data below
	//		Id:                 2,
	//		DestinationAddress: []byte("tb1qq9f55z0qs8ergppxv54gyrvglhhjtkx84gd8vq"),
	//		Amount:             8000,
	//	},
	//	Completed: false,
	//}); err != nil {
	//	return nil, err
	//}

	err := k.VerifyUnsignedRedemptionTX(ctx, msg)
	if err != nil {
		return nil, err
	}

	keyIDs := make([]uint64, len(msg.Inputs))
	hashes := make([]string, len(msg.Inputs))
	for i, input := range msg.Inputs {
		keyIDs[i] = input.Keyid
		hashes[i] = input.Hash
	}

	sigReq := &treasurytypes.MsgNewSignatureRequest{
		Creator:        msg.Creator,
		KeyIds:         keyIDs,
		DataForSigning: strings.Join(hashes, ","), // hex string, each unsigned utxo is separated by comma
		CacheId:        msg.CacheId,
		ZenbtcTxBytes:  msg.Txbytes,
	}
	if _, err := k.treasuryKeeper.HandleSignatureRequest(ctx, sigReq); err != nil {
		return nil, err
	}

	return &types.MsgSubmitUnsignedRedemptionTxResponse{}, nil
}
