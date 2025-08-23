package service

import (
	"ai-kcal-agent/pkg/appContext"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func IssueJWT(login string) (string, error) {
	ctx := appContext.Get()
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   login,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(10 * time.Minute)),
		NotBefore: jwt.NewNumericDate(now.Add(-1 * time.Minute)),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(ctx.JwtSecret))
}

// Парсер с ограничением алгоритма и небольшим леевеем по часам
var parser = jwt.NewParser(
	jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	jwt.WithLeeway(30*time.Second),
)

func ParseJWT(tokenStr string) (string, error) {
	ctx := appContext.Get()
	if tokenStr == "" {
		return "", errors.New("token is empty")
	}

	var claims jwt.RegisteredClaims
	tok, err := parser.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(ctx.JwtSecret), nil
	})
	if err != nil || !tok.Valid {
		return "", err
	}
	if claims.Subject == "" {
		return "", errors.New("subject is empty")
	}
	return claims.Subject, nil
}
