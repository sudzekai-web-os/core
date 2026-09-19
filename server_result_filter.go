package core

import "net/http"

type ResultFilter func(http.ResponseWriter, *http.Request, HandlerResult)
