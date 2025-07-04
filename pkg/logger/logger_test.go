package logger

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newTestLogger(buf *bytes.Buffer) *AppLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(buf),
		zapcore.DebugLevel,
	)
	return &AppLogger{zapLogger: zap.New(core)}
}

func TestAppLogger_InfoWarnErrorDebug(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := newTestLogger(buf)

	logger.Info("info msg", "foo", 123)
	logger.Warn("warn msg", "bar", 456)
	logger.Error("error msg", "baz", 789)
	logger.Debug("debug msg", "qux", 999)

	out := buf.String()
	assert.Contains(t, out, "info msg")
	assert.Contains(t, out, "warn msg")
	assert.Contains(t, out, "error msg")
	assert.Contains(t, out, "debug msg")
	assert.Contains(t, out, "foo")
	assert.Contains(t, out, "bar")
	assert.Contains(t, out, "baz")
	assert.Contains(t, out, "qux")
}

func TestAppLogger_With(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := newTestLogger(buf)
	l2 := logger.With("user", "alice")
	l2.Info("with msg")
	out := buf.String()
	assert.Contains(t, out, "with msg")
	assert.Contains(t, out, "alice")
}

func TestAppLogger_Sync(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := newTestLogger(buf)
	assert.NoError(t, logger.Sync())
}

func TestToZapFields(t *testing.T) {
	fields := toZapFields("foo", 123, "bar", "baz")
	assert.Len(t, fields, 2)
	assert.Equal(t, "foo", fields[0].Key)
	assert.Equal(t, "bar", fields[1].Key)
}

func TestFieldHelpers(t *testing.T) {
	assert.Equal(t, zap.Float64("f", 1.23), Float64("f", 1.23))
	assert.Equal(t, zap.String("s", "abc"), String("s", "abc"))
	assert.Equal(t, zap.Int("i", 42), Int("i", 42))
	assert.Equal(t, zap.Error(errors.New("err")), Error(errors.New("err")))
	assert.Equal(t, zap.Duration("d", time.Second), Duration("d", time.Second))
}

func TestAppLogger_InfofWarnfErrorfDebugf(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := newTestLogger(buf)

	logger.Infof("info %s %d", "foo", 1)
	logger.Warnf("warn %s %d", "bar", 2)
	logger.Errorf("error %s %d", "baz", 3)
	logger.Debugf("debug %s %d", "qux", 4)

	out := buf.String()
	assert.Contains(t, out, "info foo 1")
	assert.Contains(t, out, "warn bar 2")
	assert.Contains(t, out, "error baz 3")
	assert.Contains(t, out, "debug qux 4")
}

// Fatal sẽ gọi os.Exit, không nên test trực tiếp trong unit test
