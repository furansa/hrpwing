# HTTP Reverse Proxy Written in Go

HRPWinG is a very simple and lightweight HTTP reverse proxy which provides the
following features:

* HTTP request and response forwarding.
* HTTP authentication.
* Logging.

This project is currently a *experimental* **Work in Progress**.

## Building

```console
$ go build
```

## Running

```console
$ ./HRPWinG 
Listening on port :3000 
```

Then access http://localhost:3000 from a web browser and the golang.org page
should open.

## Testing

```console
$ cd tests && go test
```

