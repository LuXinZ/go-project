package main

import (
	"fmt"
	"net/http"
	"net/internal/data"
	"time"
)

// post /v1/movies
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
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
