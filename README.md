<p align="center">
<img src="https://imgur.com/0ViiRT1.png" width="300" />
</p>
<h1 align="center">Login microservice</h1>
This repo uses as an example using echo framework, 
moreover it uses as a template for REST based micorservice.

## :trophy: prerequisites
* go 1.12 + .

## :zap: How to compile & run locally
* :fork_and_knife: clone git repo `git clone https://github.com/matankila/login.git`.
* :microscope: open cmd in the folder.
* :hammer: build - the following `go build ./cmd/main`
    * :runner: for windows: run the `main.exe` through terminal / ui.
    * :runner: for linux: run `./main` through terminal.

## :zap: How to build & run dockerfile
* :fork_and_knife: clone git repo `git clone https://github.com/matankila/login.git`.
* :microscope: open cmd in the folder.  
* :hammer: build - `docker build -t matan:v1 .`
* :runner: run   - `docker run -p '8080:8080' --name matanC matan:v1`
* Have fun :)

## :triangular_ruler: Curls for testing
* :thumbsup: `curl -k -X GET http://localhost:8080/health`
* :thumbsdown: `curl -X POST http://localhost:8080/auth/login -d '{"mail": "test", "password": "test"}' -H 'content-type: application/json'`
* :thumbsup: `curl -X POST http://localhost:8080/auth/login -d '{"mail": "test3", "password": "test"}' -H 'content-type: application/json'`
* :thumbsup: `curl -X POST http://localhost:8080/auth/register -d '{"mail": "test", "password": "test", "firstName": "test", "lastName": "test"}' -H 'content-type: application/json'`