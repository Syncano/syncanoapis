# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from syncano.codebox.filerepo.v1 import repo_pb2 as syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2


class RepoStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Exists = channel.unary_unary(
                '/syncano.codebox.filerepo.v1.Repo/Exists',
                request_serializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.ExistsRequest.SerializeToString,
                response_deserializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.ExistsResponse.FromString,
                )
        self.Upload = channel.stream_stream(
                '/syncano.codebox.filerepo.v1.Repo/Upload',
                request_serializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.UploadRequest.SerializeToString,
                response_deserializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.UploadResponse.FromString,
                )


class RepoServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Exists(self, request, context):
        """Exists checks if file was defined in file repo.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Upload(self, request_iterator, context):
        """Upload streams file(s) to server.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RepoServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Exists': grpc.unary_unary_rpc_method_handler(
                    servicer.Exists,
                    request_deserializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.ExistsRequest.FromString,
                    response_serializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.ExistsResponse.SerializeToString,
            ),
            'Upload': grpc.stream_stream_rpc_method_handler(
                    servicer.Upload,
                    request_deserializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.UploadRequest.FromString,
                    response_serializer=syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.UploadResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'syncano.codebox.filerepo.v1.Repo', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Repo(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Exists(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/syncano.codebox.filerepo.v1.Repo/Exists',
            syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.ExistsRequest.SerializeToString,
            syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.ExistsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Upload(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/syncano.codebox.filerepo.v1.Repo/Upload',
            syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.UploadRequest.SerializeToString,
            syncano_dot_codebox_dot_filerepo_dot_v1_dot_repo__pb2.UploadResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
