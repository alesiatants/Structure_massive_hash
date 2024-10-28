package main

import (
	"errors"
)

// Структура уникальных значений
type Structure struct {
	valueMap map[interface{}]int // Хранит значение и его индекс в массиве
	values   []interface{}       // Массив для хранения значений в правильном порядке
}

// New создаёт новую структуру
func New() *Structure {
	return &Structure{
		valueMap: make(map[interface{}]int),
		values:   []interface{}{},
	}
}

// Add добавляет элемент в структуру
func (structure *Structure) Add(value interface{}) {
	if _, exists := structure.valueMap[value]; !exists {
		structure.valueMap[value] = len(structure.values)  // Делаем индекс равным текущей длине массива
		structure.values = append(structure.values, value) // Добавляем значение в массив
	}
}

// Remove удаляет элемент из структуры
func (structure *Structure) Remove(value interface{}) error {
	index, exists := structure.valueMap[value]
	if !exists {
		return errors.New("Искомое значение для удаления не найдено в структуре!")
	}

	// Меняем текущий элемент с последним, затем удаляем последний
	lastIndex := len(structure.values) - 1
	structure.values[index], structure.values[lastIndex] = structure.values[lastIndex], structure.values[index]
	structure.values = structure.values[:lastIndex] // Убираем последний элемент

	// Обновляем индексы в мапе
	delete(structure.valueMap, value)
	if index != lastIndex {
		// Если мы обменяли элементы, обновляем индекс последнего элемента в мапе
		structure.valueMap[structure.values[index]] = index
	}
	return nil
}

// Get возвращает элемент по индексу
func (structure *Structure) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(structure.values) {
		return nil, errors.New("Индекс находится за пределами доступных!")
	}
	return structure.values[index], nil
}

// GetByValue возвращает элемент по значению
func (structure *Structure) GetByValue(value interface{}) (interface{}, int, error) {
	index, exists := structure.valueMap[value]
	if !exists {
		return nil, 0, errors.New("Элемент с запрашиваемым значением не найден!")
	}
	return structure.values[index], index, nil
}

// Update обновляет элемент в структуре
func (structure *Structure) Update(oldValue, newValue interface{}) error {
	index, exists := structure.valueMap[oldValue]
	if !exists {
		return errors.New("Элемент для обнавления с запрашиваемым значением не найден!")
	}

	structure.valueMap[newValue] = index // Добавляем новый элемент в мапу
	delete(structure.valueMap, oldValue) // Удаляем старый элемент из мапы
	structure.values[index] = newValue   // Обновляем значение в массиве

	return nil
}

// Size возвращает количество элементов в структуре данных
func (structure *Structure) Size() int {
	return len(structure.values)
}
