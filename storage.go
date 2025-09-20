package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	fileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName: fileName}
}

func (s *Storage[T]) save(data T) error {

	fileData, err := json.MarshalIndent(data, "", "       ")

	if err != nil {
		return err
	}

	if err := os.WriteFile(s.fileName, fileData, 0644); err != nil {
		return err
	}

	return nil
}

func (s *Storage[T]) load(data *T) error {

	fileData, err := os.ReadFile(s.fileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}
