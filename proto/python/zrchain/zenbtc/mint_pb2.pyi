from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class NonceData(_message.Message):
    __slots__ = ("nonce", "counter")
    NONCE_FIELD_NUMBER: _ClassVar[int]
    COUNTER_FIELD_NUMBER: _ClassVar[int]
    nonce: int
    counter: int
    def __init__(self, nonce: _Optional[int] = ..., counter: _Optional[int] = ...) -> None: ...
