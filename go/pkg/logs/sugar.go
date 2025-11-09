package logs

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
	_multipleErrMsg     = "Multiple errors without a key."
)

var (
	defaultConfig = zapcore.EncoderConfig{
		TimeKey:             "time",
		LevelKey:            "Level",
		NameKey:             "Logger",
		CallerKey:           "caller",
		MessageKey:          "message",
		StacktraceKey:       "stacktrace",
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.LowercaseLevelEncoder,
		EncodeTime:          zapcore.ISO8601TimeEncoder,
		EncodeDuration:      zapcore.SecondsDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		NewReflectedEncoder: jsoniterReflectedEncoder,
	}
)

type SugaredLogger struct {
	Logger      *zap.SugaredLogger
	AtomicLevel zap.AtomicLevel
}

// New level string, encode string, port int, pattern string, initFields map[string]interface{}
func New(cfg *Config) *SugaredLogger {
	if cfg == nil {
		cfg = &Config{
			Level:  "debug",
			Encode: "console",
			Output: "console",
		}
	}

	l := &SugaredLogger{}

	var opts []zap.Option
	opts = append(opts, zap.Development())
	opts = append(opts, zap.AddCaller())
	opts = append(opts, zap.AddCallerSkip(1))
	opts = append(opts, zap.AddStacktrace(zap.FatalLevel))

	l.AtomicLevel = zap.NewAtomicLevel()
	l.AtomicLevel.SetLevel(ParseLogLevel(cfg.Level))

	if len(cfg.LevelPattern) > 0 && cfg.LevelPort > 0 {
		http.HandleFunc(cfg.LevelPattern, l.AtomicLevel.ServeHTTP)
		go func() {
			fmt.Printf("Level serve on port:%d\nusage: [GET] curl http://localhost:%d%s\nusage: [PUT] curl -XPUT --data '{\"Level\":\"debug\"}' http://localhost:%d%s\n", cfg.LevelPort, cfg.LevelPort, cfg.LevelPattern, cfg.LevelPort, cfg.LevelPattern)
			if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.LevelPort), nil); err != nil {
				panic(err)
			}
		}()
	}

	cores := make([]zapcore.Core, 0)
	output := strings.ToLower(cfg.Output)
	if strings.Contains(output, "console") {
		cores = append(cores, initConsoleCore(cfg.Encode, l.AtomicLevel))
	}
	if strings.Contains(output, "file") {
		cores = append(cores, initFileCore(cfg.File, l.AtomicLevel))
	}

	core := zapcore.NewTee(cores...)

	initFields := cfg.InitFields
	logger := zap.New(core)
	if len(initFields) > 0 {
		initFieldList := make([]zap.Field, 0)
		for k, v := range initFields {
			var field zap.Field
			if _, ok := v.(proto.Message); ok {
				field = zap.Field{Key: k, Type: zapcore.ReflectType, Interface: v}
			} else {
				field = zap.Any(k, v)
			}
			initFieldList = append(initFieldList, field)
		}
		logger = logger.With(initFieldList...)
	}

	l.Logger = logger.WithOptions(opts...).Sugar()
	return l
}

// Desugar unwraps a SugaredLogger, exposing the original zap.Logger.
func (s *SugaredLogger) Desugar() *zap.Logger {
	return s.Logger.Desugar()
}

// Named adds a sub-scope to the logger's name. See Logger.Named for details.
func (s *SugaredLogger) Named(name string) *SugaredLogger {
	return &SugaredLogger{Logger: s.Logger.Named(name)}
}

