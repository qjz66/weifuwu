package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UserMgr) Debug() *_UserMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 鐢ㄦ埛ID
func (obj *_UserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 鐢ㄦ埛鍚
func (obj *_UserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 鐢ㄦ埛瀵嗙爜锛孧D5鍔犲瘑
func (obj *_UserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithCreateTime create_time获取 鍒涘缓鏃堕棿
func (obj *_UserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 鏇存柊鏃堕棿
func (obj *_UserMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 鐢ㄦ埛ID
func (obj *_UserMgr) GetFromID(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 鐢ㄦ埛ID
func (obj *_UserMgr) GetBatchFromID(ids []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 鐢ㄦ埛鍚
func (obj *_UserMgr) GetFromUsername(username string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 鐢ㄦ埛鍚
func (obj *_UserMgr) GetBatchFromUsername(usernames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 鐢ㄦ埛瀵嗙爜锛孧D5鍔犲瘑
func (obj *_UserMgr) GetFromPassword(password string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 鐢ㄦ埛瀵嗙爜锛孧D5鍔犲瘑
func (obj *_UserMgr) GetBatchFromPassword(passwords []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 鍒涘缓鏃堕棿
func (obj *_UserMgr) GetFromCreateTime(createTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 鍒涘缓鏃堕棿
func (obj *_UserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 鏇存柊鏃堕棿
func (obj *_UserMgr) GetFromUpdateTime(updateTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 鏇存柊鏃堕棿
func (obj *_UserMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIxUpdateTime  获取多个内容
func (obj *_UserMgr) FetchIndexByIxUpdateTime(updateTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}
