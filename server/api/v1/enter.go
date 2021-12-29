package v1

import (
	"github.com/liujianjiang/gin-vue-admin/server/api/v1/autocode"
	"github.com/liujianjiang/gin-vue-admin/server/api/v1/example"
	"github.com/liujianjiang/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
