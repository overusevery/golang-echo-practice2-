package customerhandler

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

type CreateCustomerResponse struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	ZIP           string    `json:"zip"`
	Phone         string    `json:"phone"`
	MarketSegment string    `json:"mktsegment"`
	Nation        string    `json:"nation"`
	Birthdate     time.Time `json:"birthdate"`
}

func convertToCreateCustomerResponse(customer entity.Customer) CreateCustomerResponse {
	return CreateCustomerResponse{
		ID:            customer.ID,
		Name:          customer.Name,
		Address:       customer.Address,
		ZIP:           customer.ZIP,
		Phone:         customer.Phone,
		MarketSegment: customer.MarketSegment,
		Nation:        customer.Nation,
		Birthdate:     customer.Birthdate,
	}
}
