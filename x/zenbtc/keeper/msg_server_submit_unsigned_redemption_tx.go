package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/Zenrock-Foundation/zrchain/v5/bitcoin"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/zenrocklabs/zenrock/bitcoinproxy/libs/utils"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	treasurytypes "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k msgServer) SubmitUnsignedRedemptionTx(goCtx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) (*types.MsgSubmitUnsignedRedemptionTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//For Debugging Only
	if err := k.validationKeeper.ZenBTCRedemptions.Set(ctx, 0, types.Redemption{
		Data: types.RedemptionData{ // populate data below
			Id:                 1,
			DestinationAddress: []byte("tb1qqfgae6jr3xyawcz7uuh7ez06eteksxgf3up80x"),
			Amount:             10000,
		},
		Completed: false,
	}); err != nil {
		return nil, err
	}

	// TODO: here we must check if msg.Creator is the proxy

	//
	// TODO: verification against k.validationKeeper.ZenBTCRedemptions goes here
	//

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
	}
	k.treasuryKeeper.HandleSignatureRequest(ctx, sigReq)

	return &types.MsgSubmitUnsignedRedemptionTxResponse{}, nil
}

func (k msgServer) VerifyUnsignedRedemptionTX(ctx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) error {

	err := k.checkRedemptionTXCreator(ctx, msg)
	if err != nil {
		return err
	}

	msgTX, err := k.checkChangeAddress(ctx, msg)
	if err != nil {
		return err
	}

	//Loop the outputs, derive the address for the chain
	//Get the redemptions - get that each address is in pending redemptiopn adn amount is correct
	err = k.verifyOutputsAgainstRedemptions(ctx, msg, msgTX)
	if err != nil {
		return err
	}
	return nil
}

func (k msgServer) checkChangeAddress(ctx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) (*wire.MsgTx, error) {
	zenBTCChangeAddressKeyIDs := k.validationKeeper.GetZenBTCChangeAddressKeyIDs(ctx)
	if zenBTCChangeAddressKeyIDs == nil || len(zenBTCChangeAddressKeyIDs) == 0 {
		return nil, fmt.Errorf("Failed to retrieve zenBTCChangeAddressKeyIDs")
	}

	//Decode the Transactions
	msgTX, err := bitcoin.DecodeTX(msg.Txbytes)
	if err != nil {
		return nil, fmt.Errorf("Error decodeding txbytes: %s", err)
	}

	//Check 1st Output is Change
	if len(msgTX.TxOut) == 0 {
		return nil, fmt.Errorf("BTC Transaction has zero outputs")
	}
	chaincfg := utils.ChainFromString(msg.ChainName)
	changeOutput := msgTX.TxOut[0]

	//Loop all the change address and see if they match the change address in Output 0
	_, addrs, _, err := txscript.ExtractPkScriptAddrs(changeOutput.PkScript, chaincfg)
	if addrs == nil || len(addrs) != 1 {
		return nil, fmt.Errorf("BTC Change Address Invalid")
	}
	changeAddress := addrs[0].String()

	validChangeAddress := false
	for _, keyID := range zenBTCChangeAddressKeyIDs {
		key, err := k.Keeper.treasuryKeeper.KeyStore.Get(ctx, keyID)
		if err != nil {
			return nil, fmt.Errorf("BTC Error retrieving Change Addresses")
		}
		address, err := treasurytypes.BitcoinP2WPKH(&key, chaincfg)
		if err != nil {
			return nil, fmt.Errorf("BTC Error Generating Change Address from Key")
		}
		if address == changeAddress {
			validChangeAddress = true
			break
		}
	}
	if !validChangeAddress {
		return nil, fmt.Errorf("Invalid Change Address")
	}
	return msgTX, nil
}

func (k msgServer) checkRedemptionTXCreator(ctx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) error {
	bitcoinProxyCreatorID := k.validationKeeper.GetBitcoinProxyCreatorID(ctx)
	if bitcoinProxyCreatorID == "" {
		return fmt.Errorf("Failed to retrieve BitcoinProxyCreatorID")
	}

	if msg.Creator != bitcoinProxyCreatorID {
		return fmt.Errorf("msg.Creator must be the BitcoinProxyCreatorID")
	}
	return nil
}

func (k msgServer) verifyOutputsAgainstRedemptions(ctx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx, msgTX *wire.MsgTx) error {

	chaincfg := utils.ChainFromString(msg.ChainName)
	req := &types.QueryRedemptionsRequest{
		StartIndex:       0,
		CompletionFilter: types.CompletionFilter_PENDING,
	}
	redemptions, err := k.Keeper.Redemptions(ctx, req)
	if err != nil {
		return fmt.Errorf("Error retrieving Redemptions")
	}
	if redemptions == nil {
		return fmt.Errorf("Redemptions is nil")
	}
	if len(redemptions.Redemptions) == 0 {
		return fmt.Errorf("Redemptions is empty")
	}

	//Loop all TX Outputs (except change)
	for i, output := range msgTX.TxOut[1:] {
		outputIndex := uint64(i + 1) //don't use i directly as it always starts at 0, even though the TxOut[1:] starts at 1
		//Dervice output address
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, chaincfg)
		if err != nil || addrs == nil || len(addrs) != 1 {
			return fmt.Errorf("Redemption: Invalid Output Address %d output #", outputIndex)
		}
		err = verifyOutputInRedemptions(redemptions, addrs[0].String(), output.Value, msg.RedemptionIndexes[outputIndex])
		if err != nil {
			return fmt.Errorf("Invalid Output %d in tx:%s error:%w", outputIndex, hex.EncodeToString(msg.Txbytes), err)
		}
	}
	return nil
}

func verifyOutputInRedemptions(response *types.QueryRedemptionsResponse, outputAddress string, txAmount int64, redemptionIndex uint64) error {
	found := 0

	for _, redemption := range response.Redemptions {
		//check for duplicate outputAddress
		if redemption.Data.Id != redemptionIndex {
			continue
		}
		//Found the expected Redemption Index in the list of Redemptions
		found++
		//Check the Address
		redemptionAddress := string(redemption.Data.DestinationAddress)
		if redemptionAddress != outputAddress {
			return fmt.Errorf("Destination Address is invalid")
		}

		if int64(redemption.Data.Amount) != txAmount {
			return fmt.Errorf("Output has invalid amount TXOutput:%d RedemptionAmount:%d", txAmount, redemption.Data.Amount)
		}
	}
	switch found {
	case 0:
		return fmt.Errorf("Redemption not found")
	case 1:
		//Happy path case - Address & Amount in TX correspond to the redemptions
		return nil
	default:
		return fmt.Errorf("Redemption found more than once")
	}

}
