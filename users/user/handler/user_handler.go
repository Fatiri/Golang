package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"users/model"
	"users/user"
	"users/utils"

)

type UserHandler struct {
	userUseCase user.UserUseCase
}

func (h UserHandler) register(writer http.ResponseWriter, request *http.Request) {
	var dataUser model.User

	err := json.NewDecoder(request.Body).Decode(&dataUser)
	if err != nil {
		fmt.Printf("Error when decodea data user %v\n", err)
		utils.HandleError(writer, "Error when decode data user", http.StatusBadRequest)
		//Fungsi return di sini dia untuk menghentikan proses kenpa tidak memakai break , karena break pda golang di khusus untuk switch case
		return
	}

	data, err := h.userUseCase.Register(&dataUser)
	if err != nil {
		fmt.Printf("[UserHandler.register] Got Error when register %v\n", err)
		utils.HandleError(writer, "Ops.. something wrong", http.StatusBadRequest)
		return
	}
	utils.HandleSuccess(writer, http.StatusCreated, data)
}

func (h UserHandler) getById(writer http.ResponseWriter, request *http.Request) {
	pathVariable := mux.Vars(request)
	// strconv.Atoi itu untuk convert dari string ke integer strconv.Itoa itu kebalikanya dari Integer ke string
	id, err := strconv.Atoi(pathVariable["id"])
	if err != nil {
		fmt.Printf("[UserHandler.getById] error when get path variable %v\n", err)
		utils.HandleError(writer, "Ops... something wrong", http.StatusBadRequest)
		return
	}

	data, err := h.userUseCase.GetById(id)
	if err != nil {
		fmt.Printf("[UserHandler.getById] error when get by id %v\n", err)
		utils.HandleError(writer, "Ops... something wrong", http.StatusBadRequest)
		return
	}
	utils.HandleSuccess(writer, http.StatusFound, data)

}

func (h UserHandler) getAllDataUser(writer http.ResponseWriter, request *http.Request) {
	data, err := h.userUseCase.GetAllDataUser()
	if err != nil {
		fmt.Printf("[UserHandler.getAllDataUser] got error when get all data user")
		utils.HandleError(writer, "Ops. something wrong", http.StatusBadRequest)
		return
	}
	utils.HandleSuccess(writer, http.StatusFound, data)
}

func (h UserHandler) deleteUserById(writer http.ResponseWriter, request *http.Request) {
	pathVariable := mux.Vars(request)
	id, err := strconv.Atoi(pathVariable["id"])
	if err != nil {
		fmt.Printf("[UserHandler.deteleUserByID] Got error when find path variable %v\n", err)
		utils.HandleError(writer, "Ops... something wrong", http.StatusBadRequest)
		return
	}

	_, err = h.userUseCase.GetById(id)
	if err != nil {
		fmt.Printf("[UserHandler.deleteById] Got Errror When check data with id %v\n", err)
		utils.HandleError(writer, "Ops.. Something wrong", http.StatusBadRequest)
		return
	}

	err = h.userUseCase.DeleteUserById(id)
	if err != nil {
		fmt.Printf("[UserHandler.deleteUserById] Got error when delete user by id %v\n", err)
		utils.HandleError(writer, "Ops... Something wrong", http.StatusBadRequest)
		return
	}
	utils.HandleSuccess(writer, http.StatusAccepted, nil)
}

func (h UserHandler) updateUserById(writer http.ResponseWriter, request *http.Request) {
	pathVariable := mux.Vars(request)
	id, err := strconv.Atoi(pathVariable["id"])
	if err != nil {
		fmt.Printf("[UserHandler.updateUserById] Got Error when find path variable %v\n", err)
		utils.HandleError(writer, "Ops... Something wrong", http.StatusBadRequest)
		return
	}

	data, err := h.userUseCase.GetById(id)
	if err != nil {
		fmt.Printf("[UserHandler.updateuserById] Got Error when find user by id %v\n", err)
		utils.HandleError(writer, "Ops.. Something wrong", http.StatusBadRequest)
		return
	}
	err2 := json.NewDecoder(request.Body).Decode(&data)
	if err2 != nil {
		fmt.Printf("[UserHandler.updateUserById] Got error when convert to json")
		utils.HandleError(writer, "Ops... Something wrong", http.StatusBadRequest)
		return
	}

	_, err = h.userUseCase.UpdateUserById(id, data)
	if err != nil {
		fmt.Printf("[UserHandler.updateUserById] Got Error when update user by id %w\n", err)
		utils.HandleError(writer, "Ops.. Something wrong", http.StatusBadRequest)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func CreateUserHandler(res *mux.Router, userUseCase user.UserUseCase) {
	userHandler := UserHandler{userUseCase}
	res.HandleFunc("/user", userHandler.register).Methods(http.MethodPost)
	res.HandleFunc("/user/{id}", userHandler.getById).Methods(http.MethodGet)
	res.HandleFunc("/users", userHandler.getAllDataUser).Methods(http.MethodGet)
	res.HandleFunc("/user/{id}", userHandler.deleteUserById).Methods(http.MethodDelete)
	res.HandleFunc("/user/{id}", userHandler.updateUserById).Methods(http.MethodPut)
}
