package start

import (
	"GoLandHomeworks/homeworks/Homework_22/redis_lib"
	"GoLandHomeworks/homeworks/Homework_22/users"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpEndpoints interface {
	Endpoint22() func(w http.ResponseWriter, r *http.Request)
	Endpoint22WithParameter(idParam string) func(w http.ResponseWriter, r *http.Request)
	PostEndpoint22() func(w http.ResponseWriter, r *http.Request)
	RegisterEndpoint22() func(w http.ResponseWriter, r *http.Request)
	LoginEndpoint22() func(w http.ResponseWriter, r *http.Request)
	ProfileEndpoint22() func(w http.ResponseWriter, r *http.Request)
}

type httpEndpoints struct {
	usersStore users.UsersStore
	redisStore *redis_lib.RedisStore
}

func NewHttpEndpoints(uS users.UsersStore, rS *redis_lib.RedisStore) HttpEndpoints {
	return &httpEndpoints{usersStore: uS, redisStore: rS}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func (h *httpEndpoints) Endpoint22() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := users.User{
			Id:        "100513974",
			Username:  "TestUsername1",
			Password:  "Qwerty11!",
			FirstName: "Dana",
			LastName:  "White",
			Avatar:    "picture1",
		}
		respondJSON(w, http.StatusOK, user)
		return
	}
}

func (h *httpEndpoints) Endpoint22WithParameter(idParam string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars[idParam]
		if !ok {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    "Please write parameter ID",
				StatusCode: http.StatusBadRequest,
			})
		}
		var response users.User
		usersData := []users.User{
			{
				Id:        "1",
				Username:  "Cool_Dude",
				Password:  "qweasdzxc",
				FirstName: "Alex",
				LastName:  "Hopkins",
				Avatar:    "picture2",
			},
			{
				Id:        "2",
				Username:  "FeelsBadMan",
				Password:  "rtyfghvbn",
				FirstName: "Jack",
				LastName:  "Smith",
				Avatar:    "picture3",
			},
		}
		if idStr == "1" {
			response = usersData[0]
		} else if idStr == "2" {
			response = usersData[1]
		} else {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    "User with this ID not found",
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		respondJSON(w, http.StatusOK, response)
		return
	}
}

func (h *httpEndpoints) PostEndpoint22() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		user := &users.User{}
		err = json.Unmarshal(jsonData, &user)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		user.Id = "101"
		respondJSON(w, http.StatusCreated, user)
		return
	}
}

func (h *httpEndpoints) RegisterEndpoint22() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		user := &users.User{}
		err = json.Unmarshal(jsonData, &user)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		response, err := h.usersStore.Create(user)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusCreated, response)
		return
	}
}

func (h *httpEndpoints) LoginEndpoint22() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		req := &LoginRequest{}
		err = json.Unmarshal(jsonData, &req)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		user, err := h.usersStore.GetByUsernameAndPassword(req.Username, req.Password)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		key := uuid.New().String()
		err = h.redisStore.SetValue(key, user, 1*time.Minute)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		response := &LoginResponse{AccessKey: key}
		respondJSON(w, http.StatusOK, response)
		return
	}
}

func (h *httpEndpoints) ProfileEndpoint22() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accessKey := r.Header.Get("Authorization")
		if accessKey == "" {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    "Unauthorized access",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		user := &users.User{}
		err := h.redisStore.GetValue(accessKey, &user)
		if err != nil {
			errorMessage := err.Error()
			if err.Error() == "redis: nil" {
				errorMessage = "Your access key was expired"
			}
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    errorMessage,
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		response, err := h.usersStore.Get(user.Id)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		respondJSON(w, http.StatusOK, response)
		return
	}
}
