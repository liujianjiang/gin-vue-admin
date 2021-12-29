package system

type Stock struct {
	Code       int    `json:"code"`
	Estimate   string `json:"estimate"`
	Name       string `json:"name"`
	Createtime string `json:"createtime"`
	Updatetime string `json:"updatetime"`
}
