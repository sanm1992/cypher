package database

import (
	"fmt"
	"os"
)

var (
	DB *os.File
)

const dbPath = ".cypher.json"

func InitDB() *os.File {
	homeDir, err := os.UserHomeDir()
	dir := homeDir + "/" + dbPath

	var file *os.File
	if err != nil {
		panic(err)
	}

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
