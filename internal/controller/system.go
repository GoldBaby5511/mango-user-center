package controller

import (
	"mango-user-center/internal/middleware"
)

type system struct{}

func init() {
	var r = Engine.Group("system")
	r.Use(middleware.MustIntranet())

}
