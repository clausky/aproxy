package auth

import (
	"encoding/base64"
	"errors"

	"encoding/json"

	"fmt"

	"github.com/imroc/req"
)

type RedmineUserStorage struct {
	Host string
}

func (self *RedmineUserStorage) Login(email, pwd string) (*User, error) {
	enc := base64.StdEncoding.EncodeToString([]byte(email + ":" + pwd))
	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic " + enc,
	}
	r, err := req.Get(self.Host+"/users/current.json", header)
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, errors.New("1000:email or password wrong.")
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(r.Bytes(), &data)
	if err != nil {
		return nil, err
	}
	if userI, ok := data["user"]; ok {
		user := userI.(map[string]interface{})
		return &User{
			Id:    fmt.Sprintf("%v", user["id"]),
			Email: email,
			Name:  email,
		}, nil
	}
	return nil, errors.New("1002: unkonw error.")
}

func (self *RedmineUserStorage) GetByEmail(email string) (u *User, err error) {
	return nil, errors.New("GetByEmail Failed.")
}

func (self *RedmineUserStorage) GetAll() (us []User, err error) {
	return nil, errors.New("GetAll Failed.")
}

func (self *RedmineUserStorage) Insert(user User) (err error) {
	return errors.New("Insert Failed.")
}

func (self *RedmineUserStorage) Update(id string, user User) (err error) {
	return errors.New("Update Failed.")
}

func (self *RedmineUserStorage) Delete(id string) (err error) {
	return errors.New("Delete Failed.")
}
