package models

import (
	"github.com/jinzhu/gorm"
)

//EndPoints - endpoint model
type EndPoints struct {
	gorm.Model
	Name, URL string
	Type      string          `gorm:"DEFAULT:'GET'"`
	Calls     []EndPointCalls `gorm:"ForeignKey:EndPointID"`
}

//EndPointWithCallSummary - Endpoint with last call summary
type EndPointWithCallSummary struct {
	ID            uint
	Name, URL     string
	Type          string
	LastStatus    int
	NumRequests   int
	LastRequester string
}

//GetWithCallSummary - get all endpoints with call summary details
func (ep EndPoints) GetWithCallSummary() []EndPointWithCallSummary {
	var eps []EndPoints
	var epsWithDets []EndPointWithCallSummary

	db.Preload("Calls").Find(&eps)

	for _, elem := range eps {
		calls := elem.Calls
		lastCall := calls[len(calls)-1:][0]

		newElem := EndPointWithCallSummary{
			elem.ID,
			elem.Name,
			elem.URL,
			elem.Type,
			lastCall.ResponseCode,
			len(calls),
			lastCall.RequestIP,
		}

		epsWithDets = append(epsWithDets, newElem)
	}

	return epsWithDets
}

//SaveOrCreate - save endpoint called
func (ep EndPoints) SaveOrCreate() EndPoints {
	db.FirstOrCreate(&ep, ep)
	return ep
}
