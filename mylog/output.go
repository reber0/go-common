/*
 * @Author: reber
 * @Mail: reber0ask@qq.com
 * @Date: 2022-06-17 23:17:52
 * @LastEditTime: 2022-06-17 23:18:01
 */
package mylog

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setConsole() zapcore.Core {
	// 配置终端日志显示格式，为普通文本格式
	encoderConsole := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		MessageKey:   "msg",
		EncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeLevel:  zapcore.CapitalColorLevelEncoder, // 按级别显示不同颜色
		EncodeCaller: zapcore.ShortCallerEncoder,       // 显示短文件路径
	})

	// 配置 Console 中日志格式
	consoleWriteSyncer := zapcore.AddSync(
		os.Stdout,
	)
	// zapcore.NewCore 第一个参数为日志格式，第二个参数为输出到哪里，第三个参数为日志级别
	consoleCore := zapcore.NewCore(encoderConsole, zapcore.NewMultiWriteSyncer(consoleWriteSyncer), zap.DebugLevel)

	return consoleCore
}

func setFile(info_file, error_file string) (zapcore.Core, zapcore.Core) {
	// 配置日志文件中日志的格式，为 json 格式
	encoderFile := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		MessageKey:   "msg",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.FullCallerEncoder, // 显示完整文件路径
	})

	// 设置日志级别，debug/info/warn/error/dpanic/panic/fatal 对应 -1/0/1/2/3/4/5
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // 低于 error 级别的记录
		return lev < zap.ErrorLevel
	})
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // 大于等于 error 级别的记录
		return lev >= zap.ErrorLevel
	})

	// 配置 debug/info
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   info_file, // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    10,        // 文件大小限制，单位 MB
		MaxBackups: 50,        // 最大保留日志文件数量
		MaxAge:     30,        // 日志文件保留天数
		Compress:   false,     // 是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoderFile, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer), lowPriority)

	// error 文件 writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   error_file, // 日志文件存放目录
		MaxSize:    10,         // 文件大小限制，单位 MB
		MaxBackups: 50,         // 最大保留日志文件数量
		MaxAge:     30,         // 日志文件保留天数
		Compress:   false,      // 是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoderFile, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer), highPriority)

	return infoFileCore, errorFileCore
}
