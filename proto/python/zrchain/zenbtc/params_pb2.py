# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: zrchain/zenbtc/params.proto
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
    'zrchain/zenbtc/params.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from amino import amino_pb2 as amino_dot_amino__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bzrchain/zenbtc/params.proto\x12\x0ezrchain.zenbtc\x1a\x11\x61mino/amino.proto\x1a\x14gogoproto/gogo.proto\"\xac\x03\n\x06Params\x12&\n\x0e\x65thBatcherAddr\x18\x01 \x01(\tR\x0e\x65thBatcherAddr\x12.\n\x12\x64\x65positKeyringAddr\x18\x02 \x01(\tR\x12\x64\x65positKeyringAddr\x12 \n\x0bminterKeyID\x18\x03 \x01(\x04R\x0bminterKeyID\x12(\n\x0fwithdrawerKeyID\x18\x04 \x01(\x04R\x0fwithdrawerKeyID\x12$\n\runstakerKeyID\x18\x05 \x01(\x04R\runstakerKeyID\x12\x30\n\x13rewardsDepositKeyID\x18\x06 \x01(\x04R\x13rewardsDepositKeyID\x12\x30\n\x13\x63hangeAddressKeyIDs\x18\x07 \x03(\x04R\x13\x63hangeAddressKeyIDs\x12\x34\n\x15\x62itcoinProxyCreatorID\x18\x08 \x01(\tR\x15\x62itcoinProxyCreatorID\x12\x1c\n\tauthority\x18\t \x01(\tR\tauthority: \xe8\xa0\x1f\x01\x8a\xe7\xb0*\x17zrchain/x/zenbtc/ParamsB.Z,github.com/zenrocklabs/zenbtc/x/zenbtc/typesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'zrchain.zenbtc.params_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z,github.com/zenrocklabs/zenbtc/x/zenbtc/types'
  _globals['_PARAMS']._loaded_options = None
  _globals['_PARAMS']._serialized_options = b'\350\240\037\001\212\347\260*\027zrchain/x/zenbtc/Params'
  _globals['_PARAMS']._serialized_start=89
  _globals['_PARAMS']._serialized_end=517
# @@protoc_insertion_point(module_scope)
