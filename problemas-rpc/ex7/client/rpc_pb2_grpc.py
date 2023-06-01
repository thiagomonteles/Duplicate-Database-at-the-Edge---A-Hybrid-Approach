# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import rpc_pb2 as rpc__pb2


class AposentadoriaStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.PodeAposentar = channel.unary_unary(
                '/aposentadoria.Aposentadoria/PodeAposentar',
                request_serializer=rpc__pb2.PodeAposentarRequest.SerializeToString,
                response_deserializer=rpc__pb2.PodeAposentarReply.FromString,
                )


class AposentadoriaServicer(object):
    """Missing associated documentation comment in .proto file."""

    def PodeAposentar(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AposentadoriaServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'PodeAposentar': grpc.unary_unary_rpc_method_handler(
                    servicer.PodeAposentar,
                    request_deserializer=rpc__pb2.PodeAposentarRequest.FromString,
                    response_serializer=rpc__pb2.PodeAposentarReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'aposentadoria.Aposentadoria', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Aposentadoria(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def PodeAposentar(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/aposentadoria.Aposentadoria/PodeAposentar',
            rpc__pb2.PodeAposentarRequest.SerializeToString,
            rpc__pb2.PodeAposentarReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)