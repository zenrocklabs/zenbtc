from amino import amino_pb2 as _amino_pb2
from gogoproto import gogo_pb2 as _gogo_pb2
from zrchain.zenbtc import params_pb2 as _params_pb2
from zrchain.zenbtc import supply_pb2 as _supply_pb2
from zrchain.zenbtc import redemptions_pb2 as _redemptions_pb2
from zrchain.zenbtc import mint_pb2 as _mint_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GenesisState(_message.Message):
    __slots__ = ("params", "lock_transactions", "pending_mint_transactions", "first_pending_eth_mint_transaction", "first_pending_sol_mint_transaction", "pending_mint_transaction_count", "burn_events", "first_pending_burn_event", "burn_event_count", "redemptions", "first_pending_redemption", "first_redemption_awaiting_sign", "supply", "first_pending_stake_transaction")
    PARAMS_FIELD_NUMBER: _ClassVar[int]
    LOCK_TRANSACTIONS_FIELD_NUMBER: _ClassVar[int]
    PENDING_MINT_TRANSACTIONS_FIELD_NUMBER: _ClassVar[int]
    FIRST_PENDING_ETH_MINT_TRANSACTION_FIELD_NUMBER: _ClassVar[int]
    FIRST_PENDING_SOL_MINT_TRANSACTION_FIELD_NUMBER: _ClassVar[int]
    PENDING_MINT_TRANSACTION_COUNT_FIELD_NUMBER: _ClassVar[int]
    BURN_EVENTS_FIELD_NUMBER: _ClassVar[int]
    FIRST_PENDING_BURN_EVENT_FIELD_NUMBER: _ClassVar[int]
    BURN_EVENT_COUNT_FIELD_NUMBER: _ClassVar[int]
    REDEMPTIONS_FIELD_NUMBER: _ClassVar[int]
    FIRST_PENDING_REDEMPTION_FIELD_NUMBER: _ClassVar[int]
    FIRST_REDEMPTION_AWAITING_SIGN_FIELD_NUMBER: _ClassVar[int]
    SUPPLY_FIELD_NUMBER: _ClassVar[int]
    FIRST_PENDING_STAKE_TRANSACTION_FIELD_NUMBER: _ClassVar[int]
    params: _params_pb2.Params
    lock_transactions: _containers.RepeatedCompositeFieldContainer[_mint_pb2.LockTransaction]
    pending_mint_transactions: _containers.RepeatedCompositeFieldContainer[_mint_pb2.PendingMintTransaction]
    first_pending_eth_mint_transaction: int
    first_pending_sol_mint_transaction: int
    pending_mint_transaction_count: int
    burn_events: _containers.RepeatedCompositeFieldContainer[_redemptions_pb2.BurnEvent]
    first_pending_burn_event: int
    burn_event_count: int
    redemptions: _containers.RepeatedCompositeFieldContainer[_redemptions_pb2.Redemption]
    first_pending_redemption: int
    first_redemption_awaiting_sign: int
    supply: _supply_pb2.Supply
    first_pending_stake_transaction: int
    def __init__(self, params: _Optional[_Union[_params_pb2.Params, _Mapping]] = ..., lock_transactions: _Optional[_Iterable[_Union[_mint_pb2.LockTransaction, _Mapping]]] = ..., pending_mint_transactions: _Optional[_Iterable[_Union[_mint_pb2.PendingMintTransaction, _Mapping]]] = ..., first_pending_eth_mint_transaction: _Optional[int] = ..., first_pending_sol_mint_transaction: _Optional[int] = ..., pending_mint_transaction_count: _Optional[int] = ..., burn_events: _Optional[_Iterable[_Union[_redemptions_pb2.BurnEvent, _Mapping]]] = ..., first_pending_burn_event: _Optional[int] = ..., burn_event_count: _Optional[int] = ..., redemptions: _Optional[_Iterable[_Union[_redemptions_pb2.Redemption, _Mapping]]] = ..., first_pending_redemption: _Optional[int] = ..., first_redemption_awaiting_sign: _Optional[int] = ..., supply: _Optional[_Union[_supply_pb2.Supply, _Mapping]] = ..., first_pending_stake_transaction: _Optional[int] = ...) -> None: ...
