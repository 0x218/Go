///////////////////////////////////////////////////////////////////////////////////////////
//			--------	API Testing Documentation	--------
//				GET /health → Should return "OK"
//				GET /users → List all users
//				GET /users?age=30 → Filter by age
//				GET /users?min_age=30 → Filter by age > 30
//				GET /users?max_age=30 → Filter by age <  30
//				GET /users?state=Tennessee → Filter by state (case-insensitive)
//				GET /users
//						and add below configuration under the params 'Key' and 'Value' (can be GET/POST)
//						min_age   30
//						state     Tennessee
//				POST /login
//						and add below configuration under the body -> raw
//						{
//							"username": "admin",
//							"password": "password"
//						}
//				NOTE: When using json body, request must send as POST.
//
//
//			--------	Application execution Documentation	--------
//				*** go run apiServer.go
//
//
//			--------	Application Build Documentation	--------
//				*** Windows: go build -o apiServer.exe
//				*** Linux: set GOOS=linux; set GOARCH=amd64; go build -o apiServer
//
///////////////////////////////////////////////////////////////////////////////////////////


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
type FilterParams struct {
	Age    *int   `json:"age,omitempty"`
	MinAge *int   `json:"min_age,omitempty"`
	MaxAge *int   `json:"max_age,omitempty"`
	State  string `json:"state,omitempty"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received to access user info")

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var ageParam, minAgeParam, maxAgeParam, stateParam string

	if r.Method == http.MethodGet {
		// GET params
		ageParam = r.URL.Query().Get("age")
		minAgeParam = r.URL.Query().Get("min_age")
		maxAgeParam = r.URL.Query().Get("max_age")
		stateParam = r.URL.Query().Get("state")
	} else if r.Method == http.MethodPost {
		// Check content type
		ct := r.Header.Get("Content-Type")
		if strings.Contains(ct, "application/json") {
			// JSON body
			var filters FilterParams
			if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			if filters.Age != nil {
				ageParam = strconv.Itoa(*filters.Age)
			}
			if filters.MinAge != nil {
				minAgeParam = strconv.Itoa(*filters.MinAge)
			}
			if filters.MaxAge != nil {
				maxAgeParam = strconv.Itoa(*filters.MaxAge)
			}
			stateParam = filters.State
		} else {
			// Form data
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}
			ageParam = r.Form.Get("age")
			minAgeParam = r.Form.Get("min_age")
			maxAgeParam = r.Form.Get("max_age")
			stateParam = r.Form.Get("state")
		}
	}

	// === Filtering logic (same as before) ===
	filteredUsers := allUsers

	if ageParam != "" {
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

	if minAgeParam != "" {
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

	if maxAgeParam != "" {
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

	if stateParam != "" {
		state := strings.ToLower(stateParam)
		var tmp []User
		for _, user := range filteredUsers {
			if strings.ToLower(user.State) == state {
				tmp = append(tmp, user)
			}
		}
		filteredUsers = tmp
	}

	// Return JSON
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
