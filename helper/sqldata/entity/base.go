package entity

import (
	"github/tronglv_authen_author/helper/auth"
	"github/tronglv_authen_author/helper/util"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type IdModel struct {
	Id           uint32    `gorm:"primary_key:auto_increment"`
	UId          string    `gorm:"uniqueIndex;not null;type:varchar(21)"`
	IsDeleted    bool      `gorm:"column:is_deleted"`
	CreatedByUId string    `gorm:"column:created_by_uid;type:varchar(36);index:idx_created_by_uid;"`
	CreatedBy    string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedByUId string    `gorm:"column:updated_by_uid;type:varchar(36)"`
	UpdatedBy    string    `gorm:"column:updated_by;type:varchar(255)"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (base *IdModel) BeforeCreate(tx *gorm.DB) error {
	uid, _ := gonanoid.New()
	if obj, err := auth.GetAuthData(tx.Statement.Context); err == nil {
		tx.Statement.SetColumn("CreatedBy", obj.GetName())
		tx.Statement.SetColumn("CreatedByUId", obj.GetUid())
		tx.Statement.SetColumn("UpdatedBy", obj.GetName())
		tx.Statement.SetColumn("UpdatedByUId", obj.GetUid())
	}

	tx.Statement.SetColumn("UId", uid)
	tx.Statement.SetColumn("IsDeleted", false)
	if base.CreatedAt.IsZero() {
		tx.Statement.SetColumn("CreatedAt", util.TimeNow())
	}
	if base.UpdatedAt.IsZero() {
		tx.Statement.SetColumn("UpdatedAt", util.TimeNow())
	}
	return nil
}

func (base *IdModel) BeforeUpdate(tx *gorm.DB) error {
	if obj, err := auth.GetAuthData(tx.Statement.Context); err == nil {
		tx.Statement.SetColumn("UpdatedBy", obj.GetName())
		tx.Statement.SetColumn("UpdatedByUId", obj.GetUid())
	}
	tx.Statement.SetColumn("UpdatedAt", util.TimeNow())
	return nil
}
