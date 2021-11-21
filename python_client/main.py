import logging

import grpc
import proto.go_server.calculator_pb2
import proto.go_server.calculator_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = calculator_grpc.CalculatorServiceStub(channel)
        response = stub.Add(calculator_pb2.AddRequest(a=1, b=2, c=3))
    print("Greeter client received: " + response.message)


if __name__ == '__main__':
    logging.basicConfig()
    run()
