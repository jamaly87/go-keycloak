package keycloak

import (
	"github.com/Nerzal/gocloak/v13"
)

type Keycloak struct {
	Gocloak      *gocloak.GoCloak
	ClientId     string
	ClientSecret string
	Realm        string
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient("http://localhost:8086/auth"),
		ClientId:     "your-client-id-goes-here",
		ClientSecret: "you-client-secret-goes-here",
		Realm:        "your-realm-goes-here",
	}
}
