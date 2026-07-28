package main 

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

type Storage[T any] struct {
	fileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	
	dir := "todos"
	

	root, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error saving to home dir: %w", err)
	}
  
	todosPath := filepath.Join(root, dir)
  

	err = os.MkdirAll(todosPath, 0755)
	if err != nil {
		return fmt.Errorf("Error making dir: %w", err)
	}
	
	fullPath := filepath.Join(root, dir, s.fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("Err creating file: %w", err) 
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)

}

func (s *Storage[T]) Load() (T, error) {

	var data T

	root, err := os.UserHomeDir()
	if err != nil {
		return data, fmt.Errorf("error finding home dir: %w", err)
	}

	fullPath := filepath.Join(root, "todos", s.fileName)

	
	file, err := os.Open(fullPath)
	if err != nil {
		return data, fmt.Errorf("error opening file: %w", err) 
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	return data, err

}
