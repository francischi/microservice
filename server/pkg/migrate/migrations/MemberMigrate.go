package migrations

import (
	"server/pkg/repos/models"
	"gorm.io/gorm"
)

func Member(db *gorm.DB){
	var MemberModel models.MemberModel
	db.AutoMigrate(&MemberModel)
}