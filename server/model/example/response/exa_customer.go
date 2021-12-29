package response

import "github.com/liujianjiang/gin-vue-admin/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
