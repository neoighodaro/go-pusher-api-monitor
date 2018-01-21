package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

//EndPointCalls - Object for storing endpoints call details
type EndPointCalls struct {
	gorm.Model
	EndPointID   uint `gorm:"index;not null"`
	RequestIP    string
	ResponseCode int
}

//SaveCall - Save the call details of an endpoint
func (ep EndPoints) SaveCall(context iris.Context) EndPointCalls {
	epCall := EndPointCalls{
		EndPointID:   ep.ID,
		RequestIP:    context.RemoteAddr(),
		ResponseCode: context.GetStatusCode(),
	}

	db.Create(&epCall)
	return epCall
}
