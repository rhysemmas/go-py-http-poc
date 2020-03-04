# go-py-http-poc

A small repo containing two apps (client/server) written in Golang and Python.

Mainly being used to test the behaviour of Go's net/http Post method.

## app.py

A Python Flask app which serves a simple 'hello' response (binds to `localhost:5000`).

Usage:  
```
python3 -m flask run --host=0.0.0.0
```

## app.go

A Go app which sends requests to `localhost:5000` and reads the response.

Usage:
```
go run app.go
```
