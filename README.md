# Tx Parser

With tx-parser you can track ethereum transaction for particular addresses.
Just run tx-parser add addresses to observe and it will collect all transactions coming from an to these addresses.

## First launch
You can build tx-parser from source using `make`
```bash
make build
```
And after that run it with configuration
```bash
./dist/tx-parser -c config.json
```

## Configuring

You can configure tx-parser with json configuration file.
An example config file looks like this

```json
{
  "http": {
    "port": "313"
  },
  "ethereum": {
    "rpc_url": "https://cloudflare-eth.com"
  },
  "parser": {
    "lookup_interval": 3000,
    "starting_block": 18258488
  }
}
```
`http.port` specifies port for http server with 
`ethereum.rpc_url` specifies rpc provider
`parser.lookup_interval` specifies in milliseconds how often parser will query blockchain
`parser.starting_block` specifies a block number to look for transactions after that block 

## Control tx-parser via http endpoints

### Subscribe
Add ethereum address to observe its transactions
```bash
curl -X POST 'http://localhost:313/subscribe' --data '{"address": "0xF9EB9dC26C7a6f45A43b21CCB4bFd9f6BAbfFD3E"}'
```

### Transactions
Request transaction for an address
```bash
curl -X GET 'http://localhost:313/transactions/0x69950B080A1a11B3514b45478230c6ccaBA43d9e'
```

### Latest Parsed Block
You can check last parsed block 
```bash
curl -X GET 'http://localhost:313/currentBlock'
```

## Improvement suggestions
### Storage
Currently parser stores transactions in the application memory and it loses all the data after application exit.
In future it could use some persistent storage like PostgreSQL. It could be simply achived by writing adapter for PostgreSQL which implements parser.Storage interface.

### Parsing
Due to simplicity currently tx-parser queries and stores ethereum blocks synchronously one by one. This process could be done in parallel