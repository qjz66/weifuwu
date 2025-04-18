package models

import (
	"time"
)

// User 鐢ㄦ埛琛
type User struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null;comment:'鐢ㄦ埛ID'"`                      // 鐢ㄦ埛ID
	Username   string    `gorm:"column:username;type:varchar(50);not null;default:'';comment:'鐢ㄦ埛鍚'"`                                        // 鐢ㄦ埛鍚
	Password   string    `gorm:"column:password;type:varchar(50);not null;default:'';comment:'鐢ㄦ埛瀵嗙爜锛孧D5鍔犲瘑'"`                               // 鐢ㄦ埛瀵嗙爜锛孧D5鍔犲瘑
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'鍒涘缓鏃堕棿'"`                      // 鍒涘缓鏃堕棿
	UpdateTime time.Time `gorm:"index:ix_update_time;column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'鏇存柊鏃堕棿'"` // 鏇存柊鏃堕棿
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID         string
	Username   string
	Password   string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Username:   "username",
	Password:   "password",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}
