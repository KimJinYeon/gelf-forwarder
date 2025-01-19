package internal

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
)

// Decompress tries to decompress data using gzip.
func Decompress(data []byte) ([]byte, error) {

	if r, err := gzip.NewReader(bytes.NewReader(data)); err == nil {
		defer r.Close()
		decompressed, err := io.ReadAll(r)
		if err == nil {
			return decompressed, nil
		}
	}

	return nil, errors.New("failed to decompress data")
}
