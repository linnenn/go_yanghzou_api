package library

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_yangzhou/config"
)
//记录日志
//记录日志的级别，日志的行数
//记录日志和打印到控制台，两个都可以切换
//日志大小可以做切片，分割
//这里使用两个仓库，一个uber的记录日志的仓库，一个是文件大小分割的仓库
//go get -u go.uber.org/zap
//go get -u github.com/natefinch/lumberjack

var Logger *zap.SugaredLogger
//初始化
//初始化存储文件配置，使用第三方工具做好文件大小分割
//初始化时间格式，错误级别大写变动
//初始化记录调用过程记录
//配置
func init() {
	Init()
	defer Logger.Sync()
}

func Init()  {
	file := getLogFile()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder,file,zapcore.DebugLevel)
	logger := zap.New(core,zap.AddCaller())
	Logger = logger.Sugar()
}

func getEncoder() zapcore.Encoder  {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}
//日志分割，存储
func getLogFile() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: config.InitConfig.Log.Filename,
		MaxSize: config.InitConfig.Log.MaxSize,
		MaxBackups: config.InitConfig.Log.MaxBackups,
		MaxAge: config.InitConfig.Log.MaxAge,
		Compress: config.InitConfig.Log.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}