#!/bin/bash
SERVERMAIN=./tokenServer/main.go
SERVER1=./server-1/main.go
SERVER2=./server-1/main.go
CLIENTMAIN=./tokenClient/main.go

#Running Server
# go run $SERVERMAIN -port 4200
# wait
# go run $SERVER1 -port 4201
# wait
# go run $SERVER2 -port 4202
# wait
#Running Client


go run $CLIENTMAIN -write -id 2 -name teja -low 5 -mid 10 -high 50 -host localhost -port 4201 
go run $CLIENTMAIN -write -id 1 -name teja -low 5 -mid 10 -high 50 -host localhost -port 4200 

go run $CLIENTMAIN -write -id 2 -name peddi -low 50 -mid 70 -high 100 -host localhost -port 4201 &
go run $CLIENTMAIN -read -id 2 -host localhost -port 4202 &
go run $CLIENTMAIN -write -id 2 -name pedsaidi -low 70 -mid 80 -high 100 -host localhost -port 4201 &
go run $CLIENTMAIN -read -id 2 -host localhost -port 4200

go run $CLIENTMAIN -write -id 1 -name abcd -low 50 -mid 70 -high 100 -host localhost -port 4200 &
go run $CLIENTMAIN -read -id 1 -host localhost -port 4202 &
go run $CLIENTMAIN -write -id 1 -name efghi -low 70 -mid 80 -high 100 -host localhost -port 4200 &
go run $CLIENTMAIN -read -id 1 -host localhost -port 4201







