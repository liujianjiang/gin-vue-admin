package response

import "github.com/liujianjiang/gin-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
