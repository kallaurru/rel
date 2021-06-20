package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestExecMakeMigration(t *testing.T) {
	t.Run("make migrate from template", func(t *testing.T) {
		var (
			args = []string{
				filepath.Join(os.TempDir(), "rel"),
				"make",
				"create_ru_words",
			}
		)
		assert.Equal(t, nil, ExecMakeMigration(args))

	})
}
