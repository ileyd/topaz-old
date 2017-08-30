package auth

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

var AdminGroup string = "topaz_administrators"
var validator *auth0.JWTValidator

func init() {
	//Creates a configuration with the Auth0 information
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: "https://topaz.au.auth0.com/.well-known/jwks.json"})
	audience := "https://topaz.lavender.pink/api"
	configuration := auth0.NewConfiguration(client, []string{audience}, "https://topaz.au.auth0.com/", jose.RS256)
	validator = auth0.NewValidator(configuration)
}

// LoadPublicKey loads a public key from PEM/DER-encoded data.
func LoadPublicKey(data []byte) (interface{}, error) {
	input := data

	block, _ := pem.Decode(data)
	if block != nil {
		input = block.Bytes
	}

	// Try to load SubjectPublicKeyInfo
	pub, err0 := x509.ParsePKIXPublicKey(input)
	if err0 == nil {
		return pub, nil
	}

	cert, err1 := x509.ParseCertificate(input)
	if err1 == nil {
		return cert.PublicKey, nil
	}

	return nil, fmt.Errorf("square/go-jose: parse error, got '%s' and '%s'", err0, err1)
}

func Auth0Groups(wantedGroups ...string) gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {

		tok, err := validator.ValidateRequest(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			log.Println("Invalid token:", err)
			return
		}

		claims := map[string]interface{}{}
		err = validator.Claims(c.Request, tok, &claims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			log.Println("Invalid claims:", err)
			return
		}

		metadata, okMetadata := claims["app_metadata"].(map[string]interface{})
		authorization, okAuthorization := metadata["authorization"].(map[string]interface{})
		groups, hasGroups := authorization["groups"].([]interface{})
		if !okMetadata || !okAuthorization || !hasGroups || !shouldAccess(wantedGroups, groups) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "need more privileges"})
			c.Abort()
			log.Println("Need more provileges")
			return
		}
		c.Next()
	})
}

func shouldAccess(wantedGroups []string, groups []interface{}) bool {

	if len(groups) < 1 {
		return true
	}

	for _, wantedScope := range wantedGroups {

		scopeFound := false

		for _, iScope := range groups {
			scope, ok := iScope.(string)

			if !ok {
				continue
			}
			if scope == wantedScope {
				scopeFound = true
				break
			}
		}
		if !scopeFound {
			return false
		}
	}
	return true
}
