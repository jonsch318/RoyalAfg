package authentication

import (
	"github.com/jonsch318/royalafg/pkg/mw"
)

// VerifyAuthentication could check the users session in regards to authentication or authorization.
// This is currently a NoOp because all necessary actions are done by the http handler itself.
// We could also connect to other services if needed.
func (auth *Authentication) VerifyAuthentication(user *mw.UserClaims) bool {
	return true
}
