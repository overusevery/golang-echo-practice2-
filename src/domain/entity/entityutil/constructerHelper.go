package entityutil

import "github.com/overusevery/golang-echo-practice2/src/shared/util"

type newEntitiyFun[I any, E any] func(input I) (E, util.ErrorList)

// WrapNew wraps a wrapped New function, adding error handling and returning a
// function that only returns the entity. The provided errorList is populated
// with any validation errors encountered during the entity creation.
func WrapNew[I any, E any](new newEntitiyFun[I, E], errorList *util.ErrorList) func(input I) E {
	return func(input I) E {
		entity, validationErrorList := new(input)
		*errorList = errorList.Concatenate(&validationErrorList)
		return entity
	}
}
