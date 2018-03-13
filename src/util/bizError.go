package util

type BizError struct {
	Msg string
}

func (e *BizError) Error() string {
	return e.Msg
}
