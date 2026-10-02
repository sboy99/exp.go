package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-todo/infra/exception"
	"github.com/sboy99/exp.go/go-todo/internal/ports"
	"github.com/sboy99/exp.go/go-todo/stratigies"
)

// ------------------------------CONSTANTS---------------------------------- //

var (
	AccessTokenStrategy     = &stratigies.AccessTokenStrategy{}
	BearerExtractorStrategy = &stratigies.BearerExtractorStrategy{}
)

// ------------------------------MIDDLEWARES---------------------------------- //

func (m *Middleware) UseAuthMiddleware(strategy ports.IAuthStrategyPort) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		AccessTokenStrategy = &stratigies.AccessTokenStrategy{
			JwtAdapter: m.JwtAdapter,
			Ctx:        c,
		}

		BearerExtractorStrategy = &stratigies.BearerExtractorStrategy{
			Ctx: c,
		}

		payload := new(ports.AccessTokenPayload)
		if err := strategy.Authenticate(BearerExtractorStrategy, payload); err != nil {
			exception.HandleHttpException(c,
				exception.UnauthrorizedException(err),
			)
		} else {
			c.Locals("user", payload)
			c.Next()
		}

		return nil
	}
}
