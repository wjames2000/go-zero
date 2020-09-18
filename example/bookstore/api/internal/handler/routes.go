// DO NOT EDIT, generated by goctl
package handler

import (
	"bookstore/api/internal/svc"
	"net/http"

	"github.com/wjames2000/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/add",
			Handler: addHandler(serverCtx),
		},
		{
			Method:  http.MethodGet,
			Path:    "/check",
			Handler: checkHandler(serverCtx),
		},
	})
}
