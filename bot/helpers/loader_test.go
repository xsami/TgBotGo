package helpers

import (
	"testing"
)

func TestLoadEnv(t *testing.T) {

	if err := LoadEnv(getBaseDir()); err != nil {
		t.Errorf("Failed loading .env file, got: %v", err)
	}

}

func TestGetBotToken(t *testing.T) {

	if GetBotToken() == "" {
		t.Errorf("Failed at loading bot token. Got \"\".")
	}
}

func TestGetEnviromentVariable(t *testing.T) {

}

func TestGetBaseDir(t *testing.T) {

}
