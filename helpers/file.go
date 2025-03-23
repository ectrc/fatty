package helpers

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"sync"
)

type HelperFile struct {
	mutex *sync.Mutex
	file *os.File
}

func File(path_ string) (*HelperFile, error) {
	file, err := os.OpenFile(path_, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if os.IsNotExist(err) {
		real := path.Dir(path_)

		if err := os.MkdirAll(real, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory: %s %s", real, err)
		}

		fmt.Printf("created path: %s\n", real)

		file, err = os.Create(path_)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %s %s", path_, err)
		}

		fmt.Printf("created file: %s\n", path_)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s %s", path_, err)
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