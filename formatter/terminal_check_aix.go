//go:build !appengine && !js && !windows && aix
// +build !appengine,!js,!windows,aix

/*
 * Copyright © 2019 Hedzr Yeh.
 */

package formatter

import "io"

func checkIfTerminal(w io.Writer) bool {
	return false
}
