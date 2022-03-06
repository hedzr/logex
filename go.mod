module github.com/hedzr/logex

go 1.17

//replace gopkg.in/hedzr/errors.v3 => ../05.errors

//replace github.com/hedzr/log => ../10.log

require (
	github.com/hedzr/log v1.5.29
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	gopkg.in/hedzr/errors.v3 v3.0.10 // indirect
)
