package request

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func DefaultPageInfo() PageInfo {
	return PageInfo{
		Page:     1,
		PageSize: 10,
	}
}

func DefaultPageSearch() PageSearch {
	return PageSearch{
		PageInfo: DefaultPageInfo(),
	}
}

type PageSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PageInfo
}

func QueryPage(info PageSearch, db *gorm.DB) (*gorm.DB, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	total := new(int64)
	err := db.Count(total).Error
	if err != nil {
		return nil, 0, err
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	return db, *total, nil
}

func DefaultMemberPageSearch(uid uuid.UUID) MemberPageSearch {
	return MemberPageSearch{
		MemberId: uid,
		PageSearch: PageSearch{
			PageInfo: DefaultPageInfo(),
		},
	}
}

type MemberPageSearch struct {
	PageSearch
	MemberId uuid.UUID `json:"memberId" form:"memberId" `
}

func QueryPageWithMember(info MemberPageSearch, db *gorm.DB) (*gorm.DB, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MemberId != uuid.Nil {
		db = db.Where("member_id = ?", info.MemberId)
	}

	total := new(int64)
	err := db.Count(total).Error
	if err != nil {
		return nil, 0, err
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	return db, *total, nil
}
