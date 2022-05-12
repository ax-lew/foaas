# FOAAS


### Requirements
* Go >= 1.15
* Make >= 3
* GOPATH environment variable set up and included in your $PATH


### Tests
To run the tests you can execute:

```shell script
$ make test
```

### Run locally
You will need to download the dependencies:
```shell script
$ make deps
```
Then run:
```shell script
$ go run ./main.go serve --max-requests 5 --interval-ms 10000 --foaas-timeout-ms 1000
```
The server will run on port 8080.

To get a message:
```shell script
$ curl -H 'user-id: 1234' 'http://localhost:8080/message'
```

### Run with Docker
The server can be run using Docker. Use:
```shell script
$ make docker-build
$ make docker-run
``` 