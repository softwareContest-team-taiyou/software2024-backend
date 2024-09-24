package auth0

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt" // Updated import path
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)
func AuthInterceptor(domain, clientID string, jwks *JWKS) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if vals, ok := md["authorization"]; ok && len(vals) > 0 {
				tokenStr := strings.TrimPrefix(vals[0], "Bearer ")
	
				token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
				
					cert, err := getPemCert(jwks, token)
					if err != nil {
						return nil, err
					}
					return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
				})
	

				if err != nil || !token.Valid {
					fmt.Print(err)
					return nil, status.Errorf(codes.Unauthenticated, "invalid or expired token: %v", err)
				}
				  claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					return nil, status.Error(codes.Unauthenticated, "error accessing claims")
				}

				// Access custom claim, e.g., "uid"
				uid, ok := claims["sub"].(string)
				if !ok {
					return nil, status.Error(codes.Unauthenticated, "uid not found in token")
				}
		

        		// Optionally, add the UID to the context for use in subsequent handlers
       			ctx = context.WithValue(ctx, "uid", uid)
			
				return handler(ctx, req)
			}
		}
		return nil, status.Error(codes.Unauthenticated, "authorization token is required")
	}
}