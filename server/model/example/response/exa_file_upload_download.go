package response

import "github.com/liujianjiang/gin-vue-admin/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
