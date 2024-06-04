type Customer struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	ZIP           string    `json:"zip"`
	Phone         string    `json:"phone"`
	MarketSegment string    `json:"mktsegment"`
	Nation        string    `json:"nation"`
	Birthdate     time.Time `json:"birthdate"`
}