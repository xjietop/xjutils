package xjutils

import "testing"

func TestLogOutput(t *testing.T) {
	LogDebug("这是一条调试信息")
	LogInfo("这是一条普通信息")
	LogWarn("这是一条警告信息")
	LogError("这是一条错误信息")
	LogPrefix("这是一条自定义前辍信息","[XYZ]")
}
