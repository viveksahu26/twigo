package errors

type TwigoError struct {
	Message string
	Code    int
}

func Error(twigoError TwigoError) error {
	return &TwigoError{
		Message: twigoError.Message,
		Code:    twigoError.Code,
	}
}

// Assert that we implement error at build time.
var _ error = (*TwigoError)(nil)

// Error implements error
func (te *TwigoError) Error() string {
	return te.Message
}

func (te *TwigoError) ExitCode() int {
	return te.Code
}
