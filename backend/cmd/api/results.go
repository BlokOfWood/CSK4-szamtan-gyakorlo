package main

import (
	"fmt"
	"net/http"

	"github.com/BlokOfWood/CSK4-szamtan-gyakorlo/backend/internal/data"
)

func (app *application) createResultHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	result := &data.Result{
		Name:  input.Name,
		Score: input.Score,
	}

	err = app.models.Results.Insert(result)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/result/%d", result.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"result": result}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showResultHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	result, err := app.models.Results.Get(id)
	if err != nil {
		switch {
		case err.Error() == data.ErrRecordNotFound.Error():
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"result": result}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listResultsHandler(w http.ResponseWriter, r *http.Request) {
	results, err := app.models.Results.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"results": results}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
