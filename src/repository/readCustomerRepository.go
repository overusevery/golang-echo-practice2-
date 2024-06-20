package repository

import (
	"time"

	"github.com/overusevery/golang-echo-practice2/src/domain/entity"
)

type RealCustomerRepository struct {
	localData map[int]entity.Customer
}

func NewRealCustomerRepository() *RealCustomerRepository {
	return &RealCustomerRepository{
		localData: map[int]entity.Customer{
			1: {
				ID:            1,
				Name:          "山田 太郎",
				Address:       "東京都練馬区豊玉北2-13-1",
				ZIP:           "176-0013",
				Phone:         "03-1234-5678",
				MarketSegment: "個人",
				Nation:        "日本",
				Birthdate:     time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			2: {
				ID:            2,
				Name:          "佐藤 花子",
				Address:       "神奈川県横浜市中区伊勢佐木町1-1-1",
				ZIP:           "231-0021",
				Phone:         "045-222-3333",
				MarketSegment: "法人",
				Nation:        "日本",
				Birthdate:     time.Date(1990, 7, 10, 0, 0, 0, 0, time.UTC),
			},
			3: {
				ID:            3,
				Name:          "田中 麗子",
				Address:       "大阪府大阪市北区梅田1-1-1",
				ZIP:           "530-0001",
				Phone:         "06-6666-7777",
				MarketSegment: "個人",
				Nation:        "日本",
				Birthdate:     time.Date(2000, 4, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
}

func (r *RealCustomerRepository) GetCustomer(id int) entity.Customer {
	return r.localData[id]
}

func (r *RealCustomerRepository) CreateCustomer(customer entity.Customer) error {
	return nil
}
