# genius-api
Golang bindings for Genius API

[![GoDoc](https://godoc.org/github.com/gabyshev/genius-api/genius?status.svg)](https://godoc.org/github.com/gabyshev/genius-api/genius)


## Usage

```go
import (
	"fmt"
	"github.com/gabyshev/genius-api/genius"
)

func main() {
	accessToken := "token"
	client := genius.NewClient(nil, accessToken)

	response, err := client.GetArtistHTML(16775)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Response.Artist)
}

```
