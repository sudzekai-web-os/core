package core

import "net/http"

type JwtMiddleware func(http.HandlerFunc, []string) http.HandlerFunc
