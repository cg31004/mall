package timelogger

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const ContextKey = "timeLogger"

var (
	err_TimeLoggerTypeError = errors.New("必須包含 *timeLogger 進 context 才能正常使用 TimeLog")
)

type timeLogger struct {
	startTime time.Time

	funcNames []string
	logMap    map[string]time.Duration
	mux       sync.RWMutex
}

func NewTimeLogger() *timeLogger {
	return &timeLogger{
		startTime: time.Now(),
		logMap:    make(map[string]time.Duration),
	}
}

func LogTime(ctx context.Context) func() {
	timeLogger, existed := ctx.Value(ContextKey).(*timeLogger)
	if !existed {
		return func() {}
	}
	startTime := time.Now()
	return func() {
		timeLogger.logFuncTime(getFuncName(), time.Since(startTime))
	}
}

func GetTotalDuration(ctx context.Context) (time.Duration, error) {
	timeLogger, existed := ctx.Value(ContextKey).(*timeLogger)
	if !existed {
		return time.Duration(0), err_TimeLoggerTypeError
	}

	return timeLogger.getTotalDuration(), nil
}

func GetTimeLogs(ctx context.Context) ([]string, error) {
	timeLogger, existed := ctx.Value(ContextKey).(*timeLogger)
	if !existed {
		return nil, err_TimeLoggerTypeError
	}

	logs := make([]string, 0, len(timeLogger.funcNames))
	for _, funcName := range timeLogger.funcNames {
		logs = append(logs, fmt.Sprintf("[+%s] %s", timeLogger.getDuration(funcName), funcName))
	}

	return logs, nil
}

func (logger *timeLogger) getTotalDuration() time.Duration {
	return time.Since(logger.startTime)
}

func (logger *timeLogger) logFuncTime(funcName string, duration time.Duration) {
	logger.mux.Lock()
	defer logger.mux.Unlock()

	_, existed := logger.logMap[funcName]
	if !existed {
		logger.funcNames = append(logger.funcNames, funcName)
	}

	logger.logMap[funcName] += duration
}

func (logger *timeLogger) getDuration(funcName string) time.Duration {
	logger.mux.RLock()
	defer logger.mux.RUnlock()

	duration := logger.logMap[funcName]
	return duration
}

func getFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName[strings.LastIndex(funcName, "/")+1:]
}
