package model

type ParamRegister struct {
	UserName string `json:"username,omitempty" binding:"required,min=5,max=15"`
	Password string `json:"password,omitempty" binding:"required,min=6,max=15"`
	Gender   *bool  `json:"gender"`
	Phone    string `json:"phone"`
	Email    string `json:"email" binding:"email"`
}

type ParamLogin struct {
	UserName   string `json:"username,omitempty" binding:"required,min=5,max=15"`
	Password   string `json:"password,omitempty" binding:"required,min=6,max=15"`
	RePassword string `json:"re_password,omitempty" binding:"required,eqfield=Password"`
}

type ParamCreatePost struct {
	Title   string `json:"title,omitempty" binding:"min=3,max=128,required"`
	Content string `json:"content,omitempty" binding:"min=1,max=8192,required"`
}

type ParamCreateCommunity struct {
	CommunityName string `json:"community_name,omitempty" binding:"required,max=64"`
	Introduction  string `json:"introduction,omitempty" binding:"required,max=128"`
}
