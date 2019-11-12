package common

import "errors"

var (
	ErrLockAlreadyRequired = errors.New("锁已被占用")
)