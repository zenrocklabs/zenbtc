from gogoproto import gogo_pb2 as _gogo_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Redemption(_message.Message):
    __slots__ = ("data", "completed")
    DATA_FIELD_NUMBER: _ClassVar[int]
    COMPLETED_FIELD_NUMBER: _ClassVar[int]
    data: RedemptionData
    completed: bool
    def __init__(self, data: _Optional[_Union[RedemptionData, _Mapping]] = ..., completed: bool = ...) -> None: ...

class RedemptionData(_message.Message):
    __slots__ = ("id", "destination_address", "amount")
    ID_FIELD_NUMBER: _ClassVar[int]
    DESTINATION_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    id: int
    destination_address: bytes
    amount: int
    def __init__(self, id: _Optional[int] = ..., destination_address: _Optional[bytes] = ..., amount: _Optional[int] = ...) -> None: ...
