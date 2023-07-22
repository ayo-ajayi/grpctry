from concurrent import futures
import grpc
import signal

import os
import sys
sys.path.append(os.path.join(os.getcwd(), 'pygen'))
import grpctry_pb2_grpc as rpc
import grpctry_pb2 as pb

#create a servicer
class GreeterServer(rpc.GreeterServicer):
    def __init__(self):
        self.i=0
    def SayHello(self, request, context):
        i=self.i
        while i<200:
            i=1+i
        return pb.HelloReply(message="success! here is response for request: "+request.name+"\n i is: "+str(i))


def server():
    port = "50051"
    s = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    rpc.add_GreeterServicer_to_server(GreeterServer(), s)
    s.add_insecure_port("localhost:" + port)

    
    def handle_sigint(signum, frame):
        print("\nReceived SIGINT (Ctrl+C), shutting down gracefully...")
        s.stop(0)  # Gracefully stop the server with a timeout of 0 (no waiting)

    signal.signal(signal.SIGINT, handle_sigint)

    s.start()
    print("server listening on port: " + port)
    s.wait_for_termination()

if __name__ == "__main__":
    server()

