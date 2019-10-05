# logex


<!-- ![Build Status](https://travis-ci.org/hedzr/logex.svg?branch=master)](https://travis-ci.org/hedzr/logex) -->
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/hedzr/logex.svg?label=release)](https://github.com/hedzr/logex/releases)
[![Sourcegraph](https://sourcegraph.com/github.com/hedzr/logex/-/badge.svg)](https://sourcegraph.com/github.com/hedzr/logex?badge)

an enhanced for logrus. `logex` append the context call info to the log.



![image-20190706194022859](assets/image-20190706194022859.png)



## Usage

```go
import "github.com/hedzr/logex"

func init(){
	logex.Enable()
}
```

Or:

```go
import "github.com/hedzr/logex"

func init(){
	logex.EnableWith(logrus.DebugLevel)
}
```



### import `logex` from gopkg.in:

```go
import "gopkg.in/hedzr/logex.v1"
```


## make `logrus` works in `go test`

The codes is copied from:

<https://github.com/sirupsen/logrus/issues/834>

And in a test function, you could code now:

```go
   func TestFoo(t *testing.T) {
     defer logex.CaptureLog(t).Release()
     // …
   }
```

## LICENSE

MIT
