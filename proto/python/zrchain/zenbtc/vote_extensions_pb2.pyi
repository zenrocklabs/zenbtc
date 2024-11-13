from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Network(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ETHEREUM: _ClassVar[Network]
    SOLANA: _ClassVar[Network]
ETHEREUM: Network
SOLANA: Network

class Redemption(_message.Message):
    __slots__ = ("recipient_address", "amount", "origin_chain", "block_height", "tx_index")
    RECIPIENT_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    ORIGIN_CHAIN_FIELD_NUMBER: _ClassVar[int]
    BLOCK_HEIGHT_FIELD_NUMBER: _ClassVar[int]
    TX_INDEX_FIELD_NUMBER: _ClassVar[int]
    recipient_address: str
    amount: int
    origin_chain: Network
    block_height: int
    tx_index: int
    def __init__(self, recipient_address: _Optional[str] = ..., amount: _Optional[int] = ..., origin_chain: _Optional[_Union[Network, str]] = ..., block_height: _Optional[int] = ..., tx_index: _Optional[int] = ...) -> None: ...

class Redemptions(_message.Message):
    __slots__ = ("redemptions",)
    REDEMPTIONS_FIELD_NUMBER: _ClassVar[int]
    redemptions: _containers.RepeatedCompositeFieldContainer[Redemption]
    def __init__(self, redemptions: _Optional[_Iterable[_Union[Redemption, _Mapping]]] = ...) -> None: ...
