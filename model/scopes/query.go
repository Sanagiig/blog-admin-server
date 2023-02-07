package scopes

import (
	"fmt"
	"gorm.io/gorm"
)

type PagerInfo struct {
	Page     int
	PageSize int
}

func Pager(page int, pageSize int, count *int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Count(count)
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

// AllParentID
//
//	@Description: 通过 存储过程 获取所有父节点ID (包含传入的ID)
//	@param tbName
//	@param id
//	@return func(db *gorm.DB) *gorm.DB
func AllParentID(tbName string, id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Exec("SET @parentIDs = '';").
			Exec("CALL getAllParentID('" + id + "','category',@parentIDs);").
			Exec(fmt.Sprintf("CALL getAllParentID('%s','%s',@parentIDs);", id, tbName)).
			Raw("Select @parentIDs;")
	}
}

func AllChildrenID(tbName string, id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Exec("SET @childrenIDs = '';").
			Exec("CALL getAllChildrenID('" + id + "','category',@childrenIDs);").
			Exec(fmt.Sprintf("CALL getAllChildrenID('%s','%s',@childrenIDs);", id, tbName)).
			Raw("Select @childrenIDs;")
	}
}
