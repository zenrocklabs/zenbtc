# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from zrchain.zenbtc import tx_pb2 as zrchain_dot_zenbtc_dot_tx__pb2


class MsgStub(object):
    """Msg defines the Msg service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UpdateParams = channel.unary_unary(
                '/zrchain.zenbtc.Msg/UpdateParams',
                request_serializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgUpdateParams.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgUpdateParamsResponse.FromString,
                )
        self.VerifyDepositBlockInclusion = channel.unary_unary(
                '/zrchain.zenbtc.Msg/VerifyDepositBlockInclusion',
                request_serializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgVerifyDepositBlockInclusion.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgVerifyDepositBlockInclusionResponse.FromString,
                )
        self.SubmitUnsignedRedemptionTx = channel.unary_unary(
                '/zrchain.zenbtc.Msg/SubmitUnsignedRedemptionTx',
                request_serializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgSubmitUnsignedRedemptionTx.SerializeToString,
                response_deserializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgSubmitUnsignedRedemptionTxResponse.FromString,
                )


class MsgServicer(object):
    """Msg defines the Msg service.
    """

    def UpdateParams(self, request, context):
        """UpdateParams defines a (governance) operation for updating the module
        parameters. The authority defaults to the x/gov module account.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def VerifyDepositBlockInclusion(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubmitUnsignedRedemptionTx(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MsgServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UpdateParams': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateParams,
                    request_deserializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgUpdateParams.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgUpdateParamsResponse.SerializeToString,
            ),
            'VerifyDepositBlockInclusion': grpc.unary_unary_rpc_method_handler(
                    servicer.VerifyDepositBlockInclusion,
                    request_deserializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgVerifyDepositBlockInclusion.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgVerifyDepositBlockInclusionResponse.SerializeToString,
            ),
            'SubmitUnsignedRedemptionTx': grpc.unary_unary_rpc_method_handler(
                    servicer.SubmitUnsignedRedemptionTx,
                    request_deserializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgSubmitUnsignedRedemptionTx.FromString,
                    response_serializer=zrchain_dot_zenbtc_dot_tx__pb2.MsgSubmitUnsignedRedemptionTxResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'zrchain.zenbtc.Msg', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Msg(object):
    """Msg defines the Msg service.
    """

    @staticmethod
    def UpdateParams(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Msg/UpdateParams',
            zrchain_dot_zenbtc_dot_tx__pb2.MsgUpdateParams.SerializeToString,
            zrchain_dot_zenbtc_dot_tx__pb2.MsgUpdateParamsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def VerifyDepositBlockInclusion(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Msg/VerifyDepositBlockInclusion',
            zrchain_dot_zenbtc_dot_tx__pb2.MsgVerifyDepositBlockInclusion.SerializeToString,
            zrchain_dot_zenbtc_dot_tx__pb2.MsgVerifyDepositBlockInclusionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SubmitUnsignedRedemptionTx(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/zrchain.zenbtc.Msg/SubmitUnsignedRedemptionTx',
            zrchain_dot_zenbtc_dot_tx__pb2.MsgSubmitUnsignedRedemptionTx.SerializeToString,
            zrchain_dot_zenbtc_dot_tx__pb2.MsgSubmitUnsignedRedemptionTxResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
