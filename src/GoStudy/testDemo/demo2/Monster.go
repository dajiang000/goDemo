package demo2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name   string
	Age    int
	Gender string
	Score  float64
	Skill  string
}

func (monster *Monster) Store(storePath string) bool {
	file, err := os.OpenFile(storePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("error is", err)
		return false
	}
	data, err := json.Marshal(monster)
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	defer file.Close()
	return true
}

func (monster *Monster) Restore(storePath string) bool {
	data, err := os.ReadFile(storePath)
	if err != nil {
		fmt.Println("err is", err)
		return false
	}
	err = json.Unmarshal(data, &monster)
	if err != nil {
		fmt.Println("error is", err)
		return false
	}

	return true
}
