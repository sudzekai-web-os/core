package core

import (
	"net/http"
)

type IHandlersRegistry interface {
	AddHandler(pattern string, hnd HandlerFunc) IHandlersRegistry

	AddProtectedHandler(pattern string, hnd HandlerFunc, roles []string) IHandlersRegistry

	AddNoFilterHandler(pattern string, hnd http.HandlerFunc) IHandlersRegistry

	AddProtectedNoFilterHandler(pattern string, hnd http.HandlerFunc, roles []string) IHandlersRegistry

	AddMiddleware(pos int, middleware Middleware) IHandlersRegistry

	SetJwtMiddleware(jwt JwtMiddleware) IHandlersRegistry

	SetResultFilter(filter ResultFilter) IHandlersRegistry

	GetRoutes() []string
	ClearRoutes() IHandlersRegistry

	GetHandler() http.Handler
}
