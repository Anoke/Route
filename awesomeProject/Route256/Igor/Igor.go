package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Directory struct {
	Dir     string      `json:"dir"`
	Files   []string    `json:"files"`
	Folders []Directory `json:"folders"`
}

func main() {
	var t int
	fmt.Scan(&t)

	for test := 0; test < t; test++ {
		var n int
		fmt.Scan(&n)

		// Создаем пустую структуру для текущей директории
		var root Directory

		decoder := json.NewDecoder(os.Stdin)

		for i := 0; i < n; i++ {
			// Используем map[string]interface{} для декодирования динамической структуры JSON
			var data map[string]interface{}
			err := decoder.Decode(&data)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}

			// Преобразуем map в структуру Directory
			dir, err := mapToDirectory(data)
			if err != nil {
				fmt.Println("Error converting map to Directory:", err)
				return
			}

			// Добавляем полученную структуру в корневую директорию
			root = dir
		}

		// Считаем количество зараженных файлов, начиная с корневой директории
		infectedFiles := countInfectedFiles(root)
		fmt.Println(infectedFiles)
	}
}

// Функция для преобразования map в структуру Directory
func mapToDirectory(data map[string]interface{}) (Directory, error) {
	dir := Directory{
		Dir: data["dir"].(string),
	}

	// Преобразуем файлы и поддиректории, если они присутствуют
	if files, ok := data["files"].([]interface{}); ok {
		for _, file := range files {
			dir.Files = append(dir.Files, file.(string))
		}
	}

	if folders, ok := data["folders"].([]interface{}); ok {
		for _, folder := range folders {
			subdir, err := mapToDirectory(folder.(map[string]interface{}))
			if err != nil {
				return Directory{}, err
			}
			dir.Folders = append(dir.Folders, subdir)
		}
	}

	return dir, nil
}

// Функция для подсчета количества зараженных файлов в директории и ее поддиректориях
func countInfectedFiles(directory Directory) int {
	infectedFiles := 0

	// Проверяем каждый файл в текущей директории
	for _, file := range directory.Files {
		if strings.HasSuffix(file, ".hack") {
			infectedFiles++
		}
	}

	// Рекурсивно проверяем каждую поддиректорию
	for _, subdirectory := range directory.Folders {
		infectedFiles += countInfectedFiles(subdirectory)
	}

	// Если текущая директория заражена, добавляем количество файлов в ней к общему количеству зараженных файлов
	if infectedFiles > 0 || containsVirus(directory) {
		infectedFiles += len(directory.Files)
	}

	return infectedFiles
}

// Функция для проверки, содержит ли директория вирусные файлы
func containsVirus(directory Directory) bool {
	for _, file := range directory.Files {
		if strings.HasSuffix(file, ".hack") {
			return true
		}
	}
	return false
}
