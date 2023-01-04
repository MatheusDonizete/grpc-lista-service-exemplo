package api

import (
	"grpc-lista-de-compra/controllers"
)
func StartGrpcServers() {
	var listaServer controllers.ListaServer
	listaServer.Server()
}
	