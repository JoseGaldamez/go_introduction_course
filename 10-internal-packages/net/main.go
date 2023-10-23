package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Data struct {
	User User `json:"data"`
}

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirrtName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserCreated struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Job       string     `json:"job"`
	Age       int        `json:"age,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	DeleteAt  *time.Time `json:"deleteAt,omitempty"`
}

const (
	base = "https://reqres.in/api/users"
)

func main() {
	resp, err := GetReqExample("https://reqres.in/api/users/2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resp))

	fmt.Println()

	res, err := Get("https://reqres.in/api/users/2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))

	fmt.Println()

	user, err := GetUser("2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	fmt.Println()

	newuser, err := CreateUser("Jose", "Desarrollador")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newuser)

	userJson, err := json.Marshal(newuser)
	fmt.Println(string(userJson))

}

func GetReqExample(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func GetUser(userID string) (*User, error) {
	req, err := Get(fmt.Sprintf("%s/%s", base, userID))
	if err != nil {
		return nil, err
	}

	var dataResponse Data

	if err := json.Unmarshal(req, &dataResponse); err != nil {
		return nil, err
	}

	return &dataResponse.User, nil
}

func Post(url string, data interface{}) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func CreateUser(name, job string) (*UserCreated, error) {
	user := &UserCreated{
		Name: name,
		Job:  job,
	}

	response, err := Post(base, user)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(response, user); err != nil {
		return nil, err
	}

	return user, nil
}
