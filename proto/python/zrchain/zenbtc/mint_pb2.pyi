from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class NonceData(_message.Message):
    __slots__ = ("nonce", "counter")
    NONCE_FIELD_NUMBER: _ClassVar[int]
    COUNTER_FIELD_NUMBER: _ClassVar[int]
    nonce: int
    counter: int
    def __init__(self, nonce: _Optional[int] = ..., counter: _Optional[int] = ...) -> None: ...

class RequestedBitcoinHeaders(_message.Message):
    __slots__ = ("heights",)
    HEIGHTS_FIELD_NUMBER: _ClassVar[int]
    heights: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, heights: _Optional[_Iterable[int]] = ...) -> None: ...
