package database

import (
	"fmt"
	"os"
)

var (
	DB *os.File
)

func DBPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return homeDir, err
	}

	dir := homeDir + "/" + ".cypher.json"

	return dir, nil
}

func InitDB() *os.File {
	dir, err := DBPath()
	if err != nil {
		panic(err)
	}

	var file *os.File

	if _, err := os.Stat(dir); err != nil {
		// 自动创建文件并写入空内容
		file, err = os.Create(dir)
		fmt.Fprintf(file, "[]")
	}

	file, err = os.OpenFile(dir, os.O_RDWR|os.O_CREATE, 0644)
	os.Stat(dir)
	if err != nil {
		panic(err)
	}
	return file
}

func GetDB() *os.File {
	return InitDB()
}
