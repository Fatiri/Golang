package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"users/model"
)

func HandleSuccess(writer http.ResponseWriter, status int, data interface{}) {
	result := model.Response{
		Status:  true,
		Message: "success",
		Data:    data,
	}

	writer.Header().Set("Content-type", "aplication/json")
	writer.WriteHeader(status)
	error := json.NewEncoder(writer).Encode(result)
	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Some error has been show"))
		fmt.Printf("[utils.HandleSucces] error when convert to json %v\n", error)
		return
	}
}

func HandleError(writer http.ResponseWriter, message string, status int) {
	result := model.Response{
		Status:  false,
		Message: message,
		Data:    nil,
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(status)
	error := json.NewEncoder(writer).Encode(result)
	if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		//Write hanya mengenal type byte sehingga result error yg berupa string di cpnvert dahulu ke byte
		writer.Write([]byte("some error has been show"))
		fmt.Printf("[utils.HandleSucces] error when convert to json %v\n", error)
		return
	}

}

var datas = [][]string{{"jaka1", "jaka2"}, {"jajang1", "jajang2"}}

func PrepareDataAndReturnCsv() {

	file, err := os.Create("/resultttttt.csv")
	if err != nil {
		fmt.Printf("[main] Got error when create file csf")
	}

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, value := range datas {
		err := writer.Write(value)
		if err != nil {
			fmt.Println("Ops.. something wrong")
		}
	}
}
