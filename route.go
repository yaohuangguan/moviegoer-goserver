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
		"indexHandler",
		"GET",
		"/",
		indexHandler,
	},
	Route{
		"获取用户登录态",
		"GET",
		"/api/v1/user/login",
		loginGetHandler,
	},
	Route{
		"用户登录",
		"POST",
		"/api/v1/user/login",
		loginPostHandler,
	},
}
