package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dejavxtrem/Greenlight_API/internal/data"
	// New import
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintln(w, "create a new movie")
	// Declare an anonymous struct to hold the information that we expect to be in the
	// HTTP request body (note that the field names and types in the struct are a subset
	// of the Movie struct that we created earlier). This struct will be our *target
	// decode destination*.
	var input struct {
		Title   string   `json:"title"`
		Year    int      `json:"year"`
		Runtime int      `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// Initialize a new json.Decoder instance which reads from the request body, and
	// then use the Decode() method to decode the body contents into the input struct.
	// Importantly, notice that when we call Decode() we pass a *pointer* to the input
	// struct as the target decode destination. If there was an error during decoding,
	// we use our generic errorResponse() helper to send a 400 Bad Request response
	// with the error message to the client.
	//err := json.NewDecoder(r.Body).Decode(&input)
	err := app.readJSON(w, r, &input)
	if err != nil {
		//app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

// Add a showMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// When httprouter is parsing a request, any interpolated URL parameters will be
	// stored in the request context. We can use the ParamsFromContext() function to
	// retrieve a slice containing these parameter names and values.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		//http.NotFound(w, r)
		return
	}

	// Create a new instance of the Movie struct, containing the ID we extracted from
	// the URL and some dummy data. Also notice that we deliberately haven't set a
	// value for the Year field.

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		//app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
	// We can then use the ByName() method to get the value of the "id" parameter from
	// the slice. In our project all movies will have a unique positive integer ID, but
	// the value returned by ByName() is always a string. So we try to convert it to an
	// integer. If the parameter couldn't be converted, or is less than 1, we know the
	// ID is invalid so we use the http.NotFound() function to return a 404 Not Found
	// response.
	// id, err := strconv.Atoi(params.ByName("id"))
	// if err != nil || id < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// Otherwise, interpolate the movie ID in a placeholder response.
	//fmt.Fprintf(w, "show the details of movie %d\n", id)
}
