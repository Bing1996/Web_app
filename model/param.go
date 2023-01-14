package model

type ParamRegister struct {
	UserName   string `json:"username,omitempty" binding:"required,min=5,max=15"`
	Password   string `json:"password,omitempty" binding:"required,min=6,max=15"`
	RePassword string `json:"re_password,omitempty" binding:"required,eqfield=Password"`
}
