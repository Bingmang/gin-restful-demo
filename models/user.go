package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"log"
)

type Coach struct {
	gorm.Model
	Account string				`gorm:"type:varchar(64);unique;not null" json:"account"`
	Password string  			`gorm:"type:varchar(64);not null" json:"password"`
	Majors pq.StringArray 		`gorm:"type:varchar(64)[]" json:"majors"`
	Classes pq.StringArray 		`gorm:"type:varchar(64)[]" json:"classes"`
}

type Student struct {
	gorm.Model
	Account string				`gorm:"type:varchar(64);unique;not null" json:"account"`
	Password string  			`gorm:"type:varchar(64);not null" json:"password"`
	Major string				`gorm:"type:varchar(64);not null" json:"major"`
	Class string				`gorm:"type:varchar(64);not null" json:"class"`
	Coach Coach					`gorm:"foreignkey:coach_id"`
	CoachId uint				`gorm:"column:coach_id" json:"coach_id"`
}

func init() {
	log.Println("Init user table...")
	if !DB.HasTable(Coach{}) {
		DB.CreateTable(Coach{})
	}
	if !DB.HasTable(Student{}) {
		DB.CreateTable(Student{})
	}
}

func (c *Coach) TableName() string { return "coach" }
func (c *Student) TableName() string { return "student" }

func (c *Coach) Insert() error {
	return DB.Create(c).Error
}

func (s *Student) Insert() error {
	return DB.Create(s).Error
}
