package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// Create a fixed-format JSON response from a string. Notice how we're using a raw
	// string literal (enclosed with backticks) so that we can include double quote
	// characters in the JSON without needing to escape them? We also use the %q verb to
	// wrap the interpolated values in double quotes.
	//js := `{"status": "available", "environment": %q, "version": %q}`
	//js = fmt.Sprintf(js, app.config.env, version)
	data := map[string]string{
		"status":      "avalable",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	// Pass the map to the json.Marshal() function. This returns a []byte slice
	// containing the encoded JSON. If there was an error, we log it and send the client
	// a generic error message.
	//js, err := json.Marshal(data)

	// Append a newline to the JSON. This is just a small nicety to make it easier to
	// view in terminal applications.
	//js = append(js, '\n')

	// Set the "Content-Type: application/json" header on the response. If you forget to do
	// this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.
	//w.Header().Set("Content-Type", "application/json")

	// Use w.Write() to send the []byte slice containing the JSON as the response body.
	//w.Write(js)

	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
}
