package ensurecp

import (
	"os"
	"crypto/sha512"
	"io"
	"encoding/hex"
	"fmt"
)

type CopyPath string

func (c CopyPath) Stat() error {
	_, err := os.Stat(string(c))
	if err != nil {
		return err
	}

	return nil
}

func (c CopyPath) GetHash() (string, error) {
	fh, err := os.Open(string(c))
	defer fh.Close()
	if err != nil {
		return "", err
	}

	hasher := sha512.New()

	if _, err = io.Copy(hasher, fh); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (c CopyPath) Write(src io.Reader, mode... os.FileMode) error {
	defmode := os.FileMode(0744)
	if len(mode) > 0 {
		defmode = mode[0]
	}

	fh, err := os.OpenFile(string(c), os.O_CREATE | os.O_RDWR | os.O_TRUNC, defmode)
	defer fh.Close()

	if err != nil {
		return err
	}

	_, err = io.Copy(fh, src)
	if err != nil {
		return err
	}

	return nil
}

func compareHash(src, dst CopyPath) bool {
	srcHash, err := src.GetHash()
	if err != nil {
		return false
	}

	dstHash, err := dst.GetHash()
	if err != nil {
		return false
	}

	state := srcHash == dstHash

	if state && localConfig.EnableLogging {
		logEntry := CopyLog{string(src), string(dst), srcHash,}
		currentLog = append(currentLog, logEntry)
	}

	return state
}

func (c CopyPath) CopyTo(dest string, verify... bool) (CopyPath, error) {
	cpDest := CopyPath(dest)

	src, err := os.Open(string(c))
	defer src.Close()
	if err != nil {
		return cpDest, err
	}

	stat, err := src.Stat()
	if err != nil {
		return cpDest, err
	}

	err = cpDest.Write(src, stat.Mode())
	if err != nil {
		return cpDest, err
	}

	if len(verify) > 0 && verify[0] && !compareHash(c, cpDest) {
		return cpDest, fmt.Errorf("File hashes do not match!")
	}

	return cpDest, err
}
