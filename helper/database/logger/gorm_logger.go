package logger

import (
	"context"
	"errors"
	"fmt"
	"github/tronglv_authen_author/helper/logify"

	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	glr "gorm.io/gorm/logger"
)

type logger struct {
	glr.Writer
	glr.Config
	printer                             logify.Logger
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New(writer glr.Writer, config glr.Config) glr.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = glr.Green + "%s\n" + glr.Reset + glr.Green + "[info] " + glr.Reset
		warnStr = glr.BlueBold + "%s\n" + glr.Reset + glr.Magenta + "[warn] " + glr.Reset
		errStr = glr.Magenta + "%s\n" + glr.Reset + glr.Red + "[error] " + glr.Reset
		traceStr = glr.Green + "%s\n" + glr.Reset + glr.Yellow + "[%.3fms] " + glr.BlueBold + "[rows:%v]" + glr.Reset + " %s"
		traceWarnStr = glr.Green + "%s " + glr.Yellow + "%s\n" + glr.Reset + glr.RedBold + "[%.3fms] " + glr.Yellow + "[rows:%v]" + glr.Magenta + " %s" + glr.Reset
		traceErrStr = glr.RedBold + "%s " + glr.MagentaBold + "%s\n" + glr.Reset + glr.Yellow + "[%.3fms] " + glr.BlueBold + "[rows:%v]" + glr.Reset + " %s"
	}

	return &logger{
		Writer:       writer,
		Config:       config,
		printer:      logify.New(),
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

func (l *logger) LogMode(level glr.LogLevel) glr.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *logger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glr.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{FileWithLineNum()}, data...)...)
	}
}

func (l *logger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glr.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{FileWithLineNum()}, data...)...)
	}
}

func (l *logger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= glr.Error {
		l.Printf(l.errStr+msg, append([]interface{}{FileWithLineNum()}, data...)...)
	}
}

func (l *logger) fields(elapsed time.Duration, fc func() (string, int64)) []logc.LogField {
	sql, rows := fc()
	fields := []logc.LogField{
		logc.Field("sql", sql),
		logc.Field("duration", float64(elapsed.Nanoseconds())/1e6),
		logc.Field("file", FileWithLineNum()),
	}
	if rows != -1 {
		fields = append(fields, logc.Field("rows", rows))
	}
	return fields
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= glr.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= glr.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		l.printer.WithContext(ctx).Errorw(err.Error(), l.fields(elapsed, fc)...)

	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= glr.Warn:
		l.printer.WithContext(ctx).Infow(fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold), l.fields(elapsed, fc)...)

	case l.LogLevel == glr.Info:
		l.printer.WithContext(ctx).Infow("SQL", l.fields(elapsed, fc)...)
	}
}

func (l *logger) ParamsFilter(_ context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if l.Config.ParameterizedQueries {
		return sql, nil
	}
	return sql, params
}

type traceRecorder struct {
	glr.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func (l *traceRecorder) New() *traceRecorder {
	return &traceRecorder{Interface: l.Interface, BeginAt: time.Now()}
}

func (l *traceRecorder) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.BeginAt = begin
	l.SQL, l.RowsAffected = fc()
	l.Err = err
}

func FileWithLineNum() string {
	for i := 3; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.Contains(file, "gorm.io") || strings.HasSuffix(file, "_test.go")) {
			files := strings.Split(file, "/")
			if len(files) >= 2 {
				return strings.Join(files[len(files)-2:], "/") + ":" + strconv.FormatInt(int64(line), 10)
			}
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}
