package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "example.com/AOS_PRJ2/tokenmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type YmlData struct {
	Token   string
	Writer  string
	Readers []string
}

var (
	create = flag.Bool("create", false, "Calls Create method if true")
	write  = flag.Bool("write", false, "Calls Write method if true")
	read   = flag.Bool("read", false, "Calls Read method if true")
	drop   = flag.Bool("drop", false, "Calls Drop method if true")
	id     = flag.String("id", "", "Defines what id to use")
	name   = flag.String("name", "", "Defines what name to use")
	host   = flag.String("host", "localhost", "Defines hosting domain")
	port   = flag.String("port", "4200", "Defines port")
	low    = flag.Uint64("low", 0, "Defines low")
	mid    = flag.Uint64("mid", 10, "Defines mid")
	high   = flag.Uint64("high", 20, "Defines high")
)

func main() {
	var address string

	flag.Parse()

	address = *host + ":" + *port
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTokenManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if *read {

		res, err := c.Read(ctx, &pb.NormalRequest{Id: *id})
		if err != nil {
			log.Fatalf("Unable to read token val: %v", err)
		}
		//log.Println("Read Rpc Invoked on Port", *port)
		log.Println("Read Rpc: Final Value: ", res.GetMessage())

	} else if *write {

		res, err := c.Write(ctx, &pb.WriteRequest{Id: *id, Name: *name, Low: *low, Mid: *mid, High: *high})
		if err != nil {
			log.Fatalf("Unable to write: %v", err)
		}
		//log.Println("Write Rpc Invoked on Port", *port)
		log.Println("Write Rpc: Final Value: ", res.GetMessage())

	} else if *drop {
		res, err := c.Drop(ctx, &pb.NormalRequest{Id: *id})
		if err != nil {
			log.Fatalf("Unable to drop: %v", err)
		}
		if res.GetMessage() == 1 {
			log.Printf("Successfully droped specified token")
		}
	}

}
