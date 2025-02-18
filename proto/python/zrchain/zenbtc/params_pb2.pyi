from amino import amino_pb2 as _amino_pb2
from gogoproto import gogo_pb2 as _gogo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Params(_message.Message):
    __slots__ = ("depositKeyringAddr", "stakerKeyID", "ethMinterKeyID", "unstakerKeyID", "completerKeyID", "rewardsDepositKeyID", "changeAddressKeyIDs", "bitcoinProxyAddress", "ethTokenAddr", "controllerAddr")
    DEPOSITKEYRINGADDR_FIELD_NUMBER: _ClassVar[int]
    STAKERKEYID_FIELD_NUMBER: _ClassVar[int]
    ETHMINTERKEYID_FIELD_NUMBER: _ClassVar[int]
    UNSTAKERKEYID_FIELD_NUMBER: _ClassVar[int]
    COMPLETERKEYID_FIELD_NUMBER: _ClassVar[int]
    REWARDSDEPOSITKEYID_FIELD_NUMBER: _ClassVar[int]
    CHANGEADDRESSKEYIDS_FIELD_NUMBER: _ClassVar[int]
    BITCOINPROXYADDRESS_FIELD_NUMBER: _ClassVar[int]
    ETHTOKENADDR_FIELD_NUMBER: _ClassVar[int]
    CONTROLLERADDR_FIELD_NUMBER: _ClassVar[int]
    depositKeyringAddr: str
    stakerKeyID: int
    ethMinterKeyID: int
    unstakerKeyID: int
    completerKeyID: int
    rewardsDepositKeyID: int
    changeAddressKeyIDs: _containers.RepeatedScalarFieldContainer[int]
    bitcoinProxyAddress: str
    ethTokenAddr: str
    controllerAddr: str
    def __init__(self, depositKeyringAddr: _Optional[str] = ..., stakerKeyID: _Optional[int] = ..., ethMinterKeyID: _Optional[int] = ..., unstakerKeyID: _Optional[int] = ..., completerKeyID: _Optional[int] = ..., rewardsDepositKeyID: _Optional[int] = ..., changeAddressKeyIDs: _Optional[_Iterable[int]] = ..., bitcoinProxyAddress: _Optional[str] = ..., ethTokenAddr: _Optional[str] = ..., controllerAddr: _Optional[str] = ...) -> None: ...
