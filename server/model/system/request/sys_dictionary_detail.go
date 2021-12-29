package request

import (
	"github.com/liujianjiang/gin-vue-admin/server/model/common/request"
	"github.com/liujianjiang/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
