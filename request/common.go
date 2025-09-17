package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

type PageInfoV2 struct {
	Page     int64  `json:"page" form:"page"`         // 页码
	PageSize int64  `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	Id int `json:"id" form:"id"` // 主键ID
}
type GetByIdV2 struct {
	Id int64 `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() int {
	return r.Id
}
func (r *GetByIdV2) Uint() int64 {
	return r.Id
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

type IdsReqV2 struct {
	Ids []int64 `json:"ids" form:"ids" binding:"required"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId int `json:"authorityId" form:"authorityId"` // 角色ID
}
type GetAuthorityIdV2 struct {
	AuthorityId int64 `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
