package leho

import (
	"fmt"
	"testing"
)

const File = "test.json"

type Settings struct {
	File string
	Name string
}

var settings = Settings{
	File: "file2",
	Name: "name2",
}
var leho = Leho{
	//Setting:  settings,
	FileName: File,
}

func TestLeho_WriteSetting(t *testing.T) {
	err := leho.WriteSetting(settings)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestLeho_ReadSetting(t *testing.T) {
	err := leho.ReadSetting(&settings)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}
	fmt.Printf("\tFile: %s, name: %s\n", settings.File, settings.Name)
}
