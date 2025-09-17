package error

import "errors"

var (
	ErrInvalidType = errors.New("Неправильный тип метрик")
)