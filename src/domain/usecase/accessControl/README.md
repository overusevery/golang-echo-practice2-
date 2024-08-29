# Where to place access control functionality

usecase itself has access control in my this *experimental* implemention. But I feel handler wrapper style is prefer.

## How to use
### example
Execute wraps execute(main logic) with access control.
This is written with python's style decorator pattern in mind.
example 
```
func (uc *UseCase) Execute(ctx context.Context, id string) (*entity.Customer, error) {
	if accesscontrol.New("needed permission").IsNotAllowed(ctx) {
		return nil, errors.New("not enough scope")
	}
	return uc.execute(ctx, id)
}
func (uc *UseCase) execute(ctx context.Context, id string) (*entity.Customer, error) {
    //actual business logics
	return res, err
}
```