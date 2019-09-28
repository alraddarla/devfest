package main

import (
	"context"
	pb "devfest/catfacts/catfacts"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCatFactsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if len(os.Args) > 1 {
		switch strings.ToUpper(os.Args[1]) {
		case "GET":
			var factNumber int64
			factNumber, err = strconv.ParseInt(os.Args[2], 10, 64)
			if err != nil {
				log.Fatalf("%s is not an ID number", os.Args[1])
			}
			getFact(ctx, client, factNumber)
		case "LIST":
			listFacts(ctx, client, &pb.CatFactRequest{})
		case "STREAM":
			streamFacts(ctx, client)
		}
	} else {
		log.Fatalf("At least one argument is required ")
	}
}

func streamFacts(ctx context.Context, client pb.CatFactsClient) {
	stream, err := client.StreamCatFacts(ctx)
	if err != nil {
		log.Fatalf("%v.StreamCatFacts(_) = _, %v", client, err)
	}
	waitchan := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitchan)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive cat fact: %v", err)
			}
			log.Printf("Cat Fact #%d: %s",in.FactNum, in.Fact)
		}
	}()

	for _, num := range pb.RandomFifty {
		if err := stream.Send(&pb.CatFactRequest{Id: num}); err != nil {
			log.Fatalf("Failed to send id: %v ", err)
		}
	}
	stream.CloseSend()
	<-waitchan
}

func listFacts(ctx context.Context, client pb.CatFactsClient, req *pb.CatFactRequest) {
	stream, err := client.ListCatFacts(ctx, req)
	if err != nil {
		log.Fatalf("%v.ListCatFacts(_) = _, %v", client, err)
	}
	for {
		fact, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListCatFacts(_) = _, %v", client, err)
		}
		log.Printf("Fact #%d: %s\n", fact.FactNum, fact.Fact)
	}
}

func getFact(ctx context.Context, client pb.CatFactsClient, factNumber int64) {
	resp, err := client.GetCatFact(ctx, &pb.CatFactRequest{Id: factNumber - 1})
	if err != nil {
		log.Fatalf("could not get cat fact: %v ", err)
	}
	log.Printf("Fun fact #%v: %s \n", factNumber, resp.Fact)
}
