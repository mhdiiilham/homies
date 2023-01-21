package common

import "errors"

var (
	ErrConfigNotFound         = errors.New("config file not found")
	ErrFailedMarshalingConfig = errors.New("failed marshalling config")
)
