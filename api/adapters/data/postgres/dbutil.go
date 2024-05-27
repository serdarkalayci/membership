package postgres

import "gorm.io/gorm"

// Paginate is a scope applied to a gorm.DB object to paginate the results
func Paginate(pageSize, pageNum int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
	  offset := (pageNum - 1) * pageSize
	  return db.Offset(offset).Limit(pageSize)
	}
  }
  