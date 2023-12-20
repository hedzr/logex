module github.com/hedzr/logex

go 1.17

//replace gopkg.in/hedzr/errors.v3 => ../05.errors

//replace github.com/hedzr/log => ../10.log

require (
	github.com/hedzr/log v1.6.25
	github.com/konsorten/go-windows-terminal-sequences v1.0.3
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.17.0
	gopkg.in/hedzr/errors.v3 v3.3.0
)

require (
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
)
