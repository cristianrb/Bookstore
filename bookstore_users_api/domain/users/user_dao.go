package users

import (
	"cristianrb/datasources/mysql/users_db"
	"cristianrb/logger"
	"cristianrb/utils/errors"
	"cristianrb/utils/mysql_utils"
	"fmt"
)

const queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
const queryGetUser = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
const queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
const queryDeleteUser = "DELETE FROM users WHERE id=?;"
const queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created FROM users WHERE status=?;"

func (user *User) Get() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return mysql_utils.ParseError(err)
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user by id", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return mysql_utils.ParseError(err)
	}
	defer statement.Close()

	insertResult, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return mysql_utils.ParseError(err)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update statement", err)
		return mysql_utils.ParseError(err)
	}
	defer statement.Close()

	_, updateErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		logger.Error("error when trying to update user", updateErr)
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete statement", err)
		return mysql_utils.ParseError(err)
	}
	defer statement.Close()

	_, deleteErr := statement.Exec(user.Id)
	if deleteErr != nil {
		logger.Error("error when trying to delete user", deleteErr)
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find by status statement", err)
		return nil, mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find by status", err)
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to retrieve user in find by status", err)
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
