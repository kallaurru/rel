package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// чтение рабочего файла env. Положительный и негативный случай
func TestRootEnvFile(t *testing.T) {
	targetPath := filepath.Join(t.TempDir(), "rootEnv")
	envFileName := ".env.ydict"
	err := os.MkdirAll(targetPath, 0755)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Run("file not set", func(t *testing.T) {
		TmpEnvFile(filepath.Join(targetPath, envFileName), t)
		assert.Equal(t, "", ValidateRootEnvFile(targetPath))
	})

	// положительный тест
	t.Run("file is set", func(t *testing.T) {
		err := os.Setenv(RootEnvFileKey, envFileName)
		if err != nil {
			t.Fatal(err.Error())
		}
		assert.Equal(t, envFileName, ValidateRootEnvFile(targetPath))
		_ = os.Setenv(RootEnvFileKey, "")
	})
}

func TmpEnvFile(path string, t *testing.T) {
	f, err := os.Create(path)
	if err != nil {
		t.Error("tmp env file not created")
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			t.Error("tmp env closed with problem")
			return
		}
	}(f)
	_, _ = fmt.Fprintln(f, ProjectAliasKey, "=", "ydict")
}
