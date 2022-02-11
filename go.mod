module github.com/hedzr/logex

go 1.17

//replace github.com/hedzr/log => ../10.log

require github.com/hedzr/log v1.5.11

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220210151621-f4118a5b28e2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
