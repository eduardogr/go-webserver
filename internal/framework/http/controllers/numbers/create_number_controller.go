package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewCreateNumberController(usecase usecases.CreateNewNumberUsecase) *CreateNumberController {
	return &CreateNumberController{
		CreateNewNumberUsecase: usecase,
	}
}

type CreateNumberController struct {
	CreateNewNumberUsecase usecases.CreateNewNumberUsecase
}

func (c *CreateNumberController) Execute(w http.ResponseWriter, r *http.Request) {
	// TODO: check that number does NOT exist

	// get the body of our POST request
	// unmarshal this into a new NewNumberRequest struct

	reqBody, _ := io.ReadAll(r.Body)

	var request domain.NewNumberRequest
	err := json.Unmarshal(reqBody, &request)
	if err != nil {
		response := map[string]interface{}{
			"success": false,
			"data": map[string]string{
				"error": err.Error(),
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// create request
	err = c.CreateNewNumberUsecase.Execute(request)

	if err != nil {
		response := map[string]interface{}{
			"success": false,
			"data": map[string]string{
				"error": err.Error(),
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"number": request.ID,
		},
	}
	json.NewEncoder(w).Encode(response)
}
