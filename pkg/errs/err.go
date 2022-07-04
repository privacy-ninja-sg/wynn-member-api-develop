package errs

const (
	NOT_FOUND               = "Not found"
	INVALID_VALIDATION      = "Invalid Validation"
	EXPIRED_TOKEN           = "Unauthorized, check expiration time of your token"
	ACCOUNT_ALREADY_CREATED = "Account is already created"
	NOT_FOUND_USER          = "user data from access token not found"
	INVALID_REQ             = "invalid request"
	UPDATE_TOKEN_DATA       = "an error update token data"
	GET_TOKEN_DATA          = "an error get token data"
	CREATE_TOKEN_DATA       = "an error create token data"
)

var errorCode = map[string]int{
	NOT_FOUND:               404,
	INVALID_REQ:             111,
	INVALID_VALIDATION:      112,
	EXPIRED_TOKEN:           113,
	ACCOUNT_ALREADY_CREATED: 114,
	NOT_FOUND_USER:          115,
	UPDATE_TOKEN_DATA:       116,
	GET_TOKEN_DATA:          117,
	CREATE_TOKEN_DATA:       118,
}

func ErrorCode(errMsg string) int {
	val, ok := errorCode[errMsg]
	if ok == false {
		return 99
	}
	return val
}
