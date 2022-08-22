package test_log_2

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestInitLog(t *testing.T) {
	InitLog("tiloader.log", "tiloader_err.log")
	_, err := os.Stat("tiloader.log")
	if err != nil {
		t.Fatal("日志创建失败")
	} else {
		t.Log("tiloader.log 创建成功")
	}
}

func TestWriteLog(t *testing.T) {
	var log = "WriteLog执行成功"
	sugaredLogger.Infof(log)
	data, err := ioutil.ReadFile("tiloader.log")
	if err != nil {
		t.Fatal("读取全量日志失败")
	} else {
		t.Log("读取全量日志成功")
		fmt.Println(string(data))
	}
}

// 错误日志接口
func TestError(t *testing.T) {
	var logstr = "执行失败 xxxx"

	Error(logstr)
	data, err := ioutil.ReadFile("tiloader_err.log")
	if err != nil {
		t.Fatal("读取错误日志失败")
	} else {
		t.Log("读取错误日志成功")
		fmt.Println(string(data))
	}
}
