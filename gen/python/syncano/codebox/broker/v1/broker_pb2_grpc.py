# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from syncano.codebox.broker.v1 import broker_pb2 as syncano_dot_codebox_dot_broker_dot_v1_dot_broker__pb2
from syncano.codebox.script.v1 import script_pb2 as syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2


class ScriptRunnerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Run = channel.stream_stream(
                '/syncano.codebox.broker.v1.ScriptRunner/Run',
                request_serializer=syncano_dot_codebox_dot_broker_dot_v1_dot_broker__pb2.RunRequest.SerializeToString,
                response_deserializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.RunResponse.FromString,
                )
        self.SimpleRun = channel.unary_stream(
                '/syncano.codebox.broker.v1.ScriptRunner/SimpleRun',
                request_serializer=syncano_dot_codebox_dot_broker_dot_v1_dot_broker__pb2.SimpleRunRequest.SerializeToString,
                response_deserializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.RunResponse.FromString,
                )
        self.Delete = channel.unary_unary(
                '/syncano.codebox.broker.v1.ScriptRunner/Delete',
                request_serializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.DeleteRequest.SerializeToString,
                response_deserializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.DeleteResponse.FromString,
                )


class ScriptRunnerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Run(self, request_iterator, context):
        """Run runs script in secure environment.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SimpleRun(self, request, context):
        """SimpleRun is a simpler alternative to Run that does not require streaming request.
        As such, it is only usable for small payloads and does not support chunks.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Delete deletes all containers for specified script index.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ScriptRunnerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Run': grpc.stream_stream_rpc_method_handler(
                    servicer.Run,
                    request_deserializer=syncano_dot_codebox_dot_broker_dot_v1_dot_broker__pb2.RunRequest.FromString,
                    response_serializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.RunResponse.SerializeToString,
            ),
            'SimpleRun': grpc.unary_stream_rpc_method_handler(
                    servicer.SimpleRun,
                    request_deserializer=syncano_dot_codebox_dot_broker_dot_v1_dot_broker__pb2.SimpleRunRequest.FromString,
                    response_serializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.RunResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.DeleteRequest.FromString,
                    response_serializer=syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.DeleteResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'syncano.codebox.broker.v1.ScriptRunner', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ScriptRunner(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Run(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/syncano.codebox.broker.v1.ScriptRunner/Run',
            syncano_dot_codebox_dot_broker_dot_v1_dot_broker__pb2.RunRequest.SerializeToString,
            syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.RunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SimpleRun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/syncano.codebox.broker.v1.ScriptRunner/SimpleRun',
            syncano_dot_codebox_dot_broker_dot_v1_dot_broker__pb2.SimpleRunRequest.SerializeToString,
            syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.RunResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/syncano.codebox.broker.v1.ScriptRunner/Delete',
            syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.DeleteRequest.SerializeToString,
            syncano_dot_codebox_dot_script_dot_v1_dot_script__pb2.DeleteResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
