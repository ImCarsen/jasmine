<h1 align="center">Jasmine 🫖</h1>

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
            Routes: map[string]http.HandlerFunc{
                "POST /test/{id}": test,
            },
            ProtectedRoutes: map[string]http.HandlerFunc{
                "POST /test/{id}": test,
            },
            AuthFunc: jasmine.DefaultAuthFunc,
        },
    } 

   go s.Start()

   <-sig
   s.Stop()
}

func test(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(fmt.Sprintf("ID: %v", id)))
}
```

## Contributing

I welcome any contributions, as I don't do any perfomance testing, and the practices used for Jasmine are unlikely to be the best (it's just how I felt like writing it in the moment). So, if you feel like improving this project at all, feel free to open an issue, or PR.
