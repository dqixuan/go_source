
// DO NOT EDIT. THIS IS GENERATED BY TIME
package main

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type TableNameDao struct {}

func NewTableNameDao() *TableNameDao {
	return &TableNameDao{}
}

var TABLENAMECol = []string {
 
"I_ID",
 
"I_USER_ID",
 
"I_TYPE",
 
"I_STATUS",
 
"D_CREATED_AT",
 
"D_UPDATED_AT",

}

func (tablenameDao *TableNameDao ) selectByID(ctx context.Context, db sqlx.DB, ID int64) (err error) {
	reture
}
