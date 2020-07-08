package main
import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*https://C:\Go\src\HereMaps\OnJson.go
    bbox=41.389405513925354,2.127549994463742,41.38042236108416,2.139522979169079
    &datetime=2019-01-21T09:42:10Z*/
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":80", nil)
}