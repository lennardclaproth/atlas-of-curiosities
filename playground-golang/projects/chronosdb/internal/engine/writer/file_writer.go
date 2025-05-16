package writer

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/lennardclaproth/chronosdb/internal/engine/page"
)

type FileWriter struct {
	file       *os.File
	mutex      sync.Mutex
	currentPos int64
}

func NewFileWriter(path string) (*FileWriter, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	pos, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	return &FileWriter{
		file:       file,
		currentPos: pos,
	}, nil
}

func (w *FileWriter) WritePage(p *page.Page) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	pageBytes, err := p.Encode()
	if err != nil {
		return err
	}

	n, err := w.file.Write(pageBytes)
	if err != nil {
		return err
	}

	if n != len(pageBytes) {
		return os.ErrInvalid
	}

	// Optionally, flush to disk
	if err := w.file.Sync(); err != nil {
		return err
	}

	w.currentPos += int64(n)

	return nil
}

func (w *FileWriter) Close() error {
	return w.file.Close()
}
