package repository

import (
	"awesomeProject/domain/entity"
	"awesomeProject/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
)

type CloudhubBillRecordsRepository struct {
	db *gorm.DB
}

func NewCloudhubBillRecordsRepository(db *gorm.DB) *CloudhubBillRecordsRepository {
	return &CloudhubBillRecordsRepository{db}
}

//实现接口
var _ repository.CloudhubBillRecordsRepository = &CloudhubBillRecordsRepository{}

func (r *CloudhubBillRecordsRepository) SaveOne(do *entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string) {
	//TODO implement me
	panic("implement me")
}

func (r *CloudhubBillRecordsRepository) GetById(id int64) (*entity.CloudhubBillRecords, error) {
	var cloudhubBillRecordsDo entity.CloudhubBillRecords
	err := r.db.Debug().Where("id = ?", id).Take(&cloudhubBillRecordsDo).Error
	if err != nil {
		//return nil, errors.New("database error, please try again")
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("record not found")
	}
	return &cloudhubBillRecordsDo, nil
}

func (r *CloudhubBillRecordsRepository) Query(do *entity.CloudhubBillRecords) *[]entity.CloudhubBillRecords {
	//TODO implement me
	panic("implement me")
}

func (r *CloudhubBillRecordsRepository) Update(do *entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string) {
	//TODO implement me
	panic("implement me")
}

func (r *CloudhubBillRecordsRepository) Delete(i int64) error {
	//TODO implement me
	panic("implement me")
}
