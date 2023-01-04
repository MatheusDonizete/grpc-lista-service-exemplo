package controllers

import (
	"context"
	"flag"
	"fmt"
	"net"
	"log"
	"encoding/json"
	"google.golang.org/grpc"
	pb "grpc-lista-de-compra/protos/lista"
)

type ListaServer struct {
	pb.UnimplementedListaServiceServer
	savedListas []*pb.Lista
}

var (
	port = flag.Int("port", 50001, "The server port")
)

func (s *ListaServer) GetAllListas(req *pb.ListaRequest, stream pb.ListaService_GetAllListasServer) error{
	for _, lista := range s.savedListas {
		if err := stream.Send(lista); err != nil {
			return err
		}
	}
	return nil
}

func (s *ListaServer) GetAllListasSync(ctx context.Context, req *pb.ListaRequest) (*pb.Listas, error){
	return &pb.Listas{Listas: s.savedListas}, nil
}

func (s *ListaServer) RecordLista(stream pb.ListaService_RecordListaServer) error {
	lista := new(pb.Lista)
	if err := stream.RecvMsg(lista); err != nil {
		items := []*pb.Item{}
		return stream.SendAndClose(&pb.Lista{
			Name: "",
			Version: 0,
			Items: items,
		})
	}
	s.savedListas = append(s.savedListas, lista)
	return nil
}

var exampleData = []byte(`[{
	"name": "lista",
	"description": "descrição teste",
	"version": 1,
	"items": [
		{
			"id": 1,
			"name": "arroz",
			"value": 12.0
		},
		{
			"id": 2,
			"name": "feijão",
			"value": 3.80
		},
		{
			"id": 3,
			"name": "batata",
			"value": 2.50
		}
	]
},
{
	"name": "lista_teste",
	"display_name": "Lista de exemplo",
	"version": 1,
	"items": []
}]`)

func newServer() *ListaServer {
	s := &ListaServer{}
	if err := json.Unmarshal(exampleData, &s.savedListas); err != nil {
		log.Fatalf("Failed to load default lista: %v", err)
	}
	return s
}

func (s *ListaServer) Server() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running ListaServer on port: ", *port)

	grpcServer := grpc.NewServer()
	pb.RegisterListaServiceServer(grpcServer, newServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
