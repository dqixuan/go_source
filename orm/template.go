package main

const textTemplate = `
// DO NOT EDIT. THIS IS GENERATED BY TIME
package main

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type {{.TableName}}Dao struct {}

func New{{.TableName}}Dao() *{{.TableName}}Dao {
	return &{{.TableName}}Dao{}
}

var {{ToTitle .TableName}}Col = []string {
{{ range .Columns }} 
"{{- .ColumnName -}}",
{{end}}
}

func ({{toLowerCamelCase .TableName}}Dao *{{.TableName}}Dao ) selectByID(ctx context.Context, db sqlx.DB, ID int64) (err error) {
	reture
}
`
