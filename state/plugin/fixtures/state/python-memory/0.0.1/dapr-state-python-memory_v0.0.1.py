from concurrent import futures
import sys
import time

import grpc

from proto import state_pb2_grpc
from proto import state_pb2

from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc

class StateServicer(state_pb2_grpc.StateServicer):
    """Implementation of State service."""
    data = {}

    def Init(self, metadata, context):
        self.data = {}
        return state_pb2.Empty()

    def Get(self, request, context):
        result = self.data[request.key]
        response = state_pb2.GetResponse()
        response.data = result
        return response

    def Set(self, request, context):
        self.data[request.key] = request.value
        return state_pb2.Empty()

    def Delete(self, request, context):
        self.data.pop('key', None)
        return state_pb2.Empty()

def serve():
    # We need to build a health service to work with go-plugin
    health = HealthServicer()
    health.set("plugin", health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))

    # Start the server.
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    state_pb2_grpc.add_StateServicer_to_server(StateServicer(), server)
    health_pb2_grpc.add_HealthServicer_to_server(health, server)
    
    port = server.add_insecure_port('127.0.0.1:0')
    server.start()
    # Output information for the plugin client
    print(f'1|1|tcp|127.0.0.1:{port:d}|grpc')
    sys.stdout.flush()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
