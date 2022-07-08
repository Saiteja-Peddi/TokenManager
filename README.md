gRPC Prroject-3

Changes made to the old token manager:
1. Modified Final value storage. Earlier hash value is stored now x is stored.
2. Combined both Read and Write methods  functionality and directly storing the final value.
3. Now Read function returns the final value instead of performing mid to high operations.


Follow the below steps to run the Project-3

1. Open 3 different terminals in AOS_Prj3 directory and run each of the following commands in each terminal.
    go run ./tokenServer/main.go
    go run ./server-1/main.go
    go run ./server-2/main.go

2. Open another new Terminal window in the directory of AOS_PRJ2 and run below bash command.
    bash prj3.sh
 

PS: If any dependency issues appear in any of protobuf files, run below command.
        go mod tidy
