module github.com/hedzr/logex

go 1.13

// replace github.com/hedzr/log => ../log

require (
	github.com/hedzr/log v0.1.7
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/sirupsen/logrus v1.6.0
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	golang.org/x/sys v0.0.0-20191005200804-aed5e4c7ecf9 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
