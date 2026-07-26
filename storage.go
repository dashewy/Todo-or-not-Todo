package main 

import (
	"os"
	"encoding/json"
)

type Storage[T any] struct {
	fileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	
	file, err := os.Create(s.fileName)
	if err != nil {
		return err 
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)

}

func (s *Storage[T]) Load() (T, error) {
	
	var data T 
	file, err := os.Open(s.fileName)
	if err != nil {
		return data, err 
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	return data, err

}
