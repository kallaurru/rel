package internal

import (
	"fmt"
	slUtils "github.com/kallaurru/utils"
	"os"
	"path/filepath"
)

const (
	// RootEnvFileKey регистрируем только имя файла
	RootEnvFileKey    = "MC_ROOT_ENV_FILE"
	ProjectAliasKey   = "MC_PROJECT_ALIAS"
	RootMigrationPath = "migrations"
)

/*
	MC - migration center. Хочу объединить миграции к нескольким проектам в одно место.
	Стратегии:
		1. Через переменную окружения MC_ROOT_ENV_FILE указываем имя файла окружения с которым будет работать утилита
		2. Если будет найден файл из пункта 1. Загружаем его. Обязательно нужно указать каталог проекта в котором лежат
		   миграции к данному проекту.
*/

// GetMigrationProjectDir - возвращаем относительный путь от корня проекта
// Например: migrations/ydict
// @param path - путь до проекта для проверки корректности вложенных каталогов
func GetMigrationProjectDir() (string, error) {
	projectAlias := os.Getenv(ProjectAliasKey)
	if projectAlias == "" {
		msg := fmt.Sprintf("project alias env var with key %s has not value", ProjectAliasKey)
		return "", MsgWithUserText(msg)
	}
	return filepath.Join(RootMigrationPath, projectAlias), nil
}

// ValidateRootEnvFile  - работает в контексте каталога проекта
func ValidateRootEnvFile(path string) string {
	path = filepath.Dir(path)
	// приходит только имя файла
	value := os.Getenv(RootEnvFileKey)
	if value != "" {
		isValid := slUtils.ValidateFile(
			filepath.Join(path, value),
			1024*64)
		if isValid {
			return filepath.Join(path, value)
		}
	}

	return ""
}
