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

	if err := k.validationKeeper.ZenBTCRedemptions.Set(ctx, 1, types.Redemption{
		Data: types.RedemptionData{ // populate data below
			Id:                 1,
			DestinationAddress: []byte("tb1qypwjx7yj5jz0gw0vh76348ypa2ns7tfwsnhlh9"),
			Amount:             9000,
		},
		Completed: false,
	}); err != nil {
		return nil, err
	}
	if err := k.validationKeeper.ZenBTCRedemptions.Set(ctx, 2, types.Redemption{
		Data: types.RedemptionData{ // populate data below
			Id:                 2,
			DestinationAddress: []byte("tb1qq9f55z0qs8ergppxv54gyrvglhhjtkx84gd8vq"),
			Amount:             8000,
		},
		Completed: false,
	}); err != nil {
		return nil, err
	}

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
	// Check that the transaction creator is valid
	if err := k.checkRedemptionTXCreator(ctx, msg); err != nil {
		return fmt.Errorf("failed to verify transaction creator: %w", err)
	}

	// Check that the change address is valid and retrieve the transaction
	msgTX, err := k.checkChangeAddress(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to check change address: %w", err)
	}

	// Verify that the outputs match the redemptions
	if err := k.verifyOutputsAgainstRedemptions(ctx, msg, msgTX); err != nil {
		return fmt.Errorf("failed to verify outputs against redemptions: %w", err)
	}

	// Update the redemptions to complete
	if err := k.updateRedemptionsComplete(ctx, msg.RedemptionIndexes); err != nil {
		return fmt.Errorf("failed to update redemptions to complete: %w", err)
	}

	return nil
}

func (k msgServer) updateRedemptionsComplete(ctx context.Context, redemptionIndexes []uint64) error {
	if len(redemptionIndexes) == 0 {
		return fmt.Errorf("no redemption indexes provided")
	}

	// Iterate over the redemption indexes, starting from the second element
	for _, redemptionIndex := range redemptionIndexes[1:] {
		// Retrieve the redemption entry
		redemption, err := k.validationKeeper.ZenBTCRedemptions.Get(ctx, redemptionIndex)
		if err != nil {
			return fmt.Errorf("failed to retrieve redemption at index %d: %w", redemptionIndex, err)
		}

		// Mark the redemption as completed
		redemption.Completed = true

		// Save the updated redemption entry
		if err := k.validationKeeper.ZenBTCRedemptions.Set(ctx, redemptionIndex, redemption); err != nil {
			return fmt.Errorf("failed to update redemption at index %d: %w", redemptionIndex, err)
		}
	}
	return nil
}

func (k msgServer) checkChangeAddress(ctx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) (*wire.MsgTx, error) {
	zenBTCChangeAddressKeyIDs := k.validationKeeper.GetZenBTCChangeAddressKeyIDs(ctx)
	if zenBTCChangeAddressKeyIDs == nil || len(zenBTCChangeAddressKeyIDs) == 0 {
		return nil, fmt.Errorf("failed to retrieve ZenBTCChangeAddressKeyIDs")
	}

	// Decode the transactions
	msgTX, err := bitcoin.DecodeTX(msg.Txbytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding txbytes: %w", err)
	}

	// Check that the first output is the change
	if len(msgTX.TxOut) == 0 {
		return nil, fmt.Errorf("BTC transaction has zero outputs")
	}
	chaincfg := utils.ChainFromString(msg.ChainName)
	changeOutput := msgTX.TxOut[0]

	// Extract and validate change address from Output 0
	_, addrs, _, err := txscript.ExtractPkScriptAddrs(changeOutput.PkScript, chaincfg)
	if err != nil || addrs == nil || len(addrs) != 1 {
		return nil, fmt.Errorf("BTC change address invalid")
	}
	changeAddress := addrs[0].String()

	// Validate the change address against known change addresses
	validChangeAddress := false
	for _, keyID := range zenBTCChangeAddressKeyIDs {
		key, err := k.Keeper.treasuryKeeper.KeyStore.Get(ctx, keyID)
		if err != nil {
			return nil, fmt.Errorf("error retrieving change addresses for keyID %s: %w", keyID, err)
		}
		address, err := treasurytypes.BitcoinP2WPKH(&key, chaincfg)
		if err != nil {
			return nil, fmt.Errorf("error generating change address from keyID %s: %w", keyID, err)
		}
		if address == changeAddress {
			validChangeAddress = true
			break
		}
	}
	if !validChangeAddress {
		return nil, fmt.Errorf("invalid change address: %s", changeAddress)
	}
	return msgTX, nil
}

func (k msgServer) checkRedemptionTXCreator(ctx context.Context, msg *types.MsgSubmitUnsignedRedemptionTx) error {
	bitcoinProxyCreatorID := k.validationKeeper.GetBitcoinProxyCreatorID(ctx)
	if bitcoinProxyCreatorID == "" {
		return fmt.Errorf("failed to retrieve BitcoinProxyCreatorID")
	}

	if msg.Creator != bitcoinProxyCreatorID {
		return fmt.Errorf("msg.Creator (%s) must be the BitcoinProxyCreatorID (%s)", msg.Creator, bitcoinProxyCreatorID)
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
		return fmt.Errorf("error retrieving redemptions: %w", err)
	}
	if redemptions == nil {
		return fmt.Errorf("redemptions is nil")
	}
	if len(redemptions.Redemptions) == 0 {
		return fmt.Errorf("redemptions is empty")
	}

	// Verify that the length of RedemptionIndexes matches the number of TxOut
	if len(msg.RedemptionIndexes) != len(msgTX.TxOut) {
		return fmt.Errorf("redemptionIndexes length (%d) does not match number of outputs (%d)", len(msg.RedemptionIndexes), len(msgTX.TxOut))
	}

	// Loop through all Tx Outputs (ignoring the change output at index 0)
	for i, output := range msgTX.TxOut[1:] {
		outputIndex := uint64(i + 1) // Adjust index to correspond to TxOut slice

		// Derive output address
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, chaincfg)
		if err != nil || addrs == nil || len(addrs) != 1 {
			return fmt.Errorf("invalid output address at output index %d", outputIndex)
		}

		// Verify the output against redemptions
		err = verifyOutputInRedemptions(redemptions, addrs[0].String(), output.Value, msg.RedemptionIndexes[outputIndex])
		if err != nil {
			return fmt.Errorf("invalid output %d in tx: %s, error: %w", outputIndex, hex.EncodeToString(msg.Txbytes), err)
		}
	}
	return nil
}

func verifyOutputInRedemptions(response *types.QueryRedemptionsResponse, outputAddress string, txAmount int64, redemptionIndex uint64) error {
	var found int

	for _, redemption := range response.Redemptions {
		// Check for matching redemption index
		if redemption.Data.Id != redemptionIndex {
			continue
		}

		// Found the expected Redemption Index
		found++

		// Validate the address
		if string(redemption.Data.DestinationAddress) != outputAddress {
			return fmt.Errorf("destination address is invalid: expected %s, got %s", outputAddress, redemption.Data.DestinationAddress)
		}

		// Validate the amount
		if int64(redemption.Data.Amount) != txAmount {
			return fmt.Errorf("output has invalid amount: expected %d, got %d", txAmount, redemption.Data.Amount)
		}
	}

	// Validate the number of found redemptions
	switch found {
	case 0:
		return fmt.Errorf("redemption not found")
	case 1:
		// Happy path: Redemption index, address, and amount are valid
		return nil
	default:
		return fmt.Errorf("redemption found more than once")
	}
}
