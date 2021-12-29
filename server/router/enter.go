package router

import (
	"github.com/liujianjiang/gin-vue-admin/server/router/autocode"
	"github.com/liujianjiang/gin-vue-admin/server/router/example"
	"github.com/liujianjiang/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
