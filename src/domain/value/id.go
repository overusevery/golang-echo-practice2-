package value

type ID struct {
	value        int
	isDetermined bool
}

func NewID(id int) *ID {
	return &ID{value: id, isDetermined: true}
}

func NewUndeteminedID() *ID {
	return &ID{value: 0, isDetermined: false}
}

func (i *ID) GetValue() (int, bool) {
	return i.value, i.isDetermined
}
