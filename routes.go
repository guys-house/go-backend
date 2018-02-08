package main

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"CreateRecordIndex",
		"POST",
		"/record",
		CreateRecordIndex,
	},
	Route{
		"RemoveRecordIndex",
		"DELETE",
		"/record/{id}",
		RemoveRecordIndex,
	},
}