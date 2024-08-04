package entityutil

type newEntitiyFun[I any, E any] func(input I) (E, error)

// WrapNew wraps a wrapped New function, adding error handling and returning a
// function that only returns the entity. The provided errorList is populated
// with any validation errors encountered during the entity creation.
func WrapNew[I any, E any](new newEntitiyFun[I, E], errorList *[]error) func(input I) E {
	return func(input I) E {
		entity, err := new(input)
		if err != nil {
			*errorList = append(*errorList, err)
		}
		return entity
	}
}
