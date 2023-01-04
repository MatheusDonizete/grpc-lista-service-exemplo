const path = require('path');
const fastify = require('fastify')({ logger: true });
const PROTO_PATH = path.join( __dirname, '../protos/lista/lista.proto');

const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  });

const target = 'localhost:50001';
const listasProto = grpc.loadPackageDefinition(packageDefinition).lista;

fastify.get('/sync', async () => {
  const client = new listasProto.ListaService(target, grpc.credentials.createInsecure());
  const data = await new Promise(
    resolve => client.GetAllListasSync({}, (err, res) => resolve(res?.listas))
  )
  return data;
})

fastify.post('/', async () => {
  const client = new listasProto.ListaService(target, grpc.credentials.createInsecure());
  const stream = client.RecordLista();
  stream.send(JSON.stringify(req.body));
  stream.end();
  return {};
})

fastify.get('/', async () => {
  const client = new listasProto.ListaService(target, grpc.credentials.createInsecure());
  const data = await new Promise(
    resolve => {
      const listas = [];
      const call = client.GetAllListas();
      call.on('data', (lista) => listas.push(lista));
      call.on('end', () => resolve(listas));
    }
  )
  return data;
})

const start = async () => {
  try {
    await fastify.listen({ port: 3000 })
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }
}
start()