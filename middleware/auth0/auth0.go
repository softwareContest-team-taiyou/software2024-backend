package auth0

import (
	"errors"

	"github.com/golang-jwt/jwt" // Ensure this is the library used throughout
)

// JWKSからJWTで使われているキーをPEM形式で返す
func getPemCert(jwks *JWKS, token *jwt.Token) (string, error) {
    cert := ""

    for _, k := range jwks.Keys {
        if token.Header["kid"] == k.Kid {
            cert = "-----BEGIN CERTIFICATE-----\n" + k.X5c[0] + "\n-----END CERTIFICATE-----"
            break
        }
    }

    if cert == "" {
        return "", errors.New("unable to find appropriate key")
    }

    return cert, nil
}
