package models

import "github.com/kataras/iris"

type Route struct {
	Apply func(application *iris.Application)
}
