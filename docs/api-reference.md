# API Reference Documentation

## 📚 API Overview

This document provides comprehensive reference for all APIs in the Go web application template. The API follows RESTful principles and provides standardized responses.

## 🔗 Base URLs

- **Admin Server**: `http://localhost:8000`
- **User Server**: `http://localhost:8001`
- **API Documentation**: `http://localhost:8000/swagger/`

## 📋 Response Format

### Success Response

```json
{
  "success": true,
  "data": {
    // Response data
  },
  "message": "Operation completed successfully"
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error description",
  "code": 400
}
```

### Common HTTP Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `422` - Validation Error
- `500` - Internal Server Error

## 🔐 Authentication

### JWT Token Format

```
Authorization: Bearer <jwt_token>
```

### Token Structure

```json
{
  "user_id": 123,
  "exp": 1640995200,
  "iat": 1640908800
}
```

## 📖 API Endpoints

### Common APIs

#### Health Check

```http
GET /api/health
```

**Response:**

```json
{
  "success": true,
  "data": {
    "status": "healthy",
    "timestamp": "2024-01-01T00:00:00Z",
    "version": "1.0.0"
  },
  "message": "Service is healthy"
}
```

#### Authentication

##### Login

```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  },
  "message": "Login successful"
}
```

**Validation Rules:**

- `email`: Required, valid email format
- `password`: Required, minimum 1 character, maximum 255 characters

##### Register

```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "newuser@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "message": "User registered successfully"
  },
  "message": "Registration successful"
}
```

**Validation Rules:**

- `email`: Required, valid email format, unique
- `password`: Required, minimum 1 character, maximum 255 characters

#### Internationalization

##### Get Translations

```http
GET /api/i18n/translations
Accept-Language: en
```

**Response:**

```json
{
  "success": true,
  "data": {
    "auth": {
      "login": "Login",
      "register": "Register",
      "logout": "Logout"
    },
    "common": {
      "save": "Save",
      "cancel": "Cancel",
      "delete": "Delete"
    }
  },
  "message": "Translations loaded"
}
```

##### Get Supported Locales

```http
GET /api/i18n/locales
```

**Response:**

```json
{
  "success": true,
  "data": {
    "default_locale": "en",
    "supported_locales": ["en", "vi"]
  },
  "message": "Supported locales retrieved"
}
```

#### Profile

##### Get Current User Profile

```http
GET /api/profile
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 123,
    "email": "user@example.com",
    "full_name": "John Doe",
    "is_active": true,
    "is_admin": false,
    "roles": [
      {
        "id": 1,
        "code": "user",
        "name": "User",
        "description": "Regular user role"
      }
    ],
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "Profile retrieved successfully"
}
```

##### Update Profile

```http
PUT /api/profile
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "full_name": "John Smith",
  "email": "john.smith@example.com"
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 123,
    "email": "john.smith@example.com",
    "full_name": "John Smith",
    "is_active": true,
    "is_admin": false,
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "Profile updated successfully"
}
```

### Admin APIs

#### User Management

##### List Users (DataTable)

```http
POST /api/admin/users/datatable
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "draw": 1,
  "start": 0,
  "length": 10,
  "search": {
    "value": "john",
    "regex": false
  },
  "order": [
    {
      "column": 0,
      "dir": "asc"
    }
  ],
  "columns": [
    {
      "data": "id",
      "name": "",
      "searchable": true,
      "orderable": true,
      "search": {
        "value": "",
        "regex": false
      }
    }
  ]
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "draw": 1,
    "recordsTotal": 100,
    "recordsFiltered": 25,
    "data": [
      {
        "id": 123,
        "email": "user@example.com",
        "full_name": "John Doe",
        "is_active": true,
        "is_admin": false,
        "roles": [
          {
            "id": 1,
            "code": "user",
            "name": "User"
          }
        ],
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  },
  "message": "Users retrieved successfully"
}
```

##### Get User by ID

```http
GET /api/admin/users/{id}
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 123,
    "email": "user@example.com",
    "full_name": "John Doe",
    "is_active": true,
    "is_admin": false,
    "roles": [
      {
        "id": 1,
        "code": "user",
        "name": "User",
        "description": "Regular user role"
      }
    ],
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "User retrieved successfully"
}
```

##### Create User

```http
POST /api/admin/users
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "email": "newuser@example.com",
  "password": "password123",
  "full_name": "New User",
  "is_active": true,
  "is_admin": false,
  "role_ids": [1, 2]
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 124,
    "email": "newuser@example.com",
    "full_name": "New User",
    "is_active": true,
    "is_admin": false,
    "created_at": "2024-01-01T00:00:00Z"
  },
  "message": "User created successfully"
}
```

**Validation Rules:**

- `email`: Required, valid email format, unique
- `password`: Required, minimum 8 characters
- `full_name`: Optional, maximum 255 characters
- `is_active`: Boolean, defaults to true
- `is_admin`: Boolean, defaults to false
- `role_ids`: Array of valid role IDs

##### Update User

