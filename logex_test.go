package logex

import (
	"os"
	"testing"

	"github.com/hedzr/log/dir"
)

func TestGetLevel(t *testing.T) {
	t.Logf("level = %v", GetLevel())
	t.Logf("debug = %v", GetDebugMode())
	t.Logf("trace = %v", GetTraceMode())
}

func TestReadWriteFile(t *testing.T) {
	d, err := dir.TempDir("", "")
	if err != nil {
		t.Errorf("TempDir: %v", err)
		return
	}

	defer func(d string) {
		err = dir.RemoveDirRecursive(d)
		if err != nil {
			t.Errorf("dir.RemoveDirRecursive: %v", err)
		}
	}(d)
	t.Logf("temp dir got: %v", d)

	var f *os.File
	f, err = dir.TempFile(d, "")
	if err != nil {
		t.Errorf("TempFile: %v", err)
		return
	}

	fn := f.Name()
	t.Logf("temp file got: %v", d)
	err = f.Close()
	if err != nil {
		t.Errorf("close TempFile: %v", err)
		return
	}

	defer func(name string) {
		err = os.Remove(name)
		if err != nil {
			t.Errorf("os.Remove: %v", err)
		}
	}(fn)

	err = dir.WriteFile(fn, []byte(`okok`), 0o600)
	if err != nil {
		t.Errorf("dir.WriteFile: %v", err)
		return
	}

	var b []byte
	b, err = dir.ReadFile(fn)
	if err != nil {
		t.Errorf("dir.ReadFile: %v", err)
		return
	}

	if string(b) != `okok` {
		t.Fatalf("read file content not ok")
	}
}
