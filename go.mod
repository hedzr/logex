module github.com/hedzr/logex

go 1.17

//replace github.com/hedzr/log => ../10.log

require (
	github.com/hedzr/log v1.5.16
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220210151621-f4118a5b28e2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
	gopkg.in/hedzr/errors.v2 v2.1.9 // indirect
)
