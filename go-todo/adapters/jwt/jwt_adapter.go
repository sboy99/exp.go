package jwt

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sboy99/exp.go/go-todo/infra/config"
)

// ------------------------------STRUCTS---------------------------------- //

type JwtAdapter struct{}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (a *JwtAdapter) Sign(payload interface{}) (string, error) {
	claims := a.getClaims(a.structToMap(payload))
	token, err := a.signClaims(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *JwtAdapter) Validate(tokenString string, payload interface{}) error {
	token, err := a.parseToken(tokenString)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("Invalid token")
	}

	claims, err := a.extractClaims(token)
	if err != nil {
		return err
	}

	if err := a.mapToStruct(*claims, payload); err != nil {
		return err
	}

	return nil
}

// ------------------------------PRIVATE_METHODS---------------------------------- //

func (a *JwtAdapter) getClaims(payload map[string]interface{}) jwt.MapClaims {
	claims := make(jwt.MapClaims)

	// map payload to claims
	for k, v := range payload {
		claims[k] = v
	}

	// add timestamp to claims
	claims["iat"] = time.Now().Unix()

	// done
	return claims
}

func (a *JwtAdapter) signClaims(claims jwt.MapClaims) (string, error) {
	secret := a.getSecret("access_token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (a *JwtAdapter) getSecret(secretFor string) []byte {
	jwtSecret := config.JWT_SECRET
	switch secretFor {
	case "access_token":
		return []byte(jwtSecret)
	default:
		return []byte(jwtSecret)
	}
}

func (a *JwtAdapter) parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return a.getSecret("access_token"), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *JwtAdapter) extractClaims(token *jwt.Token) (*jwt.MapClaims, error) {
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, errors.New(`Unable to extract claims`)
	}
	return claims, nil
}

// todo: move to helper
func (a *JwtAdapter) structToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(data)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
		result[field.Name] = fieldValue
	}

	return result
}

func (a *JwtAdapter) mapToStruct(m map[string]interface{}, s interface{}) error {
	sType := reflect.TypeOf(s).Elem()
	sValue := reflect.ValueOf(s).Elem()

	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		if value, ok := m[field.Name]; ok {
			sValue.Field(i).Set(reflect.ValueOf(value))
		}
	}

	return nil
}
