///////////////////////////////////////////////////////////////////////
//			--------	API Testing Documentation	--------
//				GET /health → Should return "OK"
//				GET /users → List all users
//				GET /users?age=30 → Filter by age
//				GET /users?min_age=30 → Filter by age > 30
//				GET /users?max_age=30 → Filter by age <  30
//				GET /users?state=texas → Filter by state (case-insensitive)
//				POST /login
//						and add below configuration under the body -> raw
//						{
//							"username": "admin",
//							"password": "password"
//						}
//
//			--------	Application execution Documentation	--------
//				*** go run apiServer.go
//
//			--------	Application Build Documentation	--------
//				*** Windows: go build -o apiServer.exe
//				*** Linux: set GOOS=linux; set GOARCH=amd64; go build -o apiServer
//
///////////////////////////////////////////////////////////////////////


package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Define the User struct
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	State string `json:"state"`
}

// Sample data
var allUsers = []User{
	{"Alice Johnson", 28, "California"},
	{"Bob Smith", 34, "New York"},
	{"Charlie Brown", 22, "Texas"},
	{"Diana Prince", 30, "Washington"},
	{"Ethan Hunt", 40, "Nevada"},
	{"Fiona Glenanne", 27, "Florida"},
	{"George Miller", 35, "Ohio"},
	{"Hannah Davis", 31, "Illinois"},
	{"Ian Curtis", 29, "Michigan"},
	{"Julia Roberts", 33, "Georgia"},
	{"Carl Nickson", 28, "Tennessee"},
	{"Paul Wilfred", 35, "New York"},
	{"Bob Johnson", 50, "Iowa"},
}

func main() {
	fmt.Println("Server started on port 8080.")
	
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/userInfo", userInfoHandler)

	http.ListenAndServe(":8080", nil)
}

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received to check health.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Handles listing all users and filtering by age/state
func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received to access user info")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ageParam := r.URL.Query().Get("age")
	minAgeParam := r.URL.Query().Get("min_age")
	maxAgeParam := r.URL.Query().Get("max_age")
	stateParam := r.URL.Query().Get("state")

	filteredUsers := allUsers

	// Filter by age if provided
	if ageParam != "" {
		fmt.Printf("	>> filtering based on age...")
		age, err := strconv.Atoi(ageParam)
		if err == nil {
			var tmp []User
			for _, user := range filteredUsers {
				if user.Age == age {
					tmp = append(tmp, user)
				}
			}
			filteredUsers = tmp
		}
	}

	// Filter by minimum age
	if minAgeParam != "" {
		fmt.Printf("	>> filtering users with age > %s ...", minAgeParam)
		minAge, err := strconv.Atoi(minAgeParam)
		if err == nil {
			var tmp []User
			for _, user := range filteredUsers {
				if user.Age > minAge {
					tmp = append(tmp, user)
				}
			}
			filteredUsers = tmp
		}
	}

	// Filter by maximum age
	if maxAgeParam != "" {
		fmt.Printf("	>> filtering users with age < %s ...", maxAgeParam)
		maxAge, err := strconv.Atoi(maxAgeParam)
		if err == nil {
			var tmp []User
			for _, user := range filteredUsers {
				if user.Age < maxAge {
					tmp = append(tmp, user)
				}
			}
			filteredUsers = tmp
		}
	}

	// Filter by state if provided
	if stateParam != "" {
		fmt.Printf("	>> filtering users who live in %s ...", stateParam)
		state := strings.ToLower(stateParam)
		var tmp []User
		for _, user := range filteredUsers {
			if strings.ToLower(user.State) == state {
				tmp = append(tmp, user)
			}
		}
		filteredUsers = tmp
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredUsers)
}

// Handles login POST request
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received to login")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//Check for application/json content type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Simple auth check
	if creds.Username == "admin" && creds.Password == "password" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

// Returns the full list of user info
func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received to list all users")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}
