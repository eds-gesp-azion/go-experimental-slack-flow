# go-experimental-slack-flow

Simple HTTP API that returns a random quote from thinkers in technology, AI, software engineering and the internet on every `GET /`.

## Run

```sh
go run .
curl localhost:8080/
```

The port defaults to `8080` and can be changed with the `PORT` environment variable.

## Response

```json
{"text":"Clear is better than clever.","author":"Rob Pike"}
```
