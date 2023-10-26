package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"mnc/merchant-bank-api/config"
	"mnc/merchant-bank-api/model"
	"time"
)

func CreateAccessToken(customer model.Auth) (string, error) {
	confg, _ := config.NewConfig()

	now := time.Now().UTC()
	end := now.Add(confg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    confg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: customer.Username,
	}

	token := jwt.NewWithClaims(confg.JwtSigningMethod, claims)
	ss, err := token.SignedString(confg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed to create access token :%s", err.Error())
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	confg, _ := config.NewConfig()

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != confg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return confg.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid parse token :%s", err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid || claims["iss"] != confg.ApplicationName {
		return nil, fmt.Errorf("invalid token MapClaims")
	}
	return claims, nil
}
