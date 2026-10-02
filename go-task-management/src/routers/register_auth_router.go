package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers"
)

type AuthRouter struct {
	Router     gin.IRouter
	Controller controllers.IAuthController
}

func (r *AuthRouter) Register() {
	// register user
	r.Router.POST("/register", r.Controller.Register)
}
