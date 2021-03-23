# HRPWinG is a HTTP Reverse Proxy Written in Go

HRPWinG is a very simple and lightweight HTTP reverse proxy which provides the
following features:

* HTTP request and response forwarding.
* HTTP authentication.
* Logging for monitoring purposes.

This project is currently a **Work in Progress**.

## Building

```console
go build
```

## Running

```console
./HRPWinG 
Listening on port :3000 
```

Then access http://localhost:3000 from a web browser and the golang.org page
should open.

## Testing

```console
cd tests && go test
```

