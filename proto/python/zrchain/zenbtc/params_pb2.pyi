from amino import amino_pb2 as _amino_pb2
from gogoproto import gogo_pb2 as _gogo_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Params(_message.Message):
    __slots__ = ("ethBatcherAddr", "depositKeyringAddr", "minterKeyID", "withdrawerKeyID", "unstakerKeyID", "rewardsDepositKeyID", "changeAddressKeyIDs", "bitcoinProxyAddress", "authority", "stakerKeyID", "completerKeyID", "ethMinterKeyID", "ethTokenAddr")
    ETHBATCHERADDR_FIELD_NUMBER: _ClassVar[int]
    DEPOSITKEYRINGADDR_FIELD_NUMBER: _ClassVar[int]
    MINTERKEYID_FIELD_NUMBER: _ClassVar[int]
    WITHDRAWERKEYID_FIELD_NUMBER: _ClassVar[int]
    UNSTAKERKEYID_FIELD_NUMBER: _ClassVar[int]
    REWARDSDEPOSITKEYID_FIELD_NUMBER: _ClassVar[int]
    CHANGEADDRESSKEYIDS_FIELD_NUMBER: _ClassVar[int]
    BITCOINPROXYADDRESS_FIELD_NUMBER: _ClassVar[int]
    AUTHORITY_FIELD_NUMBER: _ClassVar[int]
    STAKERKEYID_FIELD_NUMBER: _ClassVar[int]
    COMPLETERKEYID_FIELD_NUMBER: _ClassVar[int]
    ETHMINTERKEYID_FIELD_NUMBER: _ClassVar[int]
    ETHTOKENADDR_FIELD_NUMBER: _ClassVar[int]
    ethBatcherAddr: str
    depositKeyringAddr: str
    minterKeyID: int
    withdrawerKeyID: int
    unstakerKeyID: int
    rewardsDepositKeyID: int
    changeAddressKeyIDs: _containers.RepeatedScalarFieldContainer[int]
    bitcoinProxyAddress: str
    authority: str
    stakerKeyID: int
    completerKeyID: int
    ethMinterKeyID: int
    ethTokenAddr: str
    def __init__(self, ethBatcherAddr: _Optional[str] = ..., depositKeyringAddr: _Optional[str] = ..., minterKeyID: _Optional[int] = ..., withdrawerKeyID: _Optional[int] = ..., unstakerKeyID: _Optional[int] = ..., rewardsDepositKeyID: _Optional[int] = ..., changeAddressKeyIDs: _Optional[_Iterable[int]] = ..., bitcoinProxyAddress: _Optional[str] = ..., authority: _Optional[str] = ..., stakerKeyID: _Optional[int] = ..., completerKeyID: _Optional[int] = ..., ethMinterKeyID: _Optional[int] = ..., ethTokenAddr: _Optional[str] = ...) -> None: ...
