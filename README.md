<h1 style="text-align: center;">Login microservice</h1>
This repo uses as an example using echo framework, 
moreover it uses as a template for REST based micorservice.


## How to compile & run on you computer
* clone git repo
* in the folder write the following `go build ./cmd/main`
    * for windows: run `main.exe`
    * for linux: run `./main`

## How to run docker file
* clone git repo
* build - `docker build -t matan:v1 .`
* run   - `docker run -p '8080:8080' --name matanC matan:v1`
* Have fun :)

## Curls for testing
* **(expected 200)** `curl -k -X GET http://localhost:8080/health`
* **(expected 400)** `curl -X POST http://localhost:8080/auth/login -d '{"mail": "test", "password": "test"}' -H 'content-type: application/json'`
* **(expected 200)** `curl -X POST http://localhost:8080/auth/login -d '{"mail": "test3", "password": "test"}' -H 'content-type: application/json'`
* **(expected 200)** `curl -X POST http://localhost:8080/auth/register -d '{"mail": "test", "password": "test", "firstName": "test", "lastName": "test"}' -H 'content-type: application/json'`

