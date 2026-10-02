package stratigies

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-todo/internal/ports"
)

// --------------------------------CONSTANTS---------------------------------- //

// --------------------------------STRUCT---------------------------------- //

type AccessTokenStrategy struct {
	Ctx        *fiber.Ctx
	JwtAdapter ports.IJWTAdapterPort
}

// --------------------------------PUBLIC_METHODS---------------------------------- //

func (a *AccessTokenStrategy) Authenticate(strategy ports.IJwtExtractorStrategyPort, payload *ports.AccessTokenPayload) error {
	token, err := a.ExtractToken(strategy)
	if err != nil {
		return err
	}

	if err := a.validate(token, payload); err != nil {
		return err
	}

	return nil
}

func (a *AccessTokenStrategy) ExtractToken(strategy ports.IJwtExtractorStrategyPort) (string, error) {
	token, err := strategy.Extract()
	if err != nil {
		return "", err
	}
	return token, nil
}

// --------------------------------PRIVATE_METHODS---------------------------------- //

func (a *AccessTokenStrategy) validate(token string, payload *ports.AccessTokenPayload) error {
	if err := a.JwtAdapter.Validate(token, payload); err != nil {
		return err
	}
	return nil
}
