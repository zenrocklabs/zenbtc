# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/zenbtc/redemptions.proto
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
    'zrchain/zenbtc/redemptions.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n zrchain/zenbtc/redemptions.proto\x12\x0ezrchain.zenbtc\x1a\x14gogoproto/gogo.proto\"\x80\x01\n\nRedemption\x12\x38\n\x04\x64\x61ta\x18\x01 \x01(\x0b\x32\x1e.zrchain.zenbtc.RedemptionDataB\x04\xc8\xde\x1f\x00R\x04\x64\x61ta\x12\x38\n\x06status\x18\x02 \x01(\x0e\x32 .zrchain.zenbtc.RedemptionStatusR\x06status\"i\n\x0eRedemptionData\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12/\n\x13\x64\x65stination_address\x18\x02 \x01(\x0cR\x12\x64\x65stinationAddress\x12\x16\n\x06\x61mount\x18\x03 \x01(\x04R\x06\x61mount\"\x97\x01\n\tBurnEvent\x12\x12\n\x04txID\x18\x01 \x01(\tR\x04txID\x12\x1a\n\x08logIndex\x18\x02 \x01(\x04R\x08logIndex\x12\x18\n\x07\x63hainID\x18\x03 \x01(\tR\x07\x63hainID\x12(\n\x0f\x64\x65stinationAddr\x18\x04 \x01(\x0cR\x0f\x64\x65stinationAddr\x12\x16\n\x06\x61mount\x18\x05 \x01(\x04R\x06\x61mount\"?\n\nBurnEvents\x12\x31\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x19.zrchain.zenbtc.BurnEventR\x06\x65vents*[\n\x10RedemptionStatus\x12\x0f\n\x0bUNSPECIFIED\x10\x00\x12\r\n\tINITIATED\x10\x01\x12\x0c\n\x08UNSTAKED\x10\x02\x12\n\n\x06\x42URNED\x10\x03\x12\r\n\tCOMPLETED\x10\x04\x42.Z,github.com/zenrocklabs/zenbtc/x/zenbtc/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.zenbtc.redemptions_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,github.com/zenrocklabs/zenbtc/x/zenbtc/types'
  _globals['_REDEMPTION'].fields_by_name['data']._loaded_options = None
  _globals['_REDEMPTION'].fields_by_name['data']._serialized_options = b'\310\336\037\000'
  _globals['_REDEMPTIONSTATUS']._serialized_start=531
  _globals['_REDEMPTIONSTATUS']._serialized_end=622
  _globals['_REDEMPTION']._serialized_start=75
  _globals['_REDEMPTION']._serialized_end=203
  _globals['_REDEMPTIONDATA']._serialized_start=205
  _globals['_REDEMPTIONDATA']._serialized_end=310
  _globals['_BURNEVENT']._serialized_start=313
  _globals['_BURNEVENT']._serialized_end=464
  _globals['_BURNEVENTS']._serialized_start=466
  _globals['_BURNEVENTS']._serialized_end=529
# @@protoc_insertion_point(module_scope)
