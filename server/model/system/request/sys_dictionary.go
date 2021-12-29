package request

import (
	"github.com/liujianjiang/gin-vue-admin/server/model/common/request"
	"github.com/liujianjiang/gin-vue-admin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
