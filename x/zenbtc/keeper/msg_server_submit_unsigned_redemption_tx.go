package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	treasurytypes "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k msgServer) SubmitUnsignedRedemptionTx(goCtx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) (*types.MsgSubmitUnsignedRedemptionTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: here we must check if msg.Creator is the proxy

	//
	// TODO: verification against k.validationKeeper.ZenBTCRedemptions goes here
	//

	k.treasuryKeeper.HandleSignatureRequest(ctx, &treasurytypes.MsgNewSignatureRequest{
		Creator:        msg.Creator,
		KeyId:          k.validationKeeper.GetZenBTCWithdrawerKeyID(ctx),
		DataForSigning: msg.Outputs, // hex string, each unsigned utxo is separated by comma
	})

	return &types.MsgSubmitUnsignedRedemptionTxResponse{}, nil
}
