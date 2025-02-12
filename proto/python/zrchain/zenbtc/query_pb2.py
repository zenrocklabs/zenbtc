# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/zenbtc/query.proto
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
    'zrchain/zenbtc/query.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from cosmos.base.query.v1beta1 import pagination_pb2 as cosmos_dot_base_dot_query_dot_v1beta1_dot_pagination__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from zrchain.zenbtc import mint_pb2 as zrchain_dot_zenbtc_dot_mint__pb2
from zrchain.zenbtc import params_pb2 as zrchain_dot_zenbtc_dot_params__pb2
from zrchain.zenbtc import redemptions_pb2 as zrchain_dot_zenbtc_dot_redemptions__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1azrchain/zenbtc/query.proto\x12\x0ezrchain.zenbtc\x1a\x11\x61mino/amino.proto\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a\x14gogoproto/gogo.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x19zrchain/zenbtc/mint.proto\x1a\x1bzrchain/zenbtc/params.proto\x1a zrchain/zenbtc/redemptions.proto\"\x14\n\x12QueryParamsRequest\"P\n\x13QueryParamsResponse\x12\x39\n\x06params\x18\x01 \x01(\x0b\x32\x16.zrchain.zenbtc.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06params\"\x1e\n\x1cQueryLockTransactionsRequest\"m\n\x1dQueryLockTransactionsResponse\x12L\n\x11lock_transactions\x18\x01 \x03(\x0b\x32\x1f.zrchain.zenbtc.LockTransactionR\x10lockTransactions\"t\n\x17QueryRedemptionsRequest\x12\x1f\n\x0bstart_index\x18\x01 \x01(\x04R\nstartIndex\x12\x38\n\x06status\x18\x02 \x01(\x0e\x32 .zrchain.zenbtc.RedemptionStatusR\x06status\"^\n\x18QueryRedemptionsResponse\x12\x42\n\x0bredemptions\x18\x01 \x03(\x0b\x32\x1a.zrchain.zenbtc.RedemptionB\x04\xc8\xde\x1f\x00R\x0bredemptions\"\x85\x01\n#QueryPendingMintTransactionsRequest\x12\x1f\n\x0bstart_index\x18\x01 \x01(\x04R\nstartIndex\x12=\n\x06status\x18\x02 \x01(\x0e\x32%.zrchain.zenbtc.MintTransactionStatusR\x06status\"\x8a\x01\n$QueryPendingMintTransactionsResponse\x12\x62\n\x19pending_mint_transactions\x18\x01 \x03(\x0b\x32&.zrchain.zenbtc.PendingMintTransactionR\x17pendingMintTransactions\"\x14\n\x12QuerySupplyRequest\"\xc9\x01\n\x13QuerySupplyResponse\x12\"\n\x0c\x63ustodiedBTC\x18\x01 \x01(\x04R\x0c\x63ustodiedBTC\x12 \n\x0btotalZenBTC\x18\x02 \x01(\x04R\x0btotalZenBTC\x12\"\n\x0cmintedZenBTC\x18\x03 \x01(\x04R\x0cmintedZenBTC\x12$\n\rpendingZenBTC\x18\x04 \x01(\x04R\rpendingZenBTC\x12\"\n\x0c\x65xchangeRate\x18\x05 \x01(\tR\x0c\x65xchangeRate\"\xb7\x01\n\x16QueryBurnEventsRequest\x12\x1f\n\x0bstart_index\x18\x01 \x01(\x04R\nstartIndex\x12\x12\n\x04txID\x18\x02 \x01(\tR\x04txID\x12\x1a\n\x08logIndex\x18\x03 \x01(\x04R\x08logIndex\x12\x18\n\x07\x63hainID\x18\x04 \x01(\tR\x07\x63hainID\x12\x32\n\x06status\x18\x05 \x01(\x0e\x32\x1a.zrchain.zenbtc.BurnStatusR\x06status\"T\n\x17QueryBurnEventsResponse\x12\x39\n\nburnEvents\x18\x01 \x03(\x0b\x32\x19.zrchain.zenbtc.BurnEventR\nburnEvents2\xb5\x06\n\x05Query\x12l\n\tGetParams\x12\".zrchain.zenbtc.QueryParamsRequest\x1a#.zrchain.zenbtc.QueryParamsResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/zenbtc/params\x12\x92\x01\n\x10LockTransactions\x12,.zrchain.zenbtc.QueryLockTransactionsRequest\x1a-.zrchain.zenbtc.QueryLockTransactionsResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/zenbtc/lock_transactions\x12\x80\x01\n\x0eGetRedemptions\x12\'.zrchain.zenbtc.QueryRedemptionsRequest\x1a(.zrchain.zenbtc.QueryRedemptionsResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/zenbtc/redemptions\x12\xb4\x01\n\x1cQueryPendingMintTransactions\x12\x33.zrchain.zenbtc.QueryPendingMintTransactionsRequest\x1a\x34.zrchain.zenbtc.QueryPendingMintTransactionsResponse\")\x82\xd3\xe4\x93\x02#\x12!/zenbtc/pending_mint_transactions\x12n\n\x0bQuerySupply\x12\".zrchain.zenbtc.QuerySupplyRequest\x1a#.zrchain.zenbtc.QuerySupplyResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/zenbtc/supply\x12\x7f\n\x0fQueryBurnEvents\x12&.zrchain.zenbtc.QueryBurnEventsRequest\x1a\'.zrchain.zenbtc.QueryBurnEventsResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/zenbtc/burn_eventsB.Z,github.com/zenrocklabs/zenbtc/x/zenbtc/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.zenbtc.query_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,github.com/zenrocklabs/zenbtc/x/zenbtc/types'
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['params']._loaded_options = None
  _globals['_QUERYPARAMSRESPONSE'].fields_by_name['params']._serialized_options = b'\310\336\037\000\250\347\260*\001'
  _globals['_QUERYREDEMPTIONSRESPONSE'].fields_by_name['redemptions']._loaded_options = None
  _globals['_QUERYREDEMPTIONSRESPONSE'].fields_by_name['redemptions']._serialized_options = b'\310\336\037\000'
  _globals['_QUERY'].methods_by_name['GetParams']._loaded_options = None
  _globals['_QUERY'].methods_by_name['GetParams']._serialized_options = b'\202\323\344\223\002\020\022\016/zenbtc/params'
  _globals['_QUERY'].methods_by_name['LockTransactions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['LockTransactions']._serialized_options = b'\202\323\344\223\002\033\022\031/zenbtc/lock_transactions'
  _globals['_QUERY'].methods_by_name['GetRedemptions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['GetRedemptions']._serialized_options = b'\202\323\344\223\002\025\022\023/zenbtc/redemptions'
  _globals['_QUERY'].methods_by_name['QueryPendingMintTransactions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['QueryPendingMintTransactions']._serialized_options = b'\202\323\344\223\002#\022!/zenbtc/pending_mint_transactions'
  _globals['_QUERY'].methods_by_name['QuerySupply']._loaded_options = None
  _globals['_QUERY'].methods_by_name['QuerySupply']._serialized_options = b'\202\323\344\223\002\020\022\016/zenbtc/supply'
  _globals['_QUERY'].methods_by_name['QueryBurnEvents']._loaded_options = None
  _globals['_QUERY'].methods_by_name['QueryBurnEvents']._serialized_options = b'\202\323\344\223\002\025\022\023/zenbtc/burn_events'
  _globals['_QUERYPARAMSREQUEST']._serialized_start=251
  _globals['_QUERYPARAMSREQUEST']._serialized_end=271
  _globals['_QUERYPARAMSRESPONSE']._serialized_start=273
  _globals['_QUERYPARAMSRESPONSE']._serialized_end=353
  _globals['_QUERYLOCKTRANSACTIONSREQUEST']._serialized_start=355
  _globals['_QUERYLOCKTRANSACTIONSREQUEST']._serialized_end=385
  _globals['_QUERYLOCKTRANSACTIONSRESPONSE']._serialized_start=387
  _globals['_QUERYLOCKTRANSACTIONSRESPONSE']._serialized_end=496
  _globals['_QUERYREDEMPTIONSREQUEST']._serialized_start=498
  _globals['_QUERYREDEMPTIONSREQUEST']._serialized_end=614
  _globals['_QUERYREDEMPTIONSRESPONSE']._serialized_start=616
  _globals['_QUERYREDEMPTIONSRESPONSE']._serialized_end=710
  _globals['_QUERYPENDINGMINTTRANSACTIONSREQUEST']._serialized_start=713
  _globals['_QUERYPENDINGMINTTRANSACTIONSREQUEST']._serialized_end=846
  _globals['_QUERYPENDINGMINTTRANSACTIONSRESPONSE']._serialized_start=849
  _globals['_QUERYPENDINGMINTTRANSACTIONSRESPONSE']._serialized_end=987
  _globals['_QUERYSUPPLYREQUEST']._serialized_start=989
  _globals['_QUERYSUPPLYREQUEST']._serialized_end=1009
  _globals['_QUERYSUPPLYRESPONSE']._serialized_start=1012
  _globals['_QUERYSUPPLYRESPONSE']._serialized_end=1213
  _globals['_QUERYBURNEVENTSREQUEST']._serialized_start=1216
  _globals['_QUERYBURNEVENTSREQUEST']._serialized_end=1399
  _globals['_QUERYBURNEVENTSRESPONSE']._serialized_start=1401
  _globals['_QUERYBURNEVENTSRESPONSE']._serialized_end=1485
  _globals['_QUERY']._serialized_start=1488
  _globals['_QUERY']._serialized_end=2309
# @@protoc_insertion_point(module_scope)
