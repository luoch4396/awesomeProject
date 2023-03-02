package repository

import (
	"awesomeProject/domain/entity"
	"awesomeProject/domain/repository"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Repositories struct {
	CloudhubBillRecords repository.CloudhubBillRecordsRepository
	db                  *gorm.DB
}

func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", DbUser, DbPassword, DbHost, DbPort, DbName)
	fmt.Println("数据库配置字符串:", DbUrl)
	db, err := gorm.Open(DbDriver, DbUrl)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		CloudhubBillRecords: NewCloudhubBillRecordsRepository(db),
		db:                  db,
	}, nil
}

// Close closes the database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// AutoMigrate 自动增加实体类中的字段
func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.CloudhubBillRecords{}).Error
}
