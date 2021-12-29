package request

import (
	"github.com/liujianjiang/gin-vue-admin/server/model/autocode"
	"github.com/liujianjiang/gin-vue-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}