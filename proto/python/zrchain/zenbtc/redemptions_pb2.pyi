from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Redemption(_message.Message):
    __slots__ = ("recipient_address", "amount")
    RECIPIENT_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    recipient_address: str
    amount: int
    def __init__(self, recipient_address: _Optional[str] = ..., amount: _Optional[int] = ...) -> None: ...
