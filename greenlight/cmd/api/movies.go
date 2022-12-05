package main

import (
	"fmt"
	"net/http"
	"net/internal/data"
	"net/internal/validator"
	"time"
)

// post /v1/movies
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	movice := &data.Movie{
		CreatedAt: time.Time{},
		Title:     input.Title,
		Year:      input.Year,
		Runtime:   input.Runtime,
		Genres:    input.Genres,
	}
	v := validator.New()

	if data.ValidateMovie(v, movice); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// movie struct
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "casablanca",
		Year:      2022,
		Runtime:   102,
		Genres:    []string{"romance", "war", "drama"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
