from amino import amino_pb2 as _amino_pb2
from gogoproto import gogo_pb2 as _gogo_pb2
from zrchain.zenbtc import params_pb2 as _params_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GenesisState(_message.Message):
    __slots__ = ("params", "no_fee_msgs")
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    NO_FEE_MSGS_FIELD_NUMBER: _ClassVar[int]
    params: _params_pb2.Params
    no_fee_msgs: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, params: _Optional[_Union[_params_pb2.Params, _Mapping]] = ..., no_fee_msgs: _Optional[_Iterable[str]] = ...) -> None: ...
