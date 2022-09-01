package gerrors

var (
	ErrDB   = NewCodeMsg(100001, "operation DB error.")
	ErrAuth = NewCodeMsg(100002, "authentication error.")
	ErrCall = NewCodeMsg(100003, "call third party error.")
)
