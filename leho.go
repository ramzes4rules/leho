package leho

import (
	"encoding/json"
	"os"
)

type Leho struct {
	Setting  any    // Setting object
	FileName string // Settings file
}

func (leho *Leho) WriteSetting() error {
	// Сериализуем данные
	var response, err = json.MarshalIndent(leho.Setting, "", "\t")
	if err != nil {
		return err
	}

	// Создаем файл
	file, err := os.Create(leho.FileName)
	if err != nil {
		return err
	}

	// Пишем в файл
	_, err = file.WriteString(string(response))
	if err != nil {
		return err
	}

	// Закрываем
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func (leho *Leho) ReadSetting() error {

	data, err := os.ReadFile(leho.FileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &leho.Setting)
	if err != nil {
		return err
	}
	return nil
}
