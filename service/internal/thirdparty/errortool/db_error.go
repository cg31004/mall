package errortool

import (
	"database/sql"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func DBErrorPackage(groups iGroupRepository, codes iCodeRepository) interface{} {
	groups.Add(groupCodeDB)
	group := &errorGroup{
		codes:     codes,
		groups:    groups,
		groupCode: groupCodeDB,
	}

	return dbError{
		NoRow:                group.Error("000", "DB no row"),
		CannotCreateTable:    group.Error("001", "DB cannot create table"),
		CannotCreateDatabase: group.Error("002", "DB cannot create database"),
		DatabaseCreateExists: group.Error("003", "DB database create exists"),
		TooManyConns:         group.Error("004", "DB too many conns"),
		AccessDenied:         group.Error("005", "DB access denied"),
		UnknownTable:         group.Error("006", "DB unknown table"),
		DuplicateEntry:       group.Error("007", "DB duplicate entry"),
		NoDefaultForField:    group.Error("008", "DB no default for field"),
	}
}

type dbError struct {
	NoRow                error
	CannotCreateTable    error
	CannotCreateDatabase error
	DatabaseCreateExists error
	TooManyConns         error
	AccessDenied         error
	UnknownTable         error
	DuplicateEntry       error
	NoDefaultForField    error
}

var (
	dbErrorCode = map[int]error{
		1005: ErrDB.CannotCreateTable,
		1006: ErrDB.CannotCreateDatabase,
		1007: ErrDB.DatabaseCreateExists,
		1040: ErrDB.TooManyConns,
		1045: ErrDB.AccessDenied,
		1051: ErrDB.UnknownTable,
		1062: ErrDB.DuplicateEntry,
		1364: ErrDB.NoDefaultForField,
	}
)

func ConvertDB(err error) error {
	if err == sql.ErrNoRows || err == gorm.ErrRecordNotFound {
		return ErrDB.NoRow
	}

	if e, ok := parseDBError(err); ok {
		return e
	}

	return err
}

func parseDBError(err error) (error, bool) {
	s := strings.TrimSpace(err.Error())
	data := strings.Split(s, ":")
	if len(data) == 0 {
		return nil, false
	}

	numStr := strings.ToLower(data[0])
	numStr = strings.Replace(numStr, "error", "", -1)
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, false
	}

	e, ok := dbErrorCode[num]
	if !ok {
		return nil, false
	}

	return e, true
}
