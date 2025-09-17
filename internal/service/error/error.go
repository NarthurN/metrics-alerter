package error

import "errors"

var (
	ErrInvalidType = errors.New("неправильный тип метрик")
)