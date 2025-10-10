package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type service interface {
	AddUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*User, error)
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	return &Endpoint{service: service, log: log}
}

type Endpoint struct {
	service service
	log     *slog.Logger
}

func (e *Endpoint) GetUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID := uuid.MustParse(vars["id"])

	user, err := e.service.GetUser(request.Context(), userID)
	if err != nil {
		e.log.Error("error in method Endpoint.GetUser: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&user); err != nil {
		e.log.Error(err.Error())
	}
}

func (e *Endpoint) AddUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	user.UserID = uuid.New()
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		e.log.Error("error in method Endpoint.AddUser:" + err.Error())
	}

	addedUser, err := e.service.AddUser(request.Context(), &user)
	if err != nil {
		e.log.Error("error in method Endpoint.AddUser: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&addedUser); err != nil {
		e.log.Error(err.Error())
	}
}

func (e *Endpoint) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		e.log.Error(err.Error())
	}

	updatedUser, err := e.service.UpdateUser(request.Context(), &user)
	if err != nil {
		e.log.Error("error in method Endpoint.UpdateUser: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&updatedUser); err != nil {
		e.log.Error("error in method Endpoint.UpdateUser:" + err.Error())
	}
}

func (e *Endpoint) DeleteUser(_ http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID := uuid.MustParse(vars["id"])
	fmt.Println(userID)
	err := e.service.DeleteUser(request.Context(), userID)

	if err != nil {
		e.log.Error("error in method Endpoint.DeleteUser" + err.Error())
	}
}

func (e *Endpoint) GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := e.service.GetAllUsers(request.Context())
	if err != nil {
		e.log.Error("error in method Endpoint.GetAllUsers" + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&users); err != nil {
		e.log.Error("error in method Endpoint.GetAllUsers: " + err.Error())
	}
}
