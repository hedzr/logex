module github.com/hedzr/logex

go 1.13

replace github.com/hedzr/log => ../10.log

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/hedzr/log v1.3.25
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.20.0
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
