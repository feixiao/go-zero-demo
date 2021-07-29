// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/feixiao/go-zero-demo/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/user",
				Handler: CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/user/:id",
				Handler: GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/v1/user/:id",
				Handler: DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/v1/user/:id",
				Handler: UpdateUserHandler(serverCtx),
			},
		},
	)
}
