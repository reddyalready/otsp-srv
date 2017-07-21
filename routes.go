package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ProjectList",
		"GET",
		"/projects",
		ProjectList,
	},
	Route{
		"ProjectGet",
		"GET",
		"/projects/{projectId}",
		ProjectGet,
	},
}
