# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/zenbtc/query.proto
# Protobuf Python Version: 5.28.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
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
from zrchain.zenbtc import params_pb2 as zrchain_dot_zenbtc_dot_params__pb2
from zrchain.zenbtc import redemptions_pb2 as zrchain_dot_zenbtc_dot_redemptions__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1azrchain/zenbtc/query.proto\x12\x0ezrchain.zenbtc\x1a\x11\x61mino/amino.proto\x1a*cosmos/base/query/v1beta1/pagination.proto\x1a\x14gogoproto/gogo.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1bzrchain/zenbtc/params.proto\x1a zrchain/zenbtc/redemptions.proto\"\x14\n\x12QueryParamsRequest\"P\n\x13QueryParamsResponse\x12\x39\n\x06params\x18\x01 \x01(\x0b\x32\x16.zrchain.zenbtc.ParamsB\t\xc8\xde\x1f\x00\xa8\xe7\xb0*\x01R\x06params\"\x1e\n\x1cQueryLockTransactionsRequest\"L\n\x1dQueryLockTransactionsResponse\x12+\n\x11lock_transactions\x18\x01 \x03(\tR\x10lockTransactions\"\x89\x01\n\x17QueryRedemptionsRequest\x12\x1f\n\x0bstart_index\x18\x01 \x01(\x04R\nstartIndex\x12M\n\x11\x63ompletion_filter\x18\x02 \x01(\x0e\x32 .zrchain.zenbtc.CompletionFilterR\x10\x63ompletionFilter\"^\n\x18QueryRedemptionsResponse\x12\x42\n\x0bredemptions\x18\x01 \x03(\x0b\x32\x1a.zrchain.zenbtc.RedemptionB\x04\xc8\xde\x1f\x00R\x0bredemptions*7\n\x10\x43ompletionFilter\x12\x07\n\x03\x41LL\x10\x00\x12\x0b\n\x07PENDING\x10\x01\x12\r\n\tCOMPLETED\x10\x02\x32\x86\x03\n\x05Query\x12i\n\x06Params\x12\".zrchain.zenbtc.QueryParamsRequest\x1a#.zrchain.zenbtc.QueryParamsResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/zenbtc/params\x12\x92\x01\n\x10LockTransactions\x12,.zrchain.zenbtc.QueryLockTransactionsRequest\x1a-.zrchain.zenbtc.QueryLockTransactionsResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/zenbtc/lock_transactions\x12}\n\x0bRedemptions\x12\'.zrchain.zenbtc.QueryRedemptionsRequest\x1a(.zrchain.zenbtc.QueryRedemptionsResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/zenbtc/redemptionsB.Z,github.com/zenrocklabs/zenbtc/x/zenbtc/typesb\x06proto3')

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
  _globals['_QUERY'].methods_by_name['Params']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Params']._serialized_options = b'\202\323\344\223\002\020\022\016/zenbtc/params'
  _globals['_QUERY'].methods_by_name['LockTransactions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['LockTransactions']._serialized_options = b'\202\323\344\223\002\033\022\031/zenbtc/lock_transactions'
  _globals['_QUERY'].methods_by_name['Redemptions']._loaded_options = None
  _globals['_QUERY'].methods_by_name['Redemptions']._serialized_options = b'\202\323\344\223\002\025\022\023/zenbtc/redemptions'
  _globals['_COMPLETIONFILTER']._serialized_start=674
  _globals['_COMPLETIONFILTER']._serialized_end=729
  _globals['_QUERYPARAMSREQUEST']._serialized_start=224
  _globals['_QUERYPARAMSREQUEST']._serialized_end=244
  _globals['_QUERYPARAMSRESPONSE']._serialized_start=246
  _globals['_QUERYPARAMSRESPONSE']._serialized_end=326
  _globals['_QUERYLOCKTRANSACTIONSREQUEST']._serialized_start=328
  _globals['_QUERYLOCKTRANSACTIONSREQUEST']._serialized_end=358
  _globals['_QUERYLOCKTRANSACTIONSRESPONSE']._serialized_start=360
  _globals['_QUERYLOCKTRANSACTIONSRESPONSE']._serialized_end=436
  _globals['_QUERYREDEMPTIONSREQUEST']._serialized_start=439
  _globals['_QUERYREDEMPTIONSREQUEST']._serialized_end=576
  _globals['_QUERYREDEMPTIONSRESPONSE']._serialized_start=578
  _globals['_QUERYREDEMPTIONSRESPONSE']._serialized_end=672
  _globals['_QUERY']._serialized_start=732
  _globals['_QUERY']._serialized_end=1122
# @@protoc_insertion_point(module_scope)
