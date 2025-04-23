package keeper

var (
	MsgVerifyDepositBlockInclusion = "/zrchain.zenbtc.MsgVerifyDepositBlockInclusion"
	MsgSubmitUnsignedRedemptionTx  = "/zrchain.zenbtc.MsgSubmitUnsignedRedemptionTx"
)

func NewFeeExempts(noFeeMsgs []string) []string {
	return append(noFeeMsgs, MsgVerifyDepositBlockInclusion, MsgSubmitUnsignedRedemptionTx)
}

func DefaultNoFeeMsgs() []string {
	return []string{
		MsgVerifyDepositBlockInclusion,
		MsgSubmitUnsignedRedemptionTx,
	}
}

func DefaultFeeExempts() []string {
	return NewFeeExempts(DefaultNoFeeMsgs())
}
