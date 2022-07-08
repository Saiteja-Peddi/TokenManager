package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"sync"

	pb "example.com/AOS_PRJ2/tokenmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v3"
)

type Domain struct {
	Low  uint64
	Mid  uint64
	High uint64
}

type State struct {
	Final uint64
}

type Token struct {
	Id     string
	mut    sync.Mutex
	Name   string
	domain *Domain
	state  *State
	Ts     uint64
}

type YmlData struct {
	Token   string
	Writer  string
	Readers []string
}

var Tokens = make(map[string]*Token)

type tokenService struct {
	pb.UnimplementedTokenManagerServer
}

// Hash concatentates a message and a nonce and generates a hash value.
func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}

func minOf(val1 uint64, val2 uint64) uint64 {
	if val1 < val2 {
		return val1
	} else {
		return val2
	}
}

func minHash(name string, min uint64, max uint64) uint64 {
	var minVal uint64 = Hash(name, min)
	var minX uint64 = min

	for i := min + 1; i < max; i++ {
		var tempVal = Hash(name, i)
		minVal = minOf(minVal, tempVal)
		if minVal == tempVal {
			minX = i
		}
	}
	return minX
}

func printAllTokes() {
	var i = 1
	fmt.Println("\n----------------------------------------------------------------")
	for key, val := range Tokens {
		fmt.Print("\nToken ", i, ":- Id: ", key)
		if val != nil && &val.Name != nil && val.Name != "" {
			fmt.Print(" Name: ", val.Name)
		}
		if val != nil && val.state != nil && &val.state.Final != nil {
			fmt.Print(" Final: ", val.state.Final)
		}
		i += 1
	}
	fmt.Println("\n----------------------------------------------------------------")
}

func printToken(id string) {
	log.Print("\n")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Current Token:-")
	//fmt.Println("Name: ", Tokens[id].Name, " Partial Value: ", Tokens[id].state.Partial, " Final Value: ", Tokens[id].state.Final)
	fmt.Printf("Id: %s", id)
	if Tokens[id] != nil && &Tokens[id].Name != nil && Tokens[id].Name != "" {
		fmt.Printf(" Name: %s", Tokens[id].Name)
	}
	if Tokens[id] != nil && Tokens[id].state != nil && &Tokens[id].state.Final != nil {
		fmt.Printf(" Final: %d", Tokens[id].state.Final)
	}
	fmt.Println("\n----------------------------------------------------------------")

}

func (t *tokenService) Create(_ context.Context, req *pb.NormalRequest) (*pb.ServerResponse, error) {

	var createToken = &Token{
		Id: req.GetId(),
	}
	Tokens[req.GetId()] = createToken
	printToken(req.GetId())
	printAllTokes()
	return &pb.ServerResponse{
		Message: 1,
		Id:      req.GetId(),
	}, nil
}

func (t *tokenService) Read(_ context.Context, req *pb.NormalRequest) (*pb.ServerResponse, error) {
	fmt.Println("Read Rpc invoked at port 4200")
	printToken(req.GetId())
	return &pb.ServerResponse{
		Message: Tokens[req.GetId()].state.Final,
		Id:      req.GetId(),
	}, nil
}

func writeIntoToken(mut *sync.Mutex, req *pb.WriteRequest, wg *sync.WaitGroup) {

	mut.Lock()
	Tokens[req.GetId()] = &Token{
		Name: req.GetName(),
		domain: &Domain{
			Low:  req.GetLow(),
			Mid:  req.GetMid(),
			High: req.GetHigh(),
		},
		state: &State{
			Final: minHash(req.GetName(), req.GetLow(), req.GetHigh()),
		},
	}
	Tokens[req.GetId()].Ts += 1
	mut.Unlock()

	wg.Done()

}

func (t *tokenService) Write(ctx context.Context, req *pb.WriteRequest) (*pb.ServerResponse, error) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(1)
	go writeIntoToken(&mutex, req, &wg)

	wg.Wait()
	fmt.Println("Write Rpc invoked at port 4200")
	printToken(req.GetId())
	//printAllTokes()
	duplicate(ctx, req)

	return &pb.ServerResponse{
		Message: Tokens[req.GetId()].state.Final,
		Id:      req.GetId(),
	}, nil

}

func duplicate(ctx context.Context, req *pb.WriteRequest) {
	var readerAddresses []string
	yfile, err := ioutil.ReadFile("./../replScheme.yml")
	if err != nil {
		log.Fatal(err)
	}

	ymlData := make(map[string][]YmlData)
	err2 := yaml.Unmarshal(yfile, &ymlData)
	if err2 != nil {
		log.Fatal(err2)
	}
	tokens := ymlData["tokens"]
	for i := 0; i < len(tokens); i++ {
		if req.GetId() == tokens[i].Token {
			readerAddresses = tokens[i].Readers
		}
	}
	copyReq := &pb.CopyRequest{Id: req.GetId(), Name: req.GetName(), Low: req.GetLow(), Mid: req.GetMid(),
		High: req.GetHigh(), Final: Tokens[req.GetId()].state.Final, Ts: Tokens[req.GetId()].Ts}
	for i := 0; i < len(readerAddresses); i++ {
		if !strings.Contains(readerAddresses[i], "4200") {
			conn, err := grpc.DialContext(ctx, readerAddresses[i], grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			ser := pb.NewTokenManagerClient(conn)
			res, err := ser.CopyToken(ctx, copyReq)
			if err != nil {
				log.Fatalf("Unable to write: %v", err)
			}
			fmt.Println("Replicated at address: ", readerAddresses[i], " with response ", res.Message)
		}
	}
}

func (t *tokenService) Drop(_ context.Context, req *pb.NormalRequest) (*pb.ServerResponse, error) {

	delete(Tokens, req.GetId())
	log.Print("\n")
	printAllTokes()
	return &pb.ServerResponse{
		Message: 1,
		Id:      req.GetId(),
	}, nil
}

func (t *tokenService) CopyToken(ctx context.Context, req *pb.CopyRequest) (*pb.ServerResponse, error) {

	Tokens[req.GetId()] = &Token{
		Name: req.GetName(),
		domain: &Domain{
			Low:  req.GetLow(),
			Mid:  req.GetMid(),
			High: req.GetHigh(),
		},
		state: &State{
			Final: req.Final,
		},
		Ts: req.Ts,
	}
	return &pb.ServerResponse{
		Message: 1,
		Id:      req.GetId(),
	}, nil
}

func createAllTokens() {
	yfile, err := ioutil.ReadFile("./../replScheme.yml")
	if err != nil {
		log.Fatal(err)
	}

	ymlData := make(map[string][]YmlData)
	err2 := yaml.Unmarshal(yfile, &ymlData)
	if err2 != nil {
		log.Fatal(err2)
	}
	tokens := ymlData["tokens"]
	for i := 0; i < len(tokens); i++ {
		var createToken = &Token{
			Id: tokens[i].Token,
			Ts: 0,
		}
		Tokens[tokens[i].Token] = createToken
	}
}

func main() {

	var port = flag.Int("port", 4200, "Defines from which port to listen")
	flag.Parse()
	lis, err3 := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err3 != nil {
		log.Fatalf("failed to listen: %v", err3)
	}

	rpcServer := grpc.NewServer()
	tokenServer := &tokenService{}

	pb.RegisterTokenManagerServer(rpcServer, tokenServer)
	log.Printf("Server is listening at %v", lis.Addr())

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	createAllTokens()

}
