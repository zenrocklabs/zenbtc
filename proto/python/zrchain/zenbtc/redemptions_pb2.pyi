from gogoproto import gogo_pb2 as _gogo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RedemptionStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    UNSPECIFIED: _ClassVar[RedemptionStatus]
    INITIATED: _ClassVar[RedemptionStatus]
    UNSTAKED: _ClassVar[RedemptionStatus]
    BURNED: _ClassVar[RedemptionStatus]
    COMPLETED: _ClassVar[RedemptionStatus]
UNSPECIFIED: RedemptionStatus
INITIATED: RedemptionStatus
UNSTAKED: RedemptionStatus
BURNED: RedemptionStatus
COMPLETED: RedemptionStatus

class Redemption(_message.Message):
    __slots__ = ("data", "status")
    DATA_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    data: RedemptionData
    status: RedemptionStatus
    def __init__(self, data: _Optional[_Union[RedemptionData, _Mapping]] = ..., status: _Optional[_Union[RedemptionStatus, str]] = ...) -> None: ...

class RedemptionData(_message.Message):
    __slots__ = ("id", "destination_address", "amount")
    ID_FIELD_NUMBER: _ClassVar[int]
    DESTINATION_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    id: int
    destination_address: bytes
    amount: int
    def __init__(self, id: _Optional[int] = ..., destination_address: _Optional[bytes] = ..., amount: _Optional[int] = ...) -> None: ...

class BurnEvent(_message.Message):
    __slots__ = ("txID", "logIndex", "chainID", "destinationAddr", "amount")
    TXID_FIELD_NUMBER: _ClassVar[int]
    LOGINDEX_FIELD_NUMBER: _ClassVar[int]
    CHAINID_FIELD_NUMBER: _ClassVar[int]
    DESTINATIONADDR_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    txID: str
    logIndex: int
    chainID: str
    destinationAddr: str
    amount: int
    def __init__(self, txID: _Optional[str] = ..., logIndex: _Optional[int] = ..., chainID: _Optional[str] = ..., destinationAddr: _Optional[str] = ..., amount: _Optional[int] = ...) -> None: ...

class BurnEvents(_message.Message):
    __slots__ = ("events",)
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.RepeatedCompositeFieldContainer[BurnEvent]
    def __init__(self, events: _Optional[_Iterable[_Union[BurnEvent, _Mapping]]] = ...) -> None: ...
