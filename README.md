cache
-----

A simple `net.http.Handler` that adds cache and expires  headers to your `http.Handler`s


#usage
```
package main

import (
    "net/http"
    "log"
    "github.com/julien/cache"
)

func handler(func w.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello there"))

}

func main() {
    http.HandleFunc("/", cache.Cache(30, handler))
    log.Fatalf(http.ListenAndServe(":8000", nil))
}
```

