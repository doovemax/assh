// Code generated by github.com/doovemax/assh/contrib/generate-loggers.sh

package main

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("assh.contrib.webapp")
}
