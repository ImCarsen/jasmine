<h1 align="center">Jasmine 🫖</h1>


[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ImCarsen/jasmine?logo=go)](https://golang.org/doc/devel/release.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/ImCarsen/jasmine)](https://goreportcard.com/report/github.com/ImCarsen/jasmine)

> A simple framework to make building HTTP servers easier

## What is Jasmine?

Jasmine is just a wrapper on Go's net/http library I made to stop having to rewrite the same code in every project.

Why is the project called Jasmine?
I was thinking of a name, and saw my jasmine tea beside me. Yeah, no deeper thought.

## Usage

```sh
go get github.com/ImCarsen/jasmine
```

```Go
package main

import (
    "github.com/ImCarsen/jasmine"
)

func main() {
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM,
    os.Interrupt)

    s := jasmine.Server{
        Address: ":8080",
        Routes: jasmine.Routes{
            Routes: map[string]jasmine.RouteHandler{
                "POST /test/{id}": jasmine.NotImplemented,
            },
            ProtectedRoutes: map[string]jasmine.RouteHandler{
                "POST /test/{id}": jasmine.NotImplemented,
            },
            AuthFunc: jasmine.DefaultAuthFunc,
        },
    } 

   go s.Start()

   <-sig
   s.Stop()
}
```

## Contributing

I welcome any contributions, as I don't do any perfomance testing, and the practices used for Jasmine are unlikely to be the best (it's just how I felt like writing it in the moment). So, if you feel like improving this project at all, feel free to open an issue, or PR.
