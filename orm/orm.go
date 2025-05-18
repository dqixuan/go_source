package main

import (
	"bytes"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
	"text/template"
)

// user:password@/dbname
func main() {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_self")
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	var cols []*Column
	q := "SELECT COLUMN_NAME, DATA_TYPE FROM information_schema.`COLUMNS` where TABLE_SCHEMA = 'test_self' and TABLE_NAME = 'user';"
	err = db.SelectContext(ctx, &cols, q)
	if err != nil {
		panic(err)
	}
	obj := TargetTableInfo{
		TableName: "test_self",
	}

	for i, c := range cols {
		fmt.Println("no", i, "=", c)
		obj.Columns = append(obj.Columns, Column{
			ColumnName: c.ColumnName,
			DataType:   c.DataType,
		})
	}
	defer db.Close()

	temp, err := template.New("s").Funcs(template.FuncMap{
		"ToUpper":          ToUpper,
		"ToTitle":          ToTitle,
		"toLowerCamelCase": toLowerCamelCase,
	}).Parse(textTemplate)
	if err != nil {
		panic(fmt.Errorf("template new failed, err=%w", err))
	}
	//f, err := os.Open("name")
	var buf bytes.Buffer

	err = temp.Execute(&buf, obj)
	if err != nil {
		panic(fmt.Errorf("execute failed, err=%w", err))
	}
	f, err := os.Create("name.go")
	if err != nil {
		panic(fmt.Errorf("open, err=%w", err))
	}
	f.Write(buf.Bytes())
	defer f.Close()

}

// mysql字段类型转化为 golang 数据类型
var mysqlType2GoType = map[string]string{
	"tinyint":  "int32",
	"bigint":   "int64",
	"varchar":  "string",
	"datetime": "time.Time",
	"bool":     "bool",
	"int":      "int32",
}

// 目标数据库所需信息
type TargetTableInfo struct {
	TableName       string   // 表名
	Columns         []Column // 字段名称和类型
	ModelOutputPath string
	DaoOutputPath   string
}

/*
 */

type Column struct {
	ColumnName string `db:"COLUMN_NAME"`
	DataType   string `db:"DATA_TYPE"`
}

type SpecialField struct {
}

var cfg *Config

type Config struct {
	Database string
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToTitle(s string) string {
	return strings.ToTitle(s)
}

func toLowerCamelCase(s string) string {
	sli := strings.Split(s, "_")
	ans := ""
	for i, s1 := range sli {
		if i == 0 {
			ans += strings.ToLower(s1)
		} else {
			runeSlice := []rune(s1)
			ans += strings.ToTitle(string(runeSlice[:1])) + strings.ToLower(string(runeSlice[1:]))
		}
	}
	return ans
}
