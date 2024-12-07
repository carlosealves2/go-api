package goapi

type MiddlewareFunc func(handlerFunc HandlerFunc) HandlerFunc
