package internal

import (
	slUtils "github.com/kallaurru/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestExecMakeMigration(t *testing.T) {
	t.Run("make migrate from template", func(t *testing.T) {
		var (
			projectAlias      = "ydict"
			tablenameWithVerb = "create_ru_words"
			args              = []string{
				filepath.Join(os.TempDir(), "rel"),
				"make",
				tablenameWithVerb,
			}
		)
		err := makeTestDirs(t, projectAlias)
		if err != nil {
			t.Error(err)
		}
		_ = os.Setenv(ProjectAliasKey, projectAlias)
		assert.Equal(t, nil, ExecMakeMigration(args))
		// проверка существования физического файла миграции
		targetPath := filepath.Join(
			os.TempDir(),
			RootMigrationPath,
			projectAlias)
		err = slUtils.FindFilesWithSuffixInPath(targetPath, "go")
		// если нет ошибок, значит в указанном каталоге найдены файлы go
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func makeTestDirs(t *testing.T, projectPath string) error {
	path := filepath.Join(t.TempDir(), RootMigrationPath, projectPath)

	return os.MkdirAll(path, 0755)
}
