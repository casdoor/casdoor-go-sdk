package main

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"log"
)

var CasdoorEndpoint = "http://localhost:8000"
var ClientId = "55fca1d91a9bfd9c809d"
var ClientSecret = "3c8aa90b7d2fc9d25664e777e24922771f590609"

//var CasdoorOrganization = "built-in"
var CasdoorOrganization = "admin"
var CasdoorApplication = ""

var JwtPublicKey string

func main() {
	log.Printf("dsfsdf")

	casdoorsdk.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtPublicKey, CasdoorOrganization, CasdoorApplication)

	//userUUID := "9d857b30-dbbb-43e3-9dc0-bda9a5800978"
	//user, err := casdoorsdk.GetUserByUserId(userUUID)
	//if err != nil {
	//	log.Fatalf("Err get users: %v", err)
	//}
	//log.Printf("user: %v", user)

	tokens, _, err := casdoorsdk.GetTokens(0, 10)
	if err != nil {
		log.Fatalf("err get tokens: %v", err)
	}

	log.Printf("count: %d", len(tokens))

	tokenName := "c0eaf61d-0e7d-48c8-a681-0de94d3c2320"
	affected, err := casdoorsdk.DeleteToken(tokenName)
	if err != nil {
		log.Fatalf("Err delete: %v", err)
	}
	log.Printf("affected: %v", affected)
}
