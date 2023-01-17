package datastore

import "errors"

var (
	ErrTableNotExist     = errors.New("table not exist")
	ErrGuestNotExist     = errors.New("guest not exist")
	ErrGuestAlreadyExist = errors.New("guest already exist")
)

type logger interface {
	Panicf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}
