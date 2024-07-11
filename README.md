```markdown
# Linknau Test Project

Welcome to the Linknau Test - Simple project of implementing Struct, Interface, Jwt, and Unit Tests!

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Running Tests](#running-tests)
- [Endpoints](#endpoints)
- [License](#license)

## Project Structure

Here's a quick overview of our project layout:

```
linknau-test/
├── handlers/
│   └── handlers.go
├── middleware/
│   └── auth.go
├── models/
│   └── models.go
├── services/
│   └── fetch.go
├── tests/
│   └── handlers_test.go
├── main.go
└── go.mod
```

## Installation

Getting started is easy! Just follow these steps:

1. Clone the repository:

```sh
git clone https://github.com/aditwar-man/linknau-test.git
```

2. Navigate to the project directory:

```sh
cd linknau-test
```

3. Initialize Go modules and install dependencies:

```sh
go mod tidy
```

## Usage

Ready to run the application? Here's how:

1. Start the server:

```sh
go run main.go
```

2. The server will start on `http://localhost:8080`.

## Endpoints

### Login

- **URL:** `/login`
- **Method:** `POST`
- **Payload:**
  ```json
  {
    "username": "user",
    "password": "password"
  }
  ```

### Fetch Data

- **URL:** `/data`
- **Method:** `GET`
- **Headers:**
  - `Authorization: <JWT_TOKEN>`

## Running Tests

We have included tests to ensure everything works as expected. Run them with:

```sh
go test -v ./tests
```

The `-v` flag gives you a verbose output, so you can see detailed test results.

## Explanation of Code

### Structs

We have a simple `Data` struct defined in `models/models.go`:

```go
type Data struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```

### Interfaces

Here's an example of a simple interface and its implementation:

```go
type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}
```

### Package Management

Go's package management is handled by Go modules, which are defined in the `go.mod` file. To import a third-party package, use:

```sh
go get <package-path>
```

### Authentication

JWT authentication is implemented to keep your app secure. You can find the relevant code in `handlers/handlers.go` and `middleware/auth.go`. The `Login` endpoint generates a JWT token, and the `Authenticate` middleware verifies it.

### Testing

Unit tests for `FetchDataFromRemote` are provided in `tests/handlers_test.go`. The test checks the fetching and unmarshalling of JSON data from a mock server.

```go
func TestFetchData(t *testing.T) {
	mockData := []models.Data{
		{ID: 1, UserId: 101, Title: "Test1", Completed: false},
		{ID: 2, UserId: 102, Title: "Test2", Completed: true},
	}
	mockDataBytes, _ := json.Marshal(mockData)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockDataBytes)
	}))
	defer server.Close()

	url := server.URL
	data, err := services.FetchDataFromRemote(url)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Fetched Data: %+v", data)

	if len(data) != len(mockData) {
		t.Errorf("Expected %v elements, got %v", len(mockData), len(data))
	}

	for i := range data {
		if data[i].ID != mockData[i].ID || data[i].UserId != mockData[i].UserId || data[i].Title != mockData[i].Title || data[i].Completed != mockData[i].Completed {
			t.Errorf("Expected %v, got %v", mockData[i], data[i])
		}
	}
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

That's it! Enjoy exploring the project, and feel free to reach out if you have any questions or suggestions.
```