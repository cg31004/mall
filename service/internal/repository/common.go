package repository

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"simon/mall/service/internal/model/po"
)

func parsePaging(pager *po.Pager) func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pager != nil {
			db = db.Limit(pager.GetSize()).Offset(pager.GetOffset())
		}

		return db
	}
}

func pagingMongoOption(pager *po.Pager) *options.FindOptions {
	opt := options.Find()
	if pager != nil {
		opt.SetSkip(int64(pager.GetOffset()))
		opt.SetLimit(int64(pager.GetSize()))
	}

	return opt
}

func setForUpdate(b bool) func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if b {
			db = db.Clauses(clause.Locking{Strength: "UPDATE"})
		}
		return db
	}
}
