package sqldb

import (
	"database/sql"
	"time"
	"user-service/app/logger"
	"user-service/domain/model"
	"user-service/tool/timea"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jfeng45/gtransaction/gdbc"
	"github.com/pkg/errors"
)

const (
	INSERT_USER = "INSERT users SET fullName=?,email=?,password=?,created=?"
	QUERY_USER_BY_EMAIL = "SELECT * FROM users WHERE email=?"
)

type UserDataSql struct {
	DB gdbc.SqlGdbc
}

func (uds *UserDataSql) Insert(user *model.User) (*model.User, error) {
	stmt, err := uds.DB.Prepare(INSERT_USER)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.FullName, user.Email, user.Password, user.Created)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	user.Id = int(id)
	logger.Log.Debug("user inserted: ", user)
	return user, nil
}

func (uds *UserDataSql) FindByEmail(email string) (*model.User, error) {
	rows, err := uds.DB.Query(QUERY_USER_BY_EMAIL, email)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	if rows.Next() {

	}
	return nil, nil
}

func rowsToUser(rows *sql.Rows) (*model.User, error) {
	var ds string
	user := &model.User{}
	err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &ds)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	created, err := time.Parse(timea.FORMAT_ISO8601_DATE, ds)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	user.Created = created
	logger.Log.Debug("rows to user: ", user)
	return user, nil
}