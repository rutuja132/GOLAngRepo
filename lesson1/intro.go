package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Define the template for the form
var tmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Form Example</title>
</head>
<body>
    <h1>User Form</h1>
    <form method="POST" action="/">
        <label for="name">Name:</label><br>
        <input type="text" id="name" name="name" required><br><br>
        <label for="email">Email:</label><br>
        <input type="email" id="email" name="email" required><br><br>
        <button type="submit">Submit</button>
    </form>
    {{if .}}
    <h2>Submitted Data:</h2>
    <p><strong>Name:</strong> {{.Name}}</p>
    <p><strong>Email:</strong> {{.Email}}</p>
    {{end}}
</body>
</html>
`

// User struct to store form data
type User struct {
	Name  string
	Email string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	t := template.Must(template.New("form").Parse(tmpl))

	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		// Retrieve form values
		user := User{
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
		}

		// Render the template with submitted data
		t.Execute(w, user)
		return
	}

	// Render the form without data for GET requests
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", formHandler)

	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}