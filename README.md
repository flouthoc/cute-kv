# :bear: cute-kv

<strong>[Assignment]</strong> Is a minimal quick <strong>in-memory / persistant</strong> kv store written in go with basic <strong>HTTP API</strong> handle.

### Server

#### Building Server
```bash
cd server
go build server.go
```

#### Starting Server as in-memory store

By default server will start on <strong>localhost:9992</strong>

```bash
./server <port>
or
./server
```

#### Server Usage

It is advisable to use <strong>[cli-client](https://github.com/flouthoc/cute-kv#client)</strong> but you can always use raw http for basic operations, please scroll below for [raw http endpoints](https://github.com/flouthoc/cute-kv#server-api-or-using-with-curl).

##### :no_good: [Scroll to extreme bottom if you want to run server as persistance store instead of default in-memory store. [Would not recommend as it involves heavy disk i/o]](https://github.com/flouthoc/cute-kv#running-server-on-persistance-mode)

---
### Client

#### Building Client
```bash
cd cli-client
go build cliclient.go
```

#### Using cli-client

### Set
##### - Sets value for corresponding key or overrides
```bash
./cliclient set host:port <key> <value>
```

<strong>Example 1</strong>
```bash
./cliclient set 127.0.0.1:9992 hello world
```

<strong>Example 2 </strong>
```bash
./cliclient set 127.0.0.1:9992 jack "jack jack samurai jack jack"
```
---

### Get 
##### - Gets value for corresponding key
```bash
./cliclient get host:port <key>
```

<strong>Example 1</strong>
```bash
./cliclient get 127.0.0.1:9992 hello
```

---
### Watch 
##### - Subscribes to a given key and watches it for any change
```bash
./cliclient watch host:port <key>
```

<strong>Example 1</strong>
```bash
./cliclient watch 127.0.0.1:9992 hello
```

---
### Server API or using with Curl

##### Set
Set could be either a http <strong>POST</strong> or http <strong> GET </strong> call to <strong> host:port/set </strong>

```bash
curl -d "key=hello&value=hello" -X POST http://localhost:9992/set
```

##### Get
Get could be  http <strong> GET </strong> call to <strong> host:port/set </strong>

```bash
curl http://localhost:9992/get?key=hello
```

---
### Running Server on persistance mode

Persistance mode will ensure that server flushes every transaction to disk.
Warning: Method involves heavy disk i/o , not recommanded

<strong>flag true ensures that server is running on persistance mode</strong>

```bash
./server <port> true
```