package model

type Url struct {
	ID    string `json:"id" binding:"required"`
	Short string `json:"short"`
	Long  string `json:"long" binding:"required,url"`
}
