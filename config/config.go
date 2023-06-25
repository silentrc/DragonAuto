package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"log"
)

var Instance = new(ConfigValue)

// 初始化mysql
func init() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	Instance.Cfg = &viper.Viper{}

	var zapLevel zapcore.Level
	logLevel := viper.GetString("log.level")
	switch logLevel {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "panic":
		zapLevel = zapcore.PanicLevel
	case "dpanic":
		zapLevel = zapcore.DPanicLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	}
	//log 配置
	Instance.Log = logConfig{
		Level:            zapLevel,
		Path:             viper.GetString("log.path"),
		EnableConsoleLog: viper.GetBool("log.enableConsoleLog"),
		EnableInfoLog:    viper.GetBool("log.enableInfoLog"),
		EnableWarnLog:    viper.GetBool("log.enableWarnLog"),
		InfoPath:         viper.GetString("log.infoPath"),
		ErrorPath:        viper.GetString("log.errorPath"),
		MaxSize:          viper.GetInt("log.maxsize"),
		MaxAge:           viper.GetInt("log.max_age"),
		MaxBackups:       viper.GetInt("log.max_backups"),
	}

	Instance.DragonAuto.Enable = viper.GetBool("dragonAuto.enable")
	Instance.DragonAuto.Mode = viper.GetInt("dragonAuto.mode")
	if Instance.DragonAuto.Mode == 0 {
		Instance.DragonAuto.Mode = 1
	}
	Instance.DragonAuto.Account = viper.GetString("dragonAuto.account")
	Instance.DragonAuto.Pwd = viper.GetString("dragonAuto.pwd")
	Instance.DragonAuto.ReqToken = viper.GetString("dragonAuto.token")
	Instance.DragonAuto.CollectTime = viper.GetInt("dragonAuto.collectTime")
	Instance.DragonAuto.IncubatorId1 = viper.GetString("dragonAuto.incubatorId1")
	Instance.DragonAuto.IncubatorId2 = viper.GetString("dragonAuto.incubatorId2")
	Instance.DragonAuto.IncubatorId3 = viper.GetString("dragonAuto.incubatorId3")

}

type ConfigValue struct {
	Log        logConfig
	Cfg        *viper.Viper
	DragonAuto dragonAuto
}

type dragonAuto struct {
	Enable       bool
	Mode         int
	ReqToken     string
	Token        string
	Account      string
	Pwd          string
	CollectTime  int
	IncubatorId1 string
	IncubatorId2 string
	IncubatorId3 string
}

type logConfig struct {
	EnableConsoleLog bool          `json:"enable_console"`
	EnableInfoLog    bool          `json:"enable_info_log"`
	EnableWarnLog    bool          `json:"enable_warn_log"`
	Level            zapcore.Level `json:"level"`
	Path             string        `json:"path"`
	InfoPath         string        `json:"info_path"`
	ErrorPath        string        `json:"error_path"`
	MaxSize          int           `json:"maxsize"`
	MaxAge           int           `json:"max_age"`
	MaxBackups       int           `json:"max_backups"`
}
