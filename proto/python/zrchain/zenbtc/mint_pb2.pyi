from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class NonceData(_message.Message):
    __slots__ = ("nonce", "counter", "skip")
    NONCE_FIELD_NUMBER: _ClassVar[int]
    COUNTER_FIELD_NUMBER: _ClassVar[int]
    SKIP_FIELD_NUMBER: _ClassVar[int]
    nonce: int
    counter: int
    skip: bool
    def __init__(self, nonce: _Optional[int] = ..., counter: _Optional[int] = ..., skip: bool = ...) -> None: ...

class RequestedBitcoinHeaders(_message.Message):
    __slots__ = ("heights",)
    HEIGHTS_FIELD_NUMBER: _ClassVar[int]
    heights: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, heights: _Optional[_Iterable[int]] = ...) -> None: ...

class LockTransaction(_message.Message):
    __slots__ = ("raw_tx", "vout", "sender", "mint_recipient", "amount", "block_height")
    RAW_TX_FIELD_NUMBER: _ClassVar[int]
    VOUT_FIELD_NUMBER: _ClassVar[int]
    SENDER_FIELD_NUMBER: _ClassVar[int]
    MINT_RECIPIENT_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    BLOCK_HEIGHT_FIELD_NUMBER: _ClassVar[int]
    raw_tx: str
    vout: int
    sender: str
    mint_recipient: str
    amount: int
    block_height: int
    def __init__(self, raw_tx: _Optional[str] = ..., vout: _Optional[int] = ..., sender: _Optional[str] = ..., mint_recipient: _Optional[str] = ..., amount: _Optional[int] = ..., block_height: _Optional[int] = ...) -> None: ...
