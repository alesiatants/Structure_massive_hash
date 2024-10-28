package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Add = iota + 1
	Update
	Delete
	SearchByValue
	SearchByPosition
	Exit
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	file, errOpen := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	structure := New()
	if errOpen != nil {
		log.Panic("Невозможно создать файл с логами ", errOpen)
	}
	for {

		fmt.Println(`

    Выберите действие

		1. Добавить элемент
		2. Обновить  элемент по значению
		3. Удалить по значению
		4. Получить элемент по значению
		5. Получить элемент по позиции
		6. Выйти

    `)

		fmt.Print(">>> ")

		choiceStr, errRead := reader.ReadString('\n')

		if errRead != nil {
			mw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(mw)
			fmt.Println("Ошибка при чтении ввода:", errRead)
			continue
		}

		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)

		if err != nil {
			mw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(mw)
			log.Println("Неверный ввод! Введите числовой пункт меню.")
			defer file.Close()
			continue
		}

		switch choice {

		case Add:

			fmt.Print("Введите значение для добавления: \n>>> ")
			newElement, errRead := reader.ReadString('\n')

			if errRead != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при чтении ввода:", errRead)
				continue
			}

			structure.Add(strings.TrimSpace(newElement))
			fmt.Printf("Размерность: %d", structure.Size())

		case Update:

			fmt.Print("Введите через пробел старое и новое значение: \n>>> ")
			elements, errRead := reader.ReadString('\n')

			if errRead != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при чтении ввода:", errRead)
				continue // Продолжаем цикл, если возникла ошибка
			}
			el := strings.Fields(elements)

			if len(el) != 2 {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Было введено не два значения!")
				continue
			}

			old := strings.TrimSpace(el[0])
			new := strings.TrimSpace(el[1])
			errUpdate := structure.Update(old, new)

			if errUpdate != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при обнавлении значения:", errUpdate)
				continue
			} else {
				fmt.Println("Успешно обнавлено!")
			}

		case Delete:

			fmt.Print("Введите значение для удаления: \n>>> ")
			deleteElement, errRead := reader.ReadString('\n')

			if errRead != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при чтении ввода:", errRead)
				continue
			}

			errDelete := structure.Remove(strings.TrimSpace(deleteElement))

			if errDelete != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при удалении значения:", errDelete)
				continue
			} else {
				fmt.Println("Успешно удалено!")
				fmt.Printf("Размерность: %d", structure.Size())
			}

		case SearchByValue:

			fmt.Print("Введите значение для поиска: \n>>> ")
			searchElement, errRead := reader.ReadString('\n')

			if errRead != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при чтении ввода:", errRead)
				continue
			}

			valByValue, position, errSearch := structure.GetByValue(strings.TrimSpace(searchElement))

			if errSearch != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при поиске значения:", errSearch)
				continue
			} else {
				fmt.Printf("Значение - %s\nПозиция - %d", valByValue, position)
			}

		case SearchByPosition:

			fmt.Print("Введите позицию искомого значения: \n>>> ")
			searchPosition, errRead := reader.ReadString('\n')

			if errRead != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при чтении ввода:", errRead)
				continue
			}

			position, errConverting := strconv.Atoi(strings.TrimSpace(searchPosition))

			if errConverting != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				log.Println("Неверный ввод! Введите позицию в виде числа.")
				continue
			}

			valByPosition, errSearch := structure.Get(position)

			if errSearch != nil {
				mw := io.MultiWriter(os.Stdout, file)
				log.SetOutput(mw)
				fmt.Println("Ошибка при поиске значения:", errSearch)
				continue
			} else {
				fmt.Println(valByPosition)
			}

		case Exit:

			fmt.Println("Выход...")
			os.Exit(0)

		default:

			mw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(mw)
			log.Println("Неверный выбор. Введите корректный пункт меню.")
		}
	}
}
