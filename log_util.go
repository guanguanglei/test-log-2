package test_log_2

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogConfig struct {
	Filename string
}

//获取参数模版定义的日志路径以及日志文件名字
var (
	sugaredLogger *zap.SugaredLogger
	infoLog       FileLogConfig
	errLog        FileLogConfig
)

// check 是否是一个目录
func initFileLog(cfg *FileLogConfig) {
	if st, err := os.Stat(cfg.Filename); err == nil {
		if st.IsDir() {
			// fmt.Printf("is dir,请定义日志文件名字")
			panic("参数定义是一个目录,请定义日志文件名字")
		}
	}

}

// 初始化 zap logger
func InitLog(logFile string, errFile string) {
	// 获取全量日志文件和错误日志文件

	infoLog.Filename = logFile
	initFileLog(&infoLog)
	errLog.Filename = errFile
	initFileLog(&errLog)

	encoder := getEncoder()

	// 定义 zap WriteSyncer
	//  infoLog.Filename 记录全量日志
	logF, _ := os.OpenFile(infoLog.Filename, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, os.ModeAppend|0755)
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.InfoLevel)
	// errLog.Filename 记录ERROR级别的日志
	//errF, _ := os.OpenFile(errLog.Filename, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, os.ModeAppend|0755)
	//	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zapcore.ErrorLevel)
	// 使用NewTee将os.Stdout、c1和c2合并到core
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		c1,
	//	c2,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugaredLogger = logger.Sugar()

}

// 定义 zap encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 全量日志接口
func WriteLog(log string) {
	sugaredLogger.Infof(log)
}

// 错误日志接口
// func Error(log string) {
// 	fmt.Printf(log)
// }

func Error(logstr string) {
	f, err := os.OpenFile(errLog.Filename, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, os.ModeAppend|0755)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetFlags(log.Lmsgprefix)
	log.SetOutput(f)
	log.Printf(logstr)
}