```http
PUT /api/admin/users/{id}
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "email": "updated@example.com",
  "full_name": "Updated User",
  "is_active": true,
  "is_admin": false,
  "role_ids": [1, 3]
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 123,
    "email": "updated@example.com",
    "full_name": "Updated User",
    "is_active": true,
    "is_admin": false,
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "User updated successfully"
}
```

##### Delete User

```http
DELETE /api/admin/users/{id}
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "success": true,
  "data": null,
  "message": "User deleted successfully"
}
```

#### Role Management

##### List Roles

```http
GET /api/admin/roles
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "code": "admin",
      "name": "Administrator",
      "description": "Full system access",
      "permissions": [
        {
          "id": 1,
          "code": "user:create",
          "name": "Create Users",
          "service": "user",
          "method": "create"
        }
      ],
      "created_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "Roles retrieved successfully"
}
```

##### Create Role

```http
POST /api/admin/roles
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "code": "moderator",
  "name": "Moderator",
  "description": "Content moderation role",
  "permission_ids": [1, 2, 3]
}
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 2,
    "code": "moderator",
    "name": "Moderator",
    "description": "Content moderation role",
    "created_at": "2024-01-01T00:00:00Z"
  },
  "message": "Role created successfully"
}
```

#### Permission Management

##### List Permissions

```http
GET /api/admin/permissions
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "code": "user:create",
      "name": "Create Users",
      "description": "Permission to create new users",
      "service": "user",
      "method": "create",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "Permissions retrieved successfully"
}
```

### User APIs

#### User Dashboard

##### Get Dashboard Data

```http
GET /api/user/dashboard
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "success": true,
  "data": {
    "user": {
      "id": 123,
      "email": "user@example.com",
      "full_name": "John Doe"
    },
    "stats": {
      "total_items": 150,
      "recent_activity": 25
    },
    "recent_activities": [
      {
        "id": 1,
        "action": "login",
        "timestamp": "2024-01-01T00:00:00Z"
      }
    ]
  },
  "message": "Dashboard data retrieved"
}
```

## 🔍 DataTable API

### DataTable Request Format

```json
{
  "draw": 1,
  "start": 0,
  "length": 10,
  "search": {
    "value": "search term",
    "regex": false
  },
  "order": [
    {
      "column": 0,
      "dir": "asc"
    }
  ],
  "columns": [
    {
      "data": "column_name",
      "name": "",
      "searchable": true,
      "orderable": true,
      "search": {
        "value": "",
        "regex": false
      }
    }
  ],
  "action": "ajax"
}
```

### DataTable Response Format

```json
{
  "success": true,
  "data": {
    "draw": 1,
    "recordsTotal": 1000,
    "recordsFiltered": 250,
    "data": [
      {
        "id": 1,
        "name": "Example Item",
        "created_at": "2024-01-01T00:00:00Z"
      }
    ]
  },
  "message": "Data retrieved successfully"
}
```

### Export Actions

#### Excel Export

```http
POST /api/admin/users/datatable
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "action": "excel",
  "columns": ["id", "email", "full_name", "created_at"]
}
```

#### CSV Export

```http
POST /api/admin/users/datatable
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "action": "csv",
  "columns": ["id", "email", "full_name", "created_at"]
}
```

#### PDF Export

```http
POST /api/admin/users/datatable
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "action": "pdf",
  "columns": ["id", "email", "full_name", "created_at"]
}
```

## 🛡️ Error Handling

### Validation Errors

```json
{
  "success": false,
  "message": "Validation failed",
  "code": 422,
  "errors": {
    "email": ["Email is required", "Invalid email format"],
    "password": ["Password must be at least 8 characters"]
  }
}
```

### Authentication Errors

```json
{
  "success": false,
  "message": "Invalid or expired token",
  "code": 401
}
```

### Authorization Errors

```json
{
  "success": false,
  "message": "Insufficient permissions",
  "code": 403
}
```

### Server Errors

```json
{
  "success": false,
  "message": "Internal server error",
  "code": 500
}
```

## 📊 Rate Limiting

### Rate Limit Headers

```
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 59
X-RateLimit-Reset: 1640995200
```

### Rate Limit Response

```json
{
  "success": false,
  "message": "Too many requests. Please try again later.",
  "code": 429
}
```

## 🔧 API Versioning

The API supports versioning through URL paths:

- Current version: `/api/`
- Future versions: `/api/v2/`, `/api/v3/`

## 📝 OpenAPI/Swagger

Interactive API documentation is available at:

- **Swagger UI**: `http://localhost:8000/swagger/`
- **OpenAPI JSON**: `http://localhost:8000/swagger/doc.json`
- **OpenAPI YAML**: `http://localhost:8000/swagger/doc.yaml`

## 🧪 Testing APIs

### Health Check Test

```bash
curl -X GET http://localhost:8000/api/health
```

### Authentication Test

```bash
# Login
curl -X POST http://localhost:8000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'

# Use token
curl -X GET http://localhost:8000/api/profile \
  -H "Authorization: Bearer <token>"
```

### DataTable Test

```bash
curl -X POST http://localhost:8000/api/admin/users/datatable \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"draw":1,"start":0,"length":10}'
```

This API reference provides comprehensive documentation for all endpoints, request/response formats, and usage examples.
