package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"sync"
)

type User struct {
	id       int
	email    string
	password string
}

type InMemoryUserStorage struct {
	lock    sync.RWMutex
	storage map[string]User
}

func NewInMemoryUserStorage() *InMemoryUserStorage {
	return &InMemoryUserStorage{
		lock:    sync.RWMutex{},
		storage: make(map[string]User),
	}
}

type UserRepository interface {
	Add(string, User) error
	Get(string) (User, error)
}

type UserService struct {
	repository UserRepository
}

type UserRegisterParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *InMemoryUserStorage) Add(key string, user User) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.storage[key].email != "" {
		return errors.New("user already exist")
	}

	s.storage[key] = user
	return nil
}

func (s *InMemoryUserStorage) Get(key string) (user User, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	user, exists := s.storage[key]
	if exists {
		return user, nil
	}
	return (User{}), errors.New("Key '" + key + "' doesn't exist")
}

func validateRegisterParams(p *UserRegisterParams) error {
	match, _ := regexp.Match(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, []byte(p.Email))
	if !match {
		return errors.New("invalid email address")
	}

	if len(p.Password) < 8 {
		return errors.New("password too short")
	}

	return nil
}

func (u *UserService) Register(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	params := &UserRegisterParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}
	if err := validateRegisterParams(params); err != nil {
		handleError(err, w)
		return
	}
	userCounter++
	password := md5.New().Sum([]byte(params.Password))
	newUser := User{
		id:       userCounter,
		email:    params.Email,
		password: string(password),
	}
	err = u.repository.Add(params.Email, newUser)
	if err != nil {
		handleError(err, w)
		return
	}

	globalTodos.todos = make(map[int]UserToDo)

	globalTodos.todos[userCounter] = UserToDo{
		lists: make(map[int]ToDoList),
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("registered"))
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(err.Error()))
}
