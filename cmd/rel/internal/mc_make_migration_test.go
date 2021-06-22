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
		tmpRoot := t.TempDir()
		var (
			projectAlias      = "ydict"
			tablenameWithVerb = "create_ru_words"
			args              = []string{
				filepath.Join(tmpRoot, "rel"),
				"make",
				tablenameWithVerb,
			}
		)
		err := makeTestDirs(tmpRoot, projectAlias)
		if err != nil {
			t.Error(err)
		}
		_ = os.Setenv(ProjectAliasKey, projectAlias)
		assert.Equal(t, nil, ExecMakeMigration(args))
		// проверка существования физического файла миграции
		targetPath := filepath.Join(
			tmpRoot,
			RootMigrationPath,
			projectAlias)
		err = slUtils.FindFilesWithSuffixInPath(targetPath, "go")
		// если нет ошибок, значит в указанном каталоге найдены файлы go
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func makeTestDirs(rootPath, projectPath string) error {
	path := filepath.Join(rootPath, RootMigrationPath, projectPath)

	return os.MkdirAll(path, 0755)
}
