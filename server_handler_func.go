package core

import "net/http"

type HandlerFunc func(*http.Request) HandlerResult
