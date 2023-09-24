package errors

import (
	"errors"
	"fmt"
)

// ErrIndexOutOfRange 下标超过范围错误
func ErrIndexOutOfRange(length, index int) error {
	return fmt.Errorf("tekila: 下标超过范围, Len %d, Index %d", length, index)
}

// ErrInvalidType 类型转化失败错误
func ErrInvalidType(src, dst string) error {
	return fmt.Errorf("tekila: 类型转换失败, Src: %s, Dst: %s", src, dst)
}

// NewError 创建一个新的自定义错误
func NewError(msg string) error {
	return errors.New(msg)
}
