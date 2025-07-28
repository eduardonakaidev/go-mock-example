# Go Mock Example with Uber's mockgen

This repository demonstrates a practical implementation of mocking in Go using Uber's **`mockgen`** library. It shows how to isolate dependencies in unit tests for faster, more reliable testing.

## **Project Structure**

go-mock-example/
├── main.go
├── repo/
│   └── user_repository.go
├── service/
│   ├── user_service.go
│   └── user_service_test.go
├── mocks/
│   └── mock_user_repository.go
└── scripts/
└── [test.sh](http://test.sh/)

## **Getting Started**

### **Prerequisites**

- Go 1.24+
- **`mockgen`** tool

### **Installation**

```bash
# Install mockgen
go install go.uber.org/mock/mockgen@latest

# Clone the repository
git clone https://github.com/eduardonakaidev/go-mock-example.git
cd go-mock-example
```

**Generate Mocks**

```bash
mockgen -source=repo/user_repository.go -destination=mocks/mock_user_repository.go -package=mocks
```

**Run Tests**

```bash
# Using Go's built-in test runner
go test ./...

# Using custom test script
chmod +x scripts/test.sh
./scripts/test.sh
```

**Run the Application**

```bash
go run main.go
```

## **Code Overview**

### **Repository Interface (`repo/user_repository.go`)**

```go
package repo

type User struct {
    ID   int
    Name string
}

type UserRepository interface {
    GetUser(id int) (*User, error)
}

type realUserRepository struct{}

func NewRealUserRepository() UserRepository {
    return &realUserRepository{}
}

func (r *realUserRepository) GetUser(id int) (*User, error) {
    return &User{ID: id, Name: "Real User"}, nil
}
```

**Service Implementation (`service/user_service.go`)**

```go
package service

import "github.com/eduardonakaidev/go-mock-example/repo"

type UserService struct {
    repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUserName(id int) (string, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        return "", err
    }
    return user.Name, nil
}
```

**Test with Mock (`service/user_service_test.go`)**

```go
package service_test

import (
    "testing"
    "github.com/eduardonakaidev/go-mock-example/mocks"
    "github.com/eduardonakaidev/go-mock-example/repo"
    "github.com/eduardonakaidev/go-mock-example/service"
    "go.uber.org/mock/gomock"
)

func TestGetUserName(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockUserRepository(ctrl)
    mockRepo.EXPECT().
        GetUser(123).
        Return(&repo.User{ID: 123, Name: "John Mock"}, nil)

    userService := service.NewUserService(mockRepo)
    name, err := userService.GetUserName(123)
    
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    
    if name != "John Mock" {
        t.Errorf("Expected: John Mock, Got: %s", name)
    }
}
```

**Test Script (`scripts/test.sh`)**

```bash
#!/bin/sh

# Find and run all test files
for OUTPUT in $(find . -name '*_test.go')
do
    echo "Running tests in: $OUTPUT"
    RESULT=$(go test $OUTPUT)
    
    if echo "$RESULT" | grep "FAIL"; then
        echo "\n❌ Test failed:"
        echo "$RESULT"
        exit 1
    else 
        echo "✅ Tests passed\n"
    fi
done

echo "✨ All tests passed successfully!"
```

## **Key Benefits of Using `mockgen`**

### **1. Isolated Testing**

- Test components without external dependencies
- Simulate database/API responses
- Create error scenarios on demand
- Eliminate flaky tests caused by external services

### **2. Type-Safe Mocks**

- Auto-generated from interfaces
- Compile-time safety
- Always matches interface definitions
- Automatic synchronization with interface changes

### **3. Behavior Verification**

- Ensure methods are called with correct parameters
- Verify call sequences
- Validate expected interactions
- Set call count expectations

### **4. Performance**

- Tests run 10-100x faster
- No network/database overhead
- Parallel test execution
- Reduced test execution time in CI/CD pipelines

### **5. Reliability**

- Consistent test results
- Reproducible error states
- Deterministic test outcomes
- Reduced false positives/negatives

## **Comparison: Manual Mocks vs `mockgen`**

| **Feature** | **Manual Mocks** | **`mockgen`** |
| --- | --- | --- |
| Type safety | ❌ Error-prone | ✅ Guaranteed |
| Sync with interfaces | Manual | Automatic |
| Complex scenarios | Difficult | Easy |
| Call verification | Limited | Advanced |
| Boilerplate code | High | None |
| Maintenance effort | High | Low |
| Dynamic return values | Limited | Full |

## **When to Use Mocking**

1. **Database Operations**
    - Test business logic without real database
    - Simulate various database states
    - Validate transaction handling
2. **External APIs**
    - Test API integration points
    - Simulate different API responses
    - Validate error handling
3. **Complex Systems**
    - Isolate components in microservices
    - Test edge cases and error scenarios
    - Validate distributed system behavior
4. **Performance-Critical Tests**
    - Speed up test suites
    - Run tests in parallel
    - Reduce resource consumption