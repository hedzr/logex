module github.com/hedzr/logex

go 1.13

//replace github.com/hedzr/log => ../10.log

require (
	github.com/hedzr/log v0.3.19
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20190510104115-cbcb75029529
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
