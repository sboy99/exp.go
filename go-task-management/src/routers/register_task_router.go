package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers"
)

type TaskRouter struct {
	Router     gin.IRouter
	Controller controllers.ITaskController
}

func (r *TaskRouter) Register() {
	// create task
	r.Router.POST("/tasks", r.Controller.Create)
	// list tasks
	// r.Router.GET("/tasks", r.Controller.List)
	// get one task
	// r.Router.GET("/tasks/:id", r.Controller.GetById)
	// update one task
	// r.Router.PATCH("/tasks/:id", r.Controller.UpdateById)
	// delete one task
	// r.Router.DELETE("/tasks/:id", r.Controller.DeleteById)
}
