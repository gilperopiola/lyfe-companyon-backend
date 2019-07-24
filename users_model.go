package main

import (
	"fmt"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

func (user *User) Login() (*User, error) {
	if err := db.DB.QueryRow(`SELECT id FROM users WHERE email = ? AND password = ?`, user.Email, user.Password).Scan(&user.ID); err != nil {
		return &User{}, err
	}

	user.Token = generateToken(*user)
	return user, nil
}

func (user *User) Create() (*User, error) {
	result, err := db.DB.Exec(`INSERT INTO users (email, password, firstName, lastName) VALUES (?, ?, ?, ?)`, user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return &User{}, err
	}

	user.ID = utils.GetEntryID(result)

	return user.Get()
}

func (user *User) Get() (*User, error) {
	if err := db.DB.QueryRow(`SELECT email, password, firstName, lastName, enabled, dateCreated FROM users WHERE id = ?`, user.ID).Scan(
		&user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Enabled, &user.DateCreated); err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) Update() (*User, error) {
	_, err := db.DB.Exec(`UPDATE users SET email = ?, firstName = ?, lastName = ?, enabled = ? WHERE id = ?`,
		user.Email, user.FirstName, user.LastName, user.Enabled, user.ID)
	if err != nil {
		return &User{}, err
	}

	return user.Get()
}

func (user *User) Search(params *SearchParameters) ([]*User, error) {
	query := fmt.Sprintf(`SELECT id FROM users WHERE id LIKE ? OR email LIKE ? OR firstName LIKE ? OR lastName LIKE ?
						  	ORDER BY %s LIMIT ? OFFSET ?`, getSearchOrderBy(params))

	params.Filter = "%" + params.Filter + "%"
	rows, err := db.DB.Query(query, params.Filter, params.Filter, params.Filter, params.Filter, params.Limit, params.Offset)
	defer rows.Close()
	if err != nil {
		return []*User{}, err
	}

	users := []*User{}
	for rows.Next() {
		tempUser := &User{}
		if err = rows.Scan(&tempUser.ID); err != nil {
			return []*User{}, err
		}

		tempUser, err = tempUser.Get()
		if err != nil {
			return []*User{}, err
		}

		users = append(users, tempUser)
	}

	return users, nil
}
