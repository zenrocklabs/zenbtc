from amino import amino_pb2 as _amino_pb2
from cosmos.base.query.v1beta1 import pagination_pb2 as _pagination_pb2
from gogoproto import gogo_pb2 as _gogo_pb2
from google.api import annotations_pb2 as _annotations_pb2
from zrchain.zenbtc import params_pb2 as _params_pb2
from zrchain.zenbtc import redemptions_pb2 as _redemptions_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CompletionFilter(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    ALL: _ClassVar[CompletionFilter]
    PENDING: _ClassVar[CompletionFilter]
    COMPLETED: _ClassVar[CompletionFilter]
ALL: CompletionFilter
PENDING: CompletionFilter
COMPLETED: CompletionFilter

class QueryParamsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class QueryParamsResponse(_message.Message):
    __slots__ = ("params",)
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    params: _params_pb2.Params
    def __init__(self, params: _Optional[_Union[_params_pb2.Params, _Mapping]] = ...) -> None: ...

class QueryLockTransactionsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class QueryLockTransactionsResponse(_message.Message):
    __slots__ = ("lock_transactions",)
    LOCK_TRANSACTIONS_FIELD_NUMBER: _ClassVar[int]
    lock_transactions: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, lock_transactions: _Optional[_Iterable[str]] = ...) -> None: ...

class QueryRedemptionsRequest(_message.Message):
    __slots__ = ("start_index", "completion_filter")
    START_INDEX_FIELD_NUMBER: _ClassVar[int]
    COMPLETION_FILTER_FIELD_NUMBER: _ClassVar[int]
    start_index: int
    completion_filter: CompletionFilter
    def __init__(self, start_index: _Optional[int] = ..., completion_filter: _Optional[_Union[CompletionFilter, str]] = ...) -> None: ...

class QueryRedemptionsResponse(_message.Message):
    __slots__ = ("redemptions",)
    REDEMPTIONS_FIELD_NUMBER: _ClassVar[int]
    redemptions: _containers.RepeatedCompositeFieldContainer[_redemptions_pb2.Redemption]
    def __init__(self, redemptions: _Optional[_Iterable[_Union[_redemptions_pb2.Redemption, _Mapping]]] = ...) -> None: ...
