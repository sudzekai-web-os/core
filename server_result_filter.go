package core

import "net/http"

type ResultFilter func(http.HandlerFunc, HandlerResult)
