package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	pb "devfest/catfacts/catfacts"
)

const (
	port = ":50051"
)

type catServer struct {
	savedFacts []*pb.CatFactResponse
}

func (s *catServer) GetCatFact(ctx context.Context, in *pb.CatFactRequest) (*pb.CatFactResponse, error) {
	fact := pb.AllTheFacts[in.Id]
	resp := &pb.CatFactResponse{Fact: fact}
	return resp, nil
}

func (s *catServer) ListCatFacts(in *pb.CatFactRequest, stream pb.CatFacts_ListCatFactsServer) error {
	for _, fact := range s.savedFacts {
		if err := stream.Send(fact); err != nil {
			return err
		}
	}
	return nil
}

func (s *catServer) StreamCatFacts(stream pb.CatFacts_StreamCatFactsServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		resp := &pb.CatFactResponse{Fact: pb.AllTheFacts[in.Id-1]}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func newServer() *catServer {
	s := &catServer{savedFacts: make([]*pb.CatFactResponse, len(pb.AllTheFacts))}
	s.fillFacts()
	return s
}

func (s *catServer) fillFacts() {
	facts := make([]*pb.CatFactResponse, len(pb.AllTheFacts))
	for i, fact := range pb.AllTheFacts {
		facts[i] = &pb.CatFactResponse{Fact: fact}
	}
	copy(s.savedFacts, facts)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCatFactsServer(server, newServer())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
