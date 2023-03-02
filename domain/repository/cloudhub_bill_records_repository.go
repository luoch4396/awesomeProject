package repository

import "awesomeProject/domain/entity"

type CloudhubBillRecordsRepository interface {
	SaveOne(*entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string)

	GetById(int64) (*entity.CloudhubBillRecords, error)

	Query(*entity.CloudhubBillRecords) *[]entity.CloudhubBillRecords

	Update(*entity.CloudhubBillRecords) (*entity.CloudhubBillRecords, map[string]string)

	Delete(int64) error
}
