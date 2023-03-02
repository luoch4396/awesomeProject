package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalDateTime time.Time

// MarshalJSON 格式化输出日期字符串yyyy-MM-dd HH:mm:ss
func (t *LocalDateTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Value 解决插入时间值时，默认为空
func (t LocalDateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalDateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalDateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
