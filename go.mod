module github.com/hedzr/logex

go 1.17

//replace gopkg.in/hedzr/errors.v3 => ../05.errors

//replace github.com/hedzr/log => ../10.log

require (
	github.com/hedzr/log v1.5.53
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/rs/zerolog v1.26.1
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/stretchr/testify v1.7.5 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467 // indirect
	gopkg.in/hedzr/errors.v3 v3.0.21 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
