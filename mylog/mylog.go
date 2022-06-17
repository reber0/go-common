/*
 * @Author: reber
 * @Mail: reber0ask@qq.com
 * @Date: 2022-01-05 17:49:03
 * @LastEditTime: 2022-06-17 23:54:11
 */
package mylog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MyLog struct {
	Log          *zap.Logger
	InfoFile     string
	ErrorFile    string
	IsToConsole  bool
	IsToFile     bool
	IsShowCaller bool
}

// 初始化 MyLog
func New() *MyLog {
	return &MyLog{
		Log:          &zap.Logger{},
		InfoFile:     "./logs/info.log",
		ErrorFile:    "./logs/error.log",
		IsToConsole:  true,
		IsToFile:     false,
		IsShowCaller: false,
	}
}

// 日志输出到终端
func (mylog *MyLog) ToConsole(value bool) *MyLog {
	mylog.IsToConsole = value
	return mylog
}

// 日志输出到文件
func (mylog *MyLog) ToFile(value bool) *MyLog {
	mylog.IsToFile = value
	return mylog
}

// 设置 Info 日志路径
func (mylog *MyLog) SetInfoFile(logfile string) *MyLog {
	mylog.InfoFile = logfile
	return mylog
}

// 设置 Error 日志路径
func (mylog *MyLog) SetErrorFile(logfile string) *MyLog {
	mylog.ErrorFile = logfile
	return mylog
}

// 是否显示 Caller
func (mylog *MyLog) ShowCaller(value bool) *MyLog {
	mylog.IsShowCaller = value
	return mylog
}

// 生成 *zap.Logger
func (mylog *MyLog) Logger() *zap.Logger {

	var coreArr []zapcore.Core
	if mylog.IsToConsole && mylog.IsToFile {
		consoleCore := setConsole()
		infoFileCore, errorFileCore := setFile(mylog.InfoFile, mylog.ErrorFile)
		coreArr = []zapcore.Core{consoleCore, infoFileCore, errorFileCore}
	} else {
		if mylog.IsToConsole {
			consoleCore := setConsole()
			coreArr = []zapcore.Core{consoleCore}
		}
		if mylog.IsToFile {
			infoFileCore, errorFileCore := setFile(mylog.InfoFile, mylog.ErrorFile)
			coreArr = []zapcore.Core{infoFileCore, errorFileCore}
		}
	}

	var log *zap.Logger
	if mylog.IsShowCaller {
		log = zap.New(
			zapcore.NewTee(coreArr...),
			zap.AddCaller(), // zap.AddCaller() 设为显示文件名和行号
		)
	} else {
		log = zap.New(
			zapcore.NewTee(coreArr...),
		)
	}

	return log
}
