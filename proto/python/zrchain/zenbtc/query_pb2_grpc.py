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
        self.Params = channel.unary_unary(
                '/zrchain.zenbtc.Query/Params',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsResponse.FromString,
                )
        self.LockTransactions = channel.unary_unary(
                '/zrchain.zenbtc.Query/LockTransactions',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsResponse.FromString,
                )
        self.Redemptions = channel.unary_unary(
                '/zrchain.zenbtc.Query/Redemptions',
                request_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsRequest.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsResponse.FromString,
                )


class QueryServicer(object):
    """Query defines the gRPC querier service.
    """

    def Params(self, request, context):
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

    def Redemptions(self, request, context):
        """Queries a list of Redemptions items.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_QueryServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Params': grpc.unary_unary_rpc_method_handler(
                    servicer.Params,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryParamsResponse.SerializeToString,
            ),
            'LockTransactions': grpc.unary_unary_rpc_method_handler(
                    servicer.LockTransactions,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryLockTransactionsResponse.SerializeToString,
            ),
            'Redemptions': grpc.unary_unary_rpc_method_handler(
                    servicer.Redemptions,
                    request_deserializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsRequest.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsResponse.SerializeToString,
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
    def Params(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/Params',
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
    def Redemptions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Query/Redemptions',
            zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsRequest.SerializeToString,
            zrchain_dot_zenbtc_dot_query__pb2.QueryRedemptionsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
