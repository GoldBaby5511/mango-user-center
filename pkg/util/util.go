package util

import (
	"io"
	"math/big"
	"net"

	"gopkg.in/natefinch/lumberjack.v2"
)

func Ip2Int(ip string) int {
	i := big.NewInt(0).SetBytes(net.ParseIP(ip).To4()).Int64()
	return int(i)
}

func Int2Ip(ip int) string {
	return net.IPv4(byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip)).String()
}

func NewWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    50,
		MaxAge:     1,
		MaxBackups: 2,
		LocalTime:  true,
		Compress:   false, // 不压缩 （自己清理）
	}
}
