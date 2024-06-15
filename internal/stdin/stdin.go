package stdin

import (
	"bufio"
	"errors"
	"os"
)

const (
	bufferCap  = 64 * 1024
	bufferSize = 1024 * 1024
)

var ErrNoStdin = errors.New("unable to read from stdin")

func Stream(callback func(string)) error {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return ErrNoStdin
	}

	if (fi.Mode()&os.ModeNamedPipe == 0) && fi.Size() <= 0 {
		return ErrNoStdin
	}

	scan := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, bufferCap)
	scan.Buffer(buf, bufferSize)

	for scan.Scan() {
		callback(scan.Text())
	}

	if err := scan.Err(); err != nil {
		return ErrNoStdin
	}

	return nil
}
