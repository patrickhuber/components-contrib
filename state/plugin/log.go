package plugin

import (
	"github.com/dapr/kit/logger"
	"github.com/hashicorp/go-hclog"
)

type logAdapter struct {
	inner logger.Logger
}

func LogAdapter(l logger.Logger) hclog.Logger {
	intercept := hclog.NewInterceptLogger(nil)
	adapter := &logAdapter{
		inner: l,
	}
	intercept.RegisterSink(adapter)
	return intercept
}

func (a *logAdapter) Accept(name string, level hclog.Level, msg string, args ...interface{}) {
	switch level {
	case hclog.Trace:
		a.inner.Debugf(msg, args...)
	case hclog.Debug:
		a.inner.Debugf(msg, args...)
	case hclog.Info:
		a.inner.Infof(msg, args...)
	case hclog.Warn:
		a.inner.Warnf(msg, args...)
	case hclog.Error:
		a.inner.Errorf(msg, args...)
	}
}
