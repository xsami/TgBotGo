package helpers

import (
	"os"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

// LoadEnv this function load the .env file
func LoadEnv(path ...string) error {
	var err error
	if len(path) > 0 {
		err = godotenv.Load(
			path[0] + string(os.PathSeparator) + ".env")
	} else {
		err = godotenv.Load()
	}
	return err
}

// Getenv return a string with the value
// of the given environment key
func Getenv(key string) string {
	return os.Getenv(key)
}

// GetBotToken return a string with the token of the TG bot
func GetBotToken() string {
	return Getenv("TG_API_TOKEN")
}

// getBaseDir return a string with the base path of the
// directory when the function is called from the helper
// directory itself
func getBaseDir() string {
	dir, err := os.Getwd()

	if err != nil {
		return ""
	}

	var ss []string
	var currentDirName string

	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
		currentDirName = strings.Join(ss[:len(ss)-2], "\\")
	} else {
		ss = strings.Split(dir, "/")
		currentDirName = strings.Join(ss[:len(ss)-2], "/")
	}

	return currentDirName
}
