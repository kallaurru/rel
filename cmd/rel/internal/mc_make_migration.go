package internal

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type migrationTmplData struct {
	MigrateFuncName  string
	RollbackFuncName string
	Package          string
}

const migrationMakeTemplate = `
package {{.Package}}

import (
    "github.com/go-rel/rel"
)

const tablename = "define your table name"

// {{.MigrateFuncName}} definition
func {{.MigrateFuncName}}(schema *rel.Schema) {
    schema.CreateTable(tablename, func(t *rel.Table) {
    
	})
}

// {{.RollbackFuncName}} definition
func {{.RollbackFuncName}}(schema *rel.Schema) {
    schema.DropTable(tablename)
}

`

func ExecMakeMigration(args []string) error {
	if len(args) < 3 {
		return MsgWithUserText("you need to specify the command = make and migration suffix")
	}
	if len(args[2]) > 128 {
		return MsgWithUserText("name migration longer at 128 symbols")
	}
	prefix := makeMigrationPrefixName()
	suffix := os.Args[2]
	root, err := filepath.Abs(filepath.Dir(args[0]))
	if err != nil {
		return err
	}
	parts := strings.Split(suffix, "_")
	if len(parts) < 2 {
		return MsgWithUserText("migration name not camel_case written")
	}
	filename := fmt.Sprintf("%s_%s.go", prefix, suffix)
	fullFileName := makeMigrationFileName(filename, root)
	if fullFileName == "" {
		return MsgWithUserText("full name migration file not built")
	}
	migrateFuncName := fmt.Sprintf("Migrate%s", strcase.ToCamel(suffix))
	rollbackFuncName := fmt.Sprintf("Rollback%s", strcase.ToCamel(suffix))
	data := migrationTmplData{
		Package:          os.Getenv(ProjectAliasKey),
		MigrateFuncName:  migrateFuncName,
		RollbackFuncName: rollbackFuncName,
	}
	tmpl := template.Must(
		template.New("make_migration").Parse(migrationMakeTemplate))
	writeTemplateToFile(tmpl, fullFileName, data)

	return nil
}

func makeMigrationPrefixName() string {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	momentDay := t.Second() + 60*t.Minute() + 3600*t.Hour()

	return fmt.Sprintf(
		"%d%0d%d%d_",
		year,
		month,
		day,
		momentDay)
}

func makeMigrationFileName(filename, root string) string {
	migrDir, err := GetMigrationProjectDir(root)
	if err != nil {
		return ""
	}
	return filepath.Join(
		"/tmp",
		migrDir,
		filename)
}

func writeTemplateToFile(tmpl *template.Template, filename string, data migrationTmplData) {
	file, err := os.Create(filename)
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			check(err)
		}
	}(file)
	err = tmpl.Execute(file, data)
	check(err)
}
