package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.

	count := 3000

	// create and start new bar
	bar := pb.StartNew(count)

	var sizeBuf int64

	file, err := os.Open(fromPath)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := file.Stat()
	if err != nil {
		panic(err)
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

	s := io.NewSectionReader(file, offset, limit)

	if offset+limit > fi.Size() {
		sizeBuf = fi.Size() - offset
	} else {
		sizeBuf = limit
	}

	buf := make([]byte, sizeBuf)

	if _, err := s.Read(buf); err != nil {
		if errors.Is(err, io.EOF) {
			log.Fatal("err - ", err)
		}
	}
	file.Close()

	fileTo, _ := os.Create(toPath)

	if _, err := io.Copy(fileTo, strings.NewReader(string(buf))); err != nil {
		log.Fatal(err)
	}
	fileTo.Close()

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}

	// finish bar
	bar.Finish()

	return nil
}
