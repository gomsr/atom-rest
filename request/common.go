package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	Id int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() int {
	return int(r.Id)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId int `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
