package helpers

import (
	"bufio"
	"os"
	"sync"
)

type HelperFile struct {
	mutex *sync.Mutex
	file *os.File
}

func File(path string) (*HelperFile, error) {
	file, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return &HelperFile{
		mutex: &sync.Mutex{},
		file: file,
	}, nil
}

func (h *HelperFile) Write(data []byte) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	_, err := h.file.Write(data)
	return err
}

func (h *HelperFile) Close() error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	return h.file.Close()
}

func (h *HelperFile) ReadAllLines() ([]string, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	lines := make([]string, 0)
	scanner := bufio.NewReader(h.file)

	for {
		line, _, err := scanner.ReadLine()
		if err != nil {
			break
		}

		lines = append(lines, string(line))
	}

	return lines, nil
}