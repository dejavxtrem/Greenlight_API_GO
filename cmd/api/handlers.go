package main

import (
	//"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dejavxtrem/Greenlight_API/internal/data"
	"github.com/dejavxtrem/Greenlight_API/internal/validator"
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
		Title   string       `json:"title"`
		Year    int          `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		//app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Movie struct.
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	// Initialize a new json.Decoder instance which reads from the request body, and
	// then use the Decode() method to decode the body contents into the input struct.
	// Importantly, notice that when we call Decode() we pass a *pointer* to the input
	// struct as the target decode destination. If there was an error during decoding,
	// we use our generic errorResponse() helper to send a 400 Bad Request response
	// with the error message to the client.
	//err := json.NewDecoder(r.Body).Decode(&input)

	// Initialize a new Validator instance.
	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// When sending an HTTP response, we want to include a Location header to let the
	// client know which URL they can find the newly-created resource at. We make an
	// empty http.Header map and then use the Set() method to add a new Location header,
	// interpolating the system-generated ID for our new movie in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	// Write a JSON response with a 201 Created status code, the movie data in the
	// response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"movies": movie}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// v.Check(input.Title != "", "title", "must be provided")
	// v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")
	// v.Check(input.Year != 0, "year", "must be provided")
	// v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	// v.Check(input.Year <= time.Now().Year(), "year", "must not be in the future")

	// v.Check(input.Runtime != 0, "runtime", "must be provided")
	// v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

	// v.Check(input.Genres != nil, "genres", "must be provided")
	// v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
	// v.Check(len(input.Genres) <= 5, "genres", "must not contain more than 5 genres")

	// // Note that we're using the Unique helper in the line below to check that all
	// // values in the input.Genres slice are unique.
	// v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate values")

	// Use the Valid() method to see if any of the checks failed. If they did, then use
	// the failedValidationResponse() helper to send a response to the client, passing
	// in the v.Errors map.
	// if !v.Valid() {
	// 	app.failedValidationResponse(w, r, v.Errors)
	// 	return
	// }

	//fmt.Fprintf(w, "%+v\n", input)
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

	// Call the Get() method to fetch the data for a specific movie. We also need to
	// use the errors.Is() function to check if it returns a data.ErrRecordNotFound
	// error, in which case we send a 404 Not Found response to the client.
	movie, err := app.models.Movies.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)

		default:
			app.serverErrorResponse(w, r, err)

		}
		return
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
