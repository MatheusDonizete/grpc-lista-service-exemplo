# Introdução

Repositório criado como exemplo de implementação de uma aplicação gRPC, tendo um servidor desenvolvido em Golang e um client feito em Node.js com geração automática de código para consumo das funções do servidor.

## Como rodar

Antes de mais nada é necessário gerar os arquivos com o código necessário para o servidor poder funcionar corretamente executando os seguintes comandos:

```shell
$ protoc --go_out=./lista-service-server --go_opt=paths=source_relative
      --go-grpc_out=./lista-service-server --go-grpc_opt=paths=source_relative
      protos/**/*.proto

```

Caso ainda não tenha feito, instalar as dependências do compilador:

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

### Servidor

Apenas acesse o diretório `./lista-service-server` e execute:

```shell
$ go run main.go
```

### Cliente

Apenas acesse o diretório `./lista-service-client` e execute os seguintes comandos:

```shell
$ npm install
$ npm start
```