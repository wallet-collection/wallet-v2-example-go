package req

type ListPageReq struct {
	Page  int `form:"page" binding:"required,gte=1"`          // 页数
	Limit int `form:"limit" binding:"required,gte=1,lte=200"` // 每页返回多少
}
