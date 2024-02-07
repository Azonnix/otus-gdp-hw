package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromStat, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if offset > fromStat.Size() {
		return errors.New("invalid offset")
	}

	if limit == 0 || limit > fromStat.Size() {
		limit = fromStat.Size()
	}

	fromFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	_, err = fromFile.Seek(offset, io.SeekCurrent)
	if err != nil {
		return err
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	reader := io.LimitReader(fromFile, limit)
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(reader)
	io.Copy(toFile, barReader)
	bar.Finish()

	return nil
}
