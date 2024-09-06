package contracts

import "os"

type Logger interface {
	Printf(format string, args ...interface{})
}

type (
	FSWriter interface {
		MkdirAll
		WriteFile
	}
	MkdirAll interface {
		MkdirAll(path string, perm os.FileMode) error
	}
	WriteFile interface {
		WriteFile(filename string, data []byte, perm os.FileMode) error
	}
)
