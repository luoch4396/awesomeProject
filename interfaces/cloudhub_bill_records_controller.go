package interfaces

import (
	"awesomeProject/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CloudhubBillRecordsController 上下文层接口
type CloudhubBillRecordsController struct {
	recordsInterface application.CloudhubBillRecordInterface
}

// NewCloudhubBillRecordsController 构造函数
func NewCloudhubBillRecordsController(recordsInterface application.CloudhubBillRecordInterface) *CloudhubBillRecordsController {
	return &CloudhubBillRecordsController{
		recordsInterface: recordsInterface,
	}
}

func (cbrc *CloudhubBillRecordsController) GetById(c *gin.Context) {
	recordId, err := strconv.ParseInt(c.Param("id"), 15, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}
	record, err := cbrc.recordsInterface.GetById(recordId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, record)
}
