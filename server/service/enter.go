package service

import (
	"github.com/liujianjiang/gin-vue-admin/server/service/autocode"
	"github.com/liujianjiang/gin-vue-admin/server/service/example"
	"github.com/liujianjiang/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
