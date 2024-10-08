package accesscontrol

import (
	"context"

	"github.com/overusevery/golang-echo-practice2/src/shared/message"
)

var (
	ErrNotEnoughScope = message.ERRID00008
)

type AccessControl struct {
	allowlist []string
}

func New(allowlist ...string) AccessControl {
	return AccessControl{allowlist: allowlist}
}

func (ac AccessControl) IsAllowed(ctx context.Context) bool {
	scopelist := ctx.Value("scope").([]string)
	for _, a := range ac.allowlist {
		for _, s := range scopelist {
			if s == a {
				return true
			}
		}
	}
	return false
}

func (ac AccessControl) IsNotAllowed(ctx context.Context) bool {
	return !ac.IsAllowed(ctx)
}