// WithOptions clones the current SugaredLogger, applies the supplied Options,
// and returns the result. It's safe to use concurrently.
func (s *SugaredLogger) WithOptions(opts ...zap.Option) *SugaredLogger {
	return &SugaredLogger{Logger: s.Logger.WithOptions(opts...)}
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//
//	 sugaredLogger.With(
//	   "hello", "world",
//	   "failure", errors.New("oh no"),
//	   Stack(),
//	   "count", 42,
//	   "user", User{Name: "alice"},
//	)
//
// is the equivalent of
//
//	unsugared.With(
//	  String("hello", "world"),
//	  String("failure", "oh no"),
//	  Stack(),
//	  Int("count", 42),
//	  Object("user", User{Name: "alice"}),
//	)
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func (s *SugaredLogger) With(args ...interface{}) *SugaredLogger {
	return &SugaredLogger{Logger: s.Logger.With(args...)}
}

// Level reports the minimum enabled level for this logger.
//
// For NopLoggers, this is [zapcore.InvalidLevel].
func (s *SugaredLogger) Level() zapcore.Level {
	return s.Logger.Level()
}

// Debug logs the provided arguments at [DebugLevel].
// Spaces are added between arguments when neither is a string.
func (s *SugaredLogger) Debug(args ...interface{}) {
	s.log(zap.DebugLevel, "", args, nil)
}

// Info logs the provided arguments at [InfoLevel].
// Spaces are added between arguments when neither is a string.
func (s *SugaredLogger) Info(args ...interface{}) {
	s.log(zap.InfoLevel, "", args, nil)
}

// Warn logs the provided arguments at [WarnLevel].
// Spaces are added between arguments when neither is a string.
func (s *SugaredLogger) Warn(args ...interface{}) {
	s.log(zap.WarnLevel, "", args, nil)
}

// Error logs the provided arguments at [ErrorLevel].
// Spaces are added between arguments when neither is a string.
func (s *SugaredLogger) Error(args ...interface{}) {
	s.log(zap.ErrorLevel, "", args, nil)
}

// DPanic logs the provided arguments at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are added between arguments when neither is a string.
func (s *SugaredLogger) DPanic(args ...interface{}) {
	s.log(zap.DPanicLevel, "", args, nil)
}

// Panic constructs a message with the provided arguments and panics.
// Spaces are added between arguments when neither is a string.
func (s *SugaredLogger) Panic(args ...interface{}) {
	s.log(zap.PanicLevel, "", args, nil)
}

// Fatal constructs a message with the provided arguments and calls os.Exit.
// Spaces are added between arguments when neither is a string.
func (s *SugaredLogger) Fatal(args ...interface{}) {
	s.log(zap.FatalLevel, "", args, nil)
}

// Debugf formats the message according to the format specifier
// and logs it at [DebugLevel].
func (s *SugaredLogger) Debugf(template string, args ...interface{}) {
	s.log(zap.DebugLevel, template, args, nil)
}

// Infof formats the message according to the format specifier
// and logs it at [InfoLevel].
func (s *SugaredLogger) Infof(template string, args ...interface{}) {
	s.log(zap.InfoLevel, template, args, nil)
}

// Warnf formats the message according to the format specifier
// and logs it at [WarnLevel].
func (s *SugaredLogger) Warnf(template string, args ...interface{}) {
	s.log(zap.WarnLevel, template, args, nil)
}

// Errorf formats the message according to the format specifier
// and logs it at [ErrorLevel].
func (s *SugaredLogger) Errorf(template string, args ...interface{}) {
	s.log(zap.ErrorLevel, template, args, nil)
}

// DPanicf formats the message according to the format specifier
// and logs it at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
func (s *SugaredLogger) DPanicf(template string, args ...interface{}) {
	s.log(zap.DPanicLevel, template, args, nil)
}

// Panicf formats the message according to the format specifier
// and panics.
func (s *SugaredLogger) Panicf(template string, args ...interface{}) {
	s.log(zap.PanicLevel, template, args, nil)
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func (s *SugaredLogger) Fatalf(template string, args ...interface{}) {
	s.log(zap.FatalLevel, template, args, nil)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {
	s.log(zap.DebugLevel, msg, nil, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{}) {
	s.log(zap.InfoLevel, msg, nil, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{}) {
	s.log(zap.WarnLevel, msg, nil, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {
	s.log(zap.ErrorLevel, msg, nil, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	s.log(zap.DPanicLevel, msg, nil, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...interface{}) {
	s.log(zap.PanicLevel, msg, nil, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	s.log(zap.FatalLevel, msg, nil, keysAndValues)
}

// Debugln logs a message at [DebugLevel].
// Spaces are always added between arguments.
func (s *SugaredLogger) Debugln(args ...interface{}) {
	s.logln(zap.DebugLevel, args, nil)
}

// Infoln logs a message at [InfoLevel].
// Spaces are always added between arguments.
func (s *SugaredLogger) Infoln(args ...interface{}) {
	s.logln(zap.InfoLevel, args, nil)
}

// Warnln logs a message at [WarnLevel].
// Spaces are always added between arguments.
func (s *SugaredLogger) Warnln(args ...interface{}) {
	s.logln(zap.WarnLevel, args, nil)
}

// Errorln logs a message at [ErrorLevel].
// Spaces are always added between arguments.
func (s *SugaredLogger) Errorln(args ...interface{}) {
	s.logln(zap.ErrorLevel, args, nil)
}

// DPanicln logs a message at [DPanicLevel].
// In development, the logger then panics. (See [DPanicLevel] for details.)
// Spaces are always added between arguments.
func (s *SugaredLogger) DPanicln(args ...interface{}) {
	s.logln(zap.DPanicLevel, args, nil)
}

// Panicln logs a message at [PanicLevel] and panics.
// Spaces are always added between arguments.
func (s *SugaredLogger) Panicln(args ...interface{}) {
	s.logln(zap.PanicLevel, args, nil)
}

// Fatalln logs a message at [FatalLevel] and calls os.Exit.
// Spaces are always added between arguments.
func (s *SugaredLogger) Fatalln(args ...interface{}) {
	s.logln(zap.FatalLevel, args, nil)
}

// Sync flushes any buffered log entries.
func (s *SugaredLogger) Sync() error {
	return s.Logger.Sync()
}

func baseLogger(sugar *zap.SugaredLogger) *zap.Logger {
	p := unsafe.Pointer(sugar)
	offset := uintptr(0)
	return *(**zap.Logger)(unsafe.Pointer(uintptr(p) + offset))
}

// log message with Sprint, Sprintf, or neither.
func (s *SugaredLogger) log(lvl zapcore.Level, template string, fmtArgs []interface{}, context []interface{}) {
	// If logging at this level is completely disabled, skip the overhead of
	// string formatting.
	base := baseLogger(s.Logger)
	if lvl < zap.DPanicLevel && !base.Core().Enabled(lvl) {
		return
	}

	msg := getMessage(template, fmtArgs)
	if ce := base.Check(lvl, msg); ce != nil {
		ce.Write(s.sweetenFields(context)...)
	}
}

func (s *SugaredLogger) FormatMessage(msg string, context ...interface{}) string {
	fields := s.sweetenFields(context)

	buffer := bytes.NewBufferString(msg)
	if len(msg) > 0 && len(fields) > 0 {
		_, _ = buffer.WriteString(" ")
	}

	for i, field := range fields {
		if i > 0 {
			buffer.WriteByte(',')
		}
		buffer.WriteString(field.Key + ":")
		switch field.Type {
		case zapcore.BoolType:
			if field.Integer > 0 {
				buffer.WriteString("true")
			} else {
				buffer.WriteString("false")
			}
		case zapcore.Int64Type, zapcore.Int32Type, zapcore.Int16Type, zapcore.Int8Type,
			zapcore.Uint64Type, zapcore.Uint32Type, zapcore.Uint16Type, zapcore.Uint8Type,
			zapcore.Float64Type, zapcore.Float32Type,
			zapcore.TimeType, zapcore.UintptrType:
			buffer.WriteString(strconv.FormatInt(field.Integer, 10))
		case zapcore.StringType:
			buffer.WriteString(field.String)
		case zapcore.TimeFullType:
			buffer.WriteString(field.Interface.(time.Time).String())
		case zapcore.ReflectType:
			json, _ := jsoniter.ConfigDefault.MarshalToString(field.Interface)
			buffer.WriteString(json)
		case zapcore.StringerType:
			buffer.WriteString(field.Interface.(fmt.Stringer).String())
		case zapcore.ErrorType:
			buffer.WriteString(field.Interface.(error).Error())
		case zapcore.SkipType:
		default:
			buffer.WriteString(fmt.Sprintf("%v", field.Interface))
		}

	}

	return buffer.String()
}

// logln message with Sprintln
func (s *SugaredLogger) logln(lvl zapcore.Level, fmtArgs []interface{}, context []interface{}) {
	base := baseLogger(s.Logger)
	if lvl < zap.DPanicLevel && !base.Core().Enabled(lvl) {
		return
	}

	msg := getMessageln(fmtArgs)
	if ce := base.Check(lvl, msg); ce != nil {
		ce.Write(s.sweetenFields(context)...)
	}
}

// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

// getMessageln format with Sprintln.
func getMessageln(fmtArgs []interface{}) string {
	msg := fmt.Sprintln(fmtArgs...)
	return msg[:len(msg)-1]
}

func (s *SugaredLogger) sweetenFields(args []interface{}) []zap.Field {
	if len(args) == 0 {
		return nil
	}

	var (
		// Allocate enough space for the worst case; if users pass only structured
		// fields, we shouldn't penalize them with extra allocations.
		fields    = make([]zap.Field, 0, len(args))
		invalid   invalidPairs
		seenError bool
		base      = baseLogger(s.Logger)
	)

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// If it is an error, consume it and move on.
		if err, ok := args[i].(error); ok {
			if !seenError {
				seenError = true
				fields = append(fields, zap.Error(err))
			} else {
				base.Error(_multipleErrMsg, zap.Error(err))
			}
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			base.Error(_oddNumberErrMsg, zap.Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			switch val.(type) {
			case proto.Message:
				fields = append(fields, zap.Field{Key: keyStr, Type: zapcore.ReflectType, Interface: val})
			default:
				fields = append(fields, zap.Any(keyStr, val))
			}
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		base.Error(_nonStringKeyErrMsg, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	zap.Any("key", p.key).AddTo(enc)

	switch p.value.(type) {
	case proto.Message:
		zap.Field{Key: "value", Type: zapcore.ReflectType, Interface: p.value}.AddTo(enc)
	default:
		zap.Any("value", p.value).AddTo(enc)
	}

	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}

func initConsoleCore(encode string, level zap.AtomicLevel) zapcore.Core {
	encodeConfig := defaultConfig
	var formatEncoder zapcore.Encoder

	enc := strings.ToLower(encode)
	if enc == "json" {
		formatEncoder = zapcore.NewJSONEncoder(encodeConfig)
	} else {
		formatEncoder = zapcore.NewConsoleEncoder(encodeConfig)
	}
	consoleDebugging := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(formatEncoder, consoleDebugging, level)
}

func initFileCore(cfg FileSinkConfig, level zap.AtomicLevel) zapcore.Core {
	encodeConfig := defaultConfig

	var formatEncoder zapcore.Encoder
	enc := strings.ToLower(cfg.Encode)
	if enc == "json" {
		formatEncoder = zapcore.NewJSONEncoder(encodeConfig)
	} else {
		formatEncoder = zapcore.NewConsoleEncoder(encodeConfig)
	}

	f := handleFileName(cfg.Path)
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   f,
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge, // days
		Compress:   cfg.Compress,
	})

	return zapcore.NewCore(formatEncoder, w, level)
}

func handleFileName(filename string) string {
	filename = path.Clean(filename)
	parts := make([]string, 0)
	var ret string
	paths := strings.Split(filename, string(os.PathSeparator))
	for _, v := range paths {
		val := handleTemplateFileName(v)
		if len(val) > 0 {
			parts = append(parts, val)
		}
	}

	if path.IsAbs(filename) {
		ret = string(os.PathSeparator) + path.Join(parts...)
	} else {
		ret = path.Join(parts...)
	}
	return ret
}

func handleTemplateFileName(template string) string {
	// foo1{hostname}foo2{port}foo3
	lefts := make([]int, 0)
	rights := make([]int, 0)

	size := len(template)
	for i := 0; i < size; i++ {
		if template[i] == '{' {
			lefts = append(lefts, i)
		} else if template[i] == '}' {
			rights = append(rights, i)
		}
	}

	leftSize := len(lefts)
	rightSize := len(rights)
	var minSize int
	if leftSize < rightSize {
		minSize = leftSize
	} else {
		minSize = rightSize
	}

	ret := template
	for i := minSize - 1; i >= 0; i-- {
		variableName := ret[lefts[i]+1 : rights[i]]
		v := os.Getenv(variableName)
		ret = ret[:lefts[i]] + v + ret[rights[i]+1:]
	}
	return ret
}

func jsoniterReflectedEncoder(w io.Writer) zapcore.ReflectedEncoder {
	enc := jsoniter.NewEncoder(w)
	// For consistency with mojo custom JSON encoder.
	enc.SetEscapeHTML(false)
	return enc
}
