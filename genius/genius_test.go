package genius_test

import (
	"fmt"
	"github.com/gabyshev/genius-api/genius"
)

func ExampleNewClient() {
	accessToken := "token"
	client := genius.NewClient(nil, accessToken)

	user, err := client.GetAccount()
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Email)
}
