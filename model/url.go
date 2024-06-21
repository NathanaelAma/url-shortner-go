package model

type Url struct {
	Short string `json:"short"`
	Long  string `json:"long" binding:"required,url"`
}
