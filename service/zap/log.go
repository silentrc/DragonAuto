package zap

import (
	"dragonAuto/config"
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var lg *zap.Logger

type zaplog struct {
	system string
}

// InitLogger 初始化Logger
func Init() {
	zapCores := getEnablesLogs()
	core := zapcore.NewTee(zapCores...)
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEnablesLogs() (result []zapcore.Core) {
	encoder := getEncoder()
	infoPath := config.Instance.Log.Path + config.Instance.Log.InfoPath

	if config.Instance.Log.EnableInfoLog {
		infoWriter := getWriter(infoPath)
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			if lvl < zapcore.WarnLevel && lvl >= config.Instance.Log.Level {
				return true
			}
			return false
		})
		infoCore := zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel)
		result = append(result, infoCore)
	}

	if config.Instance.Log.EnableWarnLog {
		errorPath := config.Instance.Log.Path + config.Instance.Log.ErrorPath
		warnWriter := getWriter(errorPath)
		warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			if lvl >= zapcore.WarnLevel && lvl >= config.Instance.Log.Level {
				return true
			}
			return false
		})
		warnCore := zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel)
		result = append(result, warnCore)
	}

	if config.Instance.Log.EnableConsoleLog {
		consoleConfig := zap.NewProductionEncoderConfig()
		consoleConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
		consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
		result = append(result, consoleCore)
	}
	return result
}

func getEncoder() zapcore.Encoder {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.CallerKey = "file"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendInt64(int64(d) / 1000000)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志

	//单个日志最大占用空间
	logMaxSize := int64(config.Instance.Log.MaxSize) * 1000 * 1000
	//日志最长保留日期
	logMaxAge := time.Duration(config.Instance.Log.MaxAge) * time.Hour * 24
	//日志多保存数量
	//logMaxCount := uint(config.Instance.Log.MaxBackups)

	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d %H:%I:%S", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(logMaxAge),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithRotationSize(logMaxSize),
		//rotatelogs.WithRotationCount(logMaxCount),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
