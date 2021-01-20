# Gopherish

A translator app + REST API server to translate English into Gopherish.

As per the definitive grammar guide on English to Gopherish translation:
1. If a word starts with a vowel letter, add prefix “g” to the word (ex. apple => gapple)
2. If a word starts with the consonant letters “xr”, add the prefix “ge” to the begging of the word.
Such words as “xray” actually sound in the beginning with vowel sound as you pronounce them so a true gopher would say “gexray”.
3. If a word starts with a consonant sound, move it to the end of the word and then add “ogo” suffix to the word.
Consonant sounds can be made up of multiple consonants, a.k.a. a consonant cluster (e.g. "chair" -> "airchogo”).
4. If a word starts with a consonant sound followed by "qu", move it to the end of the word, and then add "ogo" suffix to the word (e.g. "square" -> "aresquogo").
*/
### Installation

Either clone this git repo or install using `go get`:
```
$ go get -u github.com/batyanko/gopherish
```

Then change into the main project directory and build using `go build`.

Run the tool by supplying a listen port of your choice:
```
$ ./gopherish 10000
```

### Setup

`Makefile` targets are available for testing and linting the project.

The make tool will be needed for that. Also be sure to have your `go/bin` directory in your PATH.

```
$ make help
``` 

### REST Endpoints
##### Translate words via the `/word` POST endpoint:
```
$ curl -d '{"english-word":"pogo"}' -H 'Content-Type: application/json' http://localhost:10000/word
```

```json
{"gopher-word":"ogopogo"}
```

##### Translate sentences via the `/sentence` POST endpoint.
```
$ curl -d '{"english-sentence":"Who ate pogo?"}' -H 'Content-Type: application/json' http://localhost:10000/sentence
```

```json
{"gopher-sentence":"Owhogo gate ogopogo?"}
```

##### Display history of past translations via the `history` GET endoint.
```
$ curl http://localhost:10000/history
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
