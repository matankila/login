<p align="center">
<img src="https://imgur.com/0ViiRT1.png" width="300" />
</p>
<h1 align="center">Login microservice</h1>
This repo uses as an example using echo framework, 
moreover it uses as a template for REST based micorservice.

## :trophy: prerequisites
* go 1.12 + .

## :zap: How to compile & run locally
* clone git repo.
* open cmd in the folder.
* write the following `go build ./cmd/main`
    * for windows: run the `main.exe` through terminal / ui.
    * for linux: run `./main` through terminal.

## :zap: How to build & run dockerfile
* clone git repo.
* open cmd in the folder.  
* build - `docker build -t matan:v1 .`
* run   - `docker run -p '8080:8080' --name matanC matan:v1`
* Have fun :)

## :triangular_ruler: Curls for testing
* **(expected 200)** `curl -k -X GET http://localhost:8080/health`
* **(expected 400)** `curl -X POST http://localhost:8080/auth/login -d '{"mail": "test", "password": "test"}' -H 'content-type: application/json'`
* **(expected 200)** `curl -X POST http://localhost:8080/auth/login -d '{"mail": "test3", "password": "test"}' -H 'content-type: application/json'`
* **(expected 200)** `curl -X POST http://localhost:8080/auth/register -d '{"mail": "test", "password": "test", "firstName": "test", "lastName": "test"}' -H 'content-type: application/json'`