package entity

import (
	"awesomeProject/utils"
)

type CloudhubBillRecords struct {
	ID           int64                `gorm:"primary_key" json:"id"`
	PlatformNo   uint32               `gorm:"size:50;null;" json:"platform_no"`
	SellerNick   string               `gorm:"size:100;null" json:"seller_nick"`
	RecordNo     string               `gorm:"size:100;null;" json:"record_no"`
	RecordType   uint32               `gorm:"size:255;null;" json:"record_type"`
	RecordStatus string               `gorm:"size:255;null" json:"record_status"`
	RecordMd5    string               `gorm:"size:255;null" json:"record_md5"`
	ModifiedDate *utils.LocalDateTime `json:"modified_date"`
	Topic        string               `gorm:"size:255;null;" json:"topic"`
	MsgKey       string               `gorm:"size:100;null;" json:"msg_key"`
	Tag          string               `gorm:"size:100;null;" json:"tag"`
	IsExtra      bool                 `json:"is_extra"`
	InitDate     *utils.LocalDateTime `gorm:"default:CURRENT_TIMESTAMP" json:"init_date"`
}

//func (ch *CloudhubBillRecords) Prepare() {
//	ch.InitDate = time.Now()
//}
