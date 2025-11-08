// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"errors"
	"fmt"
)

var sugar *SugaredLogger

type PanicInfo struct {
	Stack, Panic string
}

func init() {
	sugar = New(NewConfig())

	// SetPanicHandler(func(a interface{}, stack []byte) {
	//	info := PanicInfo{}
	//	info.Panic = fmt.Sprint(a)
	//	info.Stack = string(stack)
	//	bytes, _ := json.Marshal(info)
	//	Error(string(bytes))
	// })
}

func Logger() *SugaredLogger {
	return sugar
}

func SetLevelFrom(level string) {
	sugar.AtomicLevel.SetLevel(ParseLogLevel(level))
}

func SetLevel(level Level) {
	sugar.AtomicLevel.SetLevel(level)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Debug(args ...interface{}) {
	Logger().Debug(args...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Info(args ...interface{}) {
	Logger().Info(args...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Warn(args ...interface{}) {
	Logger().Warn(args...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the Logger.
func Error(args ...interface{}) {
	Logger().Error(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	Logger().Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	Logger().Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	Logger().Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	Logger().Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	Logger().Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger().Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	Logger().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Logger().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	Logger().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Logger().Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger().Fatalw(msg, keysAndValues...)
}

func NewError(args ...interface{}) error {
	Logger().Error(args...)
	return errors.New(fmt.Sprint(args...))
}

func NewErrorf(template string, args ...interface{}) error {
	Logger().Errorf(template, args...)
	return fmt.Errorf(template, args...)
}

func NewErrorw(msg string, keysAndValues ...interface{}) error {
	Logger().Errorw(msg, keysAndValues...)
	return errors.New(Logger().FormatMessage(msg, keysAndValues...))
}
