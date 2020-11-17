#Gopherish

A translator app + REST API server to translate English into Gopherish.

### Installation

Either clone this git repo or install using `go get`:
```
go get -u github.com/batyanko/gopherish
```

Then change into the main project directory and build using `go build`.

Run the tool by supplying a listen port of your choice:
```
./gopherish 10000
```

### REST Endpoints
##### Translate words via the `/word` POST endpoint:
```
curl -d '{"english-word":"pogo"}' -H 'Content-Type: application/json' http://localhost:10000/word
```

```json
{"gopher-word":"ogopogo"}
```

##### Translate sentences via the `/sentence` POST endpoint.
```
curl -d '{"english-sentence":"Who ate pogo?"}' -H 'Content-Type: application/json' http://localhost:10000/sentence
```

```json
{"gopher-sentence":"Owhogo gate ogopogo?"}
```

##### Display history of past translations via the `history` GET endoint.
```
curl http://localhost:10000/history
```

```json
{
   "history":[
      {
         "pogo":"ogopogo"
      },
      {
         "Who ate pogo?":"Owhogo gate ogopogo?"
      }
   ]
}
```

### License

This project is licensed under the Apache Version 2.0 license.  
http://www.apache.org/licenses/LICENSE-2.0