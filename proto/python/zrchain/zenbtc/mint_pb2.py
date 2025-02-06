# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/zenbtc/mint.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'zrchain/zenbtc/mint.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19zrchain/zenbtc/mint.proto\x12\x0ezrchain.zenbtc\"n\n\tNonceData\x12\x14\n\x05nonce\x18\x01 \x01(\x04R\x05nonce\x12\x18\n\x07\x63ounter\x18\x02 \x01(\x04R\x07\x63ounter\x12\x12\n\x04skip\x18\x03 \x01(\x08R\x04skip\x12\x1d\n\nprev_nonce\x18\x04 \x01(\x04R\tprevNonce\"3\n\x17RequestedBitcoinHeaders\x12\x18\n\x07heights\x18\x01 \x03(\x03R\x07heights\"\xb6\x01\n\x0fLockTransaction\x12\x15\n\x06raw_tx\x18\x01 \x01(\tR\x05rawTx\x12\x12\n\x04vout\x18\x02 \x01(\x04R\x04vout\x12\x16\n\x06sender\x18\x03 \x01(\tR\x06sender\x12%\n\x0emint_recipient\x18\x04 \x01(\tR\rmintRecipient\x12\x16\n\x06\x61mount\x18\x05 \x01(\x04R\x06\x61mount\x12!\n\x0c\x62lock_height\x18\x06 \x01(\x03R\x0b\x62lockHeight\"\xdd\x02\n\x16PendingMintTransaction\x12\x1d\n\x08\x63hain_id\x18\x01 \x01(\x04\x42\x02\x18\x01R\x07\x63hainId\x12\x39\n\nchain_type\x18\x02 \x01(\x0e\x32\x1a.zrchain.zenbtc.WalletTypeR\tchainType\x12+\n\x11recipient_address\x18\x03 \x01(\tR\x10recipientAddress\x12\x16\n\x06\x61mount\x18\x04 \x01(\x04R\x06\x61mount\x12\x18\n\x07\x63reator\x18\x05 \x01(\tR\x07\x63reator\x12\x15\n\x06key_id\x18\x06 \x01(\x04R\x05keyId\x12$\n\x0e\x63\x61ip2_chain_id\x18\x07 \x01(\tR\x0c\x63\x61ip2ChainId\x12\x0e\n\x02id\x18\x08 \x01(\x04R\x02id\x12=\n\x06status\x18\t \x01(\x0e\x32%.zrchain.zenbtc.MintTransactionStatusR\x06status\"W\n\x17PendingMintTransactions\x12\x38\n\x03txs\x18\x01 \x03(\x0b\x32&.zrchain.zenbtc.PendingMintTransactionR\x03txs:\x02\x18\x01*\xaf\x01\n\x15MintTransactionStatus\x12\'\n#MINT_TRANSACTION_STATUS_UNSPECIFIED\x10\x00\x12%\n!MINT_TRANSACTION_STATUS_DEPOSITED\x10\x01\x12\"\n\x1eMINT_TRANSACTION_STATUS_STAKED\x10\x02\x12\"\n\x1eMINT_TRANSACTION_STATUS_MINTED\x10\x03*\xc4\x01\n\nWalletType\x12\x1b\n\x17WALLET_TYPE_UNSPECIFIED\x10\x00\x12\x16\n\x12WALLET_TYPE_NATIVE\x10\x01\x12\x13\n\x0fWALLET_TYPE_EVM\x10\x02\x12\x1b\n\x17WALLET_TYPE_BTC_TESTNET\x10\x03\x12\x1b\n\x17WALLET_TYPE_BTC_MAINNET\x10\x04\x12\x1a\n\x16WALLET_TYPE_BTC_REGNET\x10\x05\x12\x16\n\x12WALLET_TYPE_SOLANA\x10\x06\x42.Z,github.com/zenrocklabs/zenbtc/x/zenbtc/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.zenbtc.mint_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,github.com/zenrocklabs/zenbtc/x/zenbtc/types'
  _globals['_PENDINGMINTTRANSACTION'].fields_by_name['chain_id']._loaded_options = None
  _globals['_PENDINGMINTTRANSACTION'].fields_by_name['chain_id']._serialized_options = b'\030\001'
  _globals['_PENDINGMINTTRANSACTIONS']._loaded_options = None
  _globals['_PENDINGMINTTRANSACTIONS']._serialized_options = b'\030\001'
  _globals['_MINTTRANSACTIONSTATUS']._serialized_start=837
  _globals['_MINTTRANSACTIONSTATUS']._serialized_end=1012
  _globals['_WALLETTYPE']._serialized_start=1015
  _globals['_WALLETTYPE']._serialized_end=1211
  _globals['_NONCEDATA']._serialized_start=45
  _globals['_NONCEDATA']._serialized_end=155
  _globals['_REQUESTEDBITCOINHEADERS']._serialized_start=157
  _globals['_REQUESTEDBITCOINHEADERS']._serialized_end=208
  _globals['_LOCKTRANSACTION']._serialized_start=211
  _globals['_LOCKTRANSACTION']._serialized_end=393
  _globals['_PENDINGMINTTRANSACTION']._serialized_start=396
  _globals['_PENDINGMINTTRANSACTION']._serialized_end=745
  _globals['_PENDINGMINTTRANSACTIONS']._serialized_start=747
  _globals['_PENDINGMINTTRANSACTIONS']._serialized_end=834
# @@protoc_insertion_point(module_scope)
