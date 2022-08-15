package config

import "testing"

const configData = `
{
	"project": "Bytehack Config",
	"version": "v0.0.0",
	"number": 100
}
`

func TestConfig(t *testing.T) {
	// Create config
	config, err := NewConfig([]byte(configData))
	if err != nil {
		t.Fail()
	}
	// Read string from the config
	stringValue, err := config.GetString("project")
	if err != nil {
		t.Log("Fail reading from config: ", err.Error())
		t.FailNow()
	}

	if stringValue != "Bytehack Config" {
		t.Log("Fail expected Bytehack Config found ", stringValue)
		t.FailNow()
	}

	// Read int from the config
	intValue, err := config.GetInt("number")
	if err != nil {
		t.Log("Fail reading int from config: ", err.Error())
		t.Fail()
	}
	if intValue != 100 {
		t.Log("Fail expected", 100, " found ", intValue)
		t.Fail()
	}
}
