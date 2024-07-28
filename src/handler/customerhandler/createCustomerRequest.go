package customerhandler

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

type CreateCustomerRequest struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	ZIP           string    `json:"zip"`
	Phone         string    `json:"phone"`
	MarketSegment string    `json:"mktsegment"`
	Nation        string    `json:"nation"`
	Birthdate     time.Time `json:"birthdate"`
}

func (r *CreateCustomerRequest) ConvertFrom() entity.Customer {
	return entity.Customer{
		ID:            r.ID,
		Name:          r.Name,
		Address:       r.Address,
		ZIP:           r.ZIP,
		Phone:         r.Phone,
		MarketSegment: r.MarketSegment,
		Nation:        r.Nation,
		Birthdate:     r.Birthdate,
	}
}
