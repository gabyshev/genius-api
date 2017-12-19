# genius-api

[![GoDoc](https://godoc.org/github.com/gabyshev/genius-api/genius?status.svg)](https://godoc.org/github.com/gabyshev/genius-api/genius)

Golang bindings for Genius API.
Package is not ready to use yet and may change drastically in the future.

To get token visit https://genius.com/developers

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
