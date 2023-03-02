package application

import (
	"awesomeProject/domain/entity"
	"awesomeProject/domain/repository"
)

// CloudhubBillRecordsContext 上下文层，为了引用仓储层
type CloudhubBillRecordsContext struct {
	CloudhubBillRecordsRepository repository.CloudhubBillRecordsRepository
}

//实现接口
var _ CloudhubBillRecordInterface = &CloudhubBillRecordsContext{}

func (c CloudhubBillRecordsContext) SaveOne(records *entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string) {
	return c.CloudhubBillRecordsRepository.SaveOne(records)
}

func (c CloudhubBillRecordsContext) GetById(i int64) (*entity.CloudhubBillRecords, error) {
	return c.CloudhubBillRecordsRepository.GetById(i)
}

func (c CloudhubBillRecordsContext) Query(records *entity.CloudhubBillRecords) *[]entity.CloudhubBillRecords {
	return c.CloudhubBillRecordsRepository.Query(records)
}

func (c CloudhubBillRecordsContext) Update(records *entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string) {
	return c.CloudhubBillRecordsRepository.Update(records)
}

func (c CloudhubBillRecordsContext) Delete(id int64) error {
	return c.CloudhubBillRecordsRepository.Delete(id)
}

type CloudhubBillRecordInterface interface {
	SaveOne(*entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string)

	GetById(int64) (*entity.CloudhubBillRecords, error)

	Query(*entity.CloudhubBillRecords) *[]entity.CloudhubBillRecords

	Update(*entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string)

	Delete(int64) error
}
