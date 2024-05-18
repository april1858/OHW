package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.

	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = fi.Size()
	}
	if limit > fi.Size() {
		limit = fi.Size() - offset
	}
	if offset > fi.Size() {
		return ErrOffsetExceedsFileSize
	}

	reader := io.NewSectionReader(file, offset, limit)

	fileTo, _ := os.Create(toPath)
	defer fileTo.Close()

	bar := pb.New(int(limit)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)

	bar.ShowSpeed = true
	bar.Start()

	pxReader := bar.NewProxyReader(reader)

	io.Copy(fileTo, pxReader)
	bar.Finish()

	return nil
}
