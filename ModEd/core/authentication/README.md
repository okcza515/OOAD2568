# Authentication Package

## Overview

The `authentication` package provides a complete authentication system for Go applications. It includes user management, authentication, authorization, and a CLI interface for user operations. The package is designed to be flexible and can be integrated with various database backends through GORM.

## Components

### 1. Authentication Provider

The authentication provider is responsible for user management and authentication operations. The package includes a GORM-based implementation (`DBAuthProvider`), but you can create custom providers by implementing the `AuthenticationProvider` interface.

### 2. Middleware

The middleware provides a simplified interface for authentication operations and handles authorization checks. It acts as a wrapper around the authentication provider.

### 3. User Context

The `UserContext` structure holds authenticated user information and is stored in the context.Context for access throughout the application.

### 4. CLI Menu

The package includes a CLI menu for authentication operations like login, registration, password changes, and account deletion.

## Installation

```go
import "yourmodule/core/authentication"
```

## Usage

### Initializing the Authentication System

```go
import (
    "time"
    "context"
    "gorm.io/gorm"
    "yourmodule/core/authentication"
)

// Initialize database connection
db, err := gorm.Open(...)
if err != nil {
    // Handle error
}

// Create authentication provider
// The token expiry age is set to 24 hours
authProvider := authentication.NewDBAuthProvider(db, 24*time.Hour)

// Create middleware
authMiddleware := authentication.NewMiddleware(authProvider)
```

### Role-Based Access Control

The package supports configurable role-based access control. You can specify which roles are allowed to access your application:

```go
// Create authentication CLI
authCLI := authentication.NewAuthenticationCLI()
authCLI.SetDB(db)

// Configure allowed roles
// Only admin users can access the application
authCLI.SetAllowedRoles([]string{"admin"})

// Or allow multiple roles
authCLI.SetAllowedRoles([]string{"admin", "user"})
```

When a user tries to log in:
- If their role is in the allowed roles list, they can access the application
- If their role is not allowed, they will receive an "Access denied" message
- The program will stay at the authentication menu, allowing them to try again or exit

### User Authentication

```go
// Authenticate a user
userCtx, err := authMiddleware.Authenticate(context.Background(), "username", "password")
if err != nil {
    // Handle authentication error
}

// Create a new context with user information
ctx := authentication.WithContext(context.Background(), userCtx)

// Get the current user from context
currentUser, err := authMiddleware.GetCurrentUser(ctx)
if err != nil {
    // Handle error
}
```

### User Management

```go
// Create a new user
err := authMiddleware.CreateUser(ctx, "newuser", "password", "user")
if err != nil {
    // Handle error
}

// Update password
err = authMiddleware.UpdatePassword(ctx, "username", "oldpassword", "newpassword")
if err != nil {
    // Handle error
}

// Delete user
err = authMiddleware.DeleteUser(ctx, "username")
if err != nil {
    // Handle error
}

// List all users (requires admin privileges)
users, err := authMiddleware.ListUsers(ctx)
if err != nil {
    // Handle error
}

// Update user role
err = authMiddleware.UpdateUserRole(ctx, "username", "admin")
if err != nil {
    // Handle error
}
```

### Role-Based Authorization

```go
// Check if user has admin privileges
err := authentication.RequireAdmin(ctx)
if err != nil {
    // User is not an admin
}
```

### Using the CLI Menu

```go
// Create a new authentication menu state
authMenuState := authentication.NewAuthMenuState(db)

// Run the menu loop
for {
    // Display the menu
    authMenuState.Render()
    
    // Get user input
    var input string
    fmt.Scanln(&input)
    
    // Handle user input
    err := authMenuState.HandleUserInput(input)
    if err != nil {
        if err.Error() == "exit" {
            break
        } else if err.Error() == "login_success" {
            // User has logged in successfully, you can perform post-login operations
            ctx := authMenuState.GetContext() // Get the context with user information
        } else {
            fmt.Println(err)
        }
    }
}
```

## Error Handling

The package defines several error types:

- `ErrInvalidAuthentications`: Invalid username or password
- `ErrUserNotFound`: User not found
- `ErrUserExists`: User already exists
- `ErrUnauthorized`: User does not have required privileges
- `ErrRoleNotAllowed`: User's role is not allowed to access the program

## Security Notes

1. Passwords are hashed using bcrypt before storage
2. User sessions expire after the configured token expiry age
3. Role-based access control is implemented through the `RequireAdmin` function and configurable role restrictions
4. User deletion and role updates should be restricted to administrators

## Implementation Details

- The `DBAuthProvider` uses GORM to interact with the database
- User records are stored in the `DBUser` table with bcrypt-hashed passwords
- The authentication menu provides a simple CLI for common user operations
- Role-based access control can be configured per application instance 