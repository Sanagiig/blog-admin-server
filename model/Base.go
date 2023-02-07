package model

import (
	"database/sql/driver"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

const formatTimeStr = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(formatTimeStr))), nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(formatTimeStr, timeStr)
	*t = LocalTime(t1)
	return err
}

func (t *LocalTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func (t LocalTime) Value() (driver.Value, error) {
	return time.Time(t).Format(formatTimeStr), nil
}

func (t *LocalTime) Scan(v []byte) error {
	vt, err := time.ParseInLocation(formatTimeStr, string(v), time.Local)
	*t = LocalTime(vt)
	return err
}

type Base struct {
	ID        string         `gorm:"type:char(36);primarykey" json:"id,omitempty"`
	CreatedAt *LocalTime     `gorm:"type:datetime" json:"createdAt,omitempty"`
	UpdatedAt *LocalTime     `gorm:"type:datetime" json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	if base.ID == "" {
		id := uuid.NewV4()
		t := LocalTime(time.Now())
		base.ID = id.String()
		base.CreatedAt = &t
		base.UpdatedAt = &t
	}
	return nil
}

func (base *Base) BeforeUpdate(db *gorm.DB) error {
	t := LocalTime(time.Now())
	base.UpdatedAt = &t
	return nil
}

func (base *Base) BeforeDelete(db *gorm.DB) error {
	base.DeletedAt = gorm.DeletedAt{
		Time: time.Now(),
	}
	return nil
}
