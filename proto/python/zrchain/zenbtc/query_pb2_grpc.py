# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from zrchain.zenbtc import query_pb2 as zrchain_dot_zenbtc_dot_query__pb2


class QueryStub(object):
    """Query defines the gRPC querier service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetParams = channel.unary_unary(
                '/zrchain.zenbtc.Query/GetParams',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsResponse.FromString,
                )
        self.LockTransactions = channel.unary_unary(
                '/zrchain.zenbtc.Query/LockTransactions',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsResponse.FromString,
                )
        self.GetRedemptions = channel.unary_unary(
                '/zrchain.zenbtc.Query/GetRedemptions',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsResponse.FromString,
                )
        self.QueryPendingMintTransactions = channel.unary_unary(
                '/zrchain.zenbtc.Query/QueryPendingMintTransactions',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryPendingMintTransactionsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryPendingMintTransactionsResponse.FromString,
                )
        self.QuerySupply = channel.unary_unary(
                '/zrchain.zenbtc.Query/QuerySupply',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QuerySupplyRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QuerySupplyResponse.FromString,
                )
        self.QueryBurnEvents = channel.unary_unary(
                '/zrchain.zenbtc.Query/QueryBurnEvents',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryBurnEventsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryBurnEventsResponse.FromString,
                )


class QueryServicer(object):
    """Query defines the gRPC querier service.
    """

    def GetParams(self, request, context):
        """Parameters queries the parameters of the module.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LockTransactions(self, request, context):
        """Queries a list of LockTransactions items.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRedemptions(self, request, context):
        """Queries a list of Redemptions items.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QueryPendingMintTransactions(self, request, context):
        """Queries a list of PendingMintTransactions items.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QuerySupply(self, request, context):
        """Queries the current supply of zenBTC.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QueryBurnEvents(self, request, context):
        """Queries a list of BurnEvents items.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_QueryServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetParams': grpc.unary_unary_rpc_method_handler(
                    servicer.GetParams,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsResponse.SerializeToString,
            ),
            'LockTransactions': grpc.unary_unary_rpc_method_handler(
                    servicer.LockTransactions,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsResponse.SerializeToString,
            ),
            'GetRedemptions': grpc.unary_unary_rpc_method_handler(
                    servicer.GetRedemptions,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsResponse.SerializeToString,
            ),
            'QueryPendingMintTransactions': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryPendingMintTransactions,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryPendingMintTransactionsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryPendingMintTransactionsResponse.SerializeToString,
            ),
            'QuerySupply': grpc.unary_unary_rpc_method_handler(
                    servicer.QuerySupply,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QuerySupplyRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QuerySupplyResponse.SerializeToString,
            ),
            'QueryBurnEvents': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryBurnEvents,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryBurnEventsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryBurnEventsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'zrchain.zenbtc.Query', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Query(object):
    """Query defines the gRPC querier service.
    """

    @staticmethod
    def GetParams(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/GetParams',
            zrchain_dot_zenbtc_dot_query__pb2.QueryParamsRequest.SerializeToString,
            zrchain_dot_zenbtc_dot_query__pb2.QueryParamsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def LockTransactions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/LockTransactions',
            zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsRequest.SerializeToString,
            zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRedemptions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/GetRedemptions',
            zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsRequest.SerializeToString,
            zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QueryPendingMintTransactions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/QueryPendingMintTransactions',
            zrchain_dot_zenbtc_dot_query__pb2.QueryPendingMintTransactionsRequest.SerializeToString,
            zrchain_dot_zenbtc_dot_query__pb2.QueryPendingMintTransactionsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QuerySupply(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/QuerySupply',
            zrchain_dot_zenbtc_dot_query__pb2.QuerySupplyRequest.SerializeToString,
            zrchain_dot_zenbtc_dot_query__pb2.QuerySupplyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QueryBurnEvents(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/QueryBurnEvents',
            zrchain_dot_zenbtc_dot_query__pb2.QueryBurnEventsRequest.SerializeToString,
            zrchain_dot_zenbtc_dot_query__pb2.QueryBurnEventsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
