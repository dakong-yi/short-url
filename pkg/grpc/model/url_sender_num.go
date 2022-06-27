package model

import (
	"context"
	"short-url/pkg/database/mysql"
	"time"
)

type UrlSenderNum struct {
	ID         int       `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	StartNum   int64     `gorm:"column:start_num;default:0;NOT NULL"` // 开始号
	EndNum     int64     `gorm:"column:end_num;default:0;NOT NULL"`   // 结束号
	Version    int       `gorm:"column:version;default:0;NOT NULL"`   // 版本号
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP"`
}

func (m *UrlSenderNum) TableName() string {
	return "url_sender_num"
}

// GetMaxNum 获取最大的号
func (m *UrlSenderNum) GetMaxNum(ctx context.Context) error {
	return mysql.MysqlClient.Last(m).Error
}

func (m *UrlSenderNum) Create(ctx context.Context) error {
	return mysql.MysqlClient.Create(m).Error
}

// UpdateByVersion 获取新的号码
func (m *UrlSenderNum) UpdateByVersion(ctx context.Context, version int) error {
	return mysql.MysqlClient.Where("id = ? and version = ?", m.ID, version).Updates(m).Error
}
