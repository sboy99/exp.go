package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers"
)

type UserRouter struct {
	Router     gin.IRouter
	Controller controllers.IUserController
}

func (r *UserRouter) Register() {
	// create user
	r.Router.POST("/users", r.Controller.Create)
	// list users
	r.Router.GET("/users", r.Controller.List)
	// get one user
	r.Router.GET("/users/:id", r.Controller.GetById)
	// update one user
	r.Router.PATCH("/users/:id", r.Controller.UpdateById)
	// delete one user
	r.Router.DELETE("/users/:id", r.Controller.DeleteById)
}
