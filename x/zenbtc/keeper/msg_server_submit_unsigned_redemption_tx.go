package keeper

import (
	"context"
	"fmt"
	"github.com/Zenrock-Foundation/zrchain/v5/bitcoin"
	"github.com/btcsuite/btcd/txscript"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	treasurytypes "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k msgServer) SubmitUnsignedRedemptionTx(goCtx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) (*types.MsgSubmitUnsignedRedemptionTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//For Debugging Only
	//if err := k.validationKeeper.ZenBTCRedemptions.Set(ctx, 0, types.Redemption{
	//	Data: types.RedemptionData{ // populate data below
	//		Id:                 1,
	//		DestinationAddress: []byte("tb1qqfgae6jr3xyawcz7uuh7ez06eteksxgf3up80x"),
	//		Amount:             10000,
	//	},
	//	Completed: false,
	//}); err != nil {
	//	return nil, err
	//}

	// TODO: here we must check if msg.Creator is the proxy

	//
	// TODO: verification against k.validationKeeper.ZenBTCRedemptions goes here
	//

	err := k.VerifyUnsignedRedemptionTX(ctx, msg)

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
	}
	k.treasuryKeeper.HandleSignatureRequest(ctx, sigReq)

	return &types.MsgSubmitUnsignedRedemptionTxResponse{}, nil
}

func (k msgServer) VerifyUnsignedRedemptionTX(ctx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) error {
	bitcoinProxyCreatorID := k.validationKeeper.GetBitcoinProxyCreatorID(ctx)
	if bitcoinProxyCreatorID == "" {
		return fmt.Errorf("Failed to retrieve BitcoinProxyCreatorID")
	}
	zenBTCChangeAddressKeyIDs := k.validationKeeper.GetZenBTCChangeAddressKeyIDs(ctx)
	if zenBTCChangeAddressKeyIDs == nil || len(zenBTCChangeAddressKeyIDs) == 0 {
		return fmt.Errorf("Failed to retrieve zenBTCChangeAddressKeyIDs")
	}

	if msg.Creator != bitcoinProxyCreatorID {
		return fmt.Errorf("msg.Creator must be the BitcoinProxyCreatorID")
	}

	//Decode the Transactions
	msgTX, err := bitcoin.DecodeTX(msg.Txbytes)
	if err != nil {
		return fmt.Errorf("Error decodeding txbytes: %s", err)
	}

	//Check 1st Output is Change
	if len(msgTX.TxOut) == 0 {
		return fmt.Errorf("BTC Transaction has zero outputs")
	}

	chain := msg.ChainName

	changeOutput := msgTX.TxOut[0]
	for _, keyID := range zenBTCChangeAddressKeyIDs {

		_, addrs, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, chaincfg)

		if changeOutput == byte(keyID) {
			return nil
		}
	}
	return fmt.Errorf("BTC Transaction first output is not a change address")
	return nil
}
