

## TYpes of gRPC

1. Unary: single req to single resp
2. Server streaming RPC: Single req -> multiple response from server.
3. CLient streaming RPC: multiple client streaming->  getting a single server response.
4. Bi Directional streaming RPC: multiple req and response from both sides in any order.


protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto
