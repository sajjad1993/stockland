package presenter

const (
	UserValueKey      = "user"
	VerifyCodeInvalid = "VerifyCodeInvalid"
	UserUpdated       = "UserUpdated"
	LoginFailed       = "LoginFailed"
	LoginForbidden    = "LoginForbidden"
)

var ResponseMessages = map[string]string{

	VerifyCodeInvalid: "verify code type is invalid",     //code:13
	UserUpdated:       "user updated successfully",       //code:12
	LoginFailed:       "username or password is invalid", //code:10
	LoginForbidden:    "user access denied",              //code:11
}
