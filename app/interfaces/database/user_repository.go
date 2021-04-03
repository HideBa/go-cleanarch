package database

import "github.com/HideBa/go-cleanarch/app/domain"

type UserRepository struct {
	SqlHandler
}

func (r *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := r.Execute("INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return id, nil
}

func (r *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := r.Query("SELECT * from users where id = ?", identifier)
	defer row.Close()
	if err != nil {
		return domain.User{}, err
	}
	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return user, nil
}

func (r *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := r.Query("SELECT * from users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return users, nil
}
