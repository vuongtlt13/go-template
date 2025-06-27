# Database Schema Documentation

## 🗄️ Database Overview

This document describes the database schema for the Go web application template. The application uses **PostgreSQL** as the primary database with **GORM** as the ORM layer.

## 📊 Database Configuration

### Connection Settings

```yaml
database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "postgres"
  dbname: "go_template"
  sslmode: "disable"
```

### Connection Pooling

```go
sqlDB.SetMaxIdleConns(50)
sqlDB.SetMaxOpenConns(200)
sqlDB.SetConnMaxLifetime(30 * time.Minute)
```

## 🏗️ Schema Design

### Core Tables

#### 1. Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    is_admin BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
```

**GORM Model:**

```go
type User struct {
    ID        uint64         `gorm:"primarykey" json:"id"`
    Email     string         `gorm:"size:255;not null;uniqueIndex" json:"email"`
    Password  string         `gorm:"size:255;not null" json:"-"`
    FullName  string         `gorm:"size:255" json:"full_name"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
    Roles     []Role         `gorm:"many2many:user_roles;" json:"roles,omitempty"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Indexes:**

- Primary Key: `id`
- Unique Index: `email`
- Soft Delete Index: `deleted_at`

#### 2. Roles Table

```sql
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
```

**GORM Model:**

```go
type Role struct {
    ID          uint64         `gorm:"primarykey" json:"id"`
    Code        string         `gorm:"size:50;not null;uniqueIndex" json:"code"`
    Name        string         `gorm:"size:100;not null" json:"name"`
    Description string         `gorm:"size:255" json:"description"`
    Users       []User         `gorm:"many2many:user_roles;" json:"users,omitempty"`
    Permissions []Permission   `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Indexes:**

- Primary Key: `id`
- Unique Index: `code`

#### 3. Permissions Table

```sql
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    code VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    service VARCHAR(50) NOT NULL,
    method VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
```

**GORM Model:**

```go
type Permission struct {
    ID          uint64         `gorm:"primarykey" json:"id"`
    Code        string         `gorm:"size:100;not null;uniqueIndex" json:"code"`
    Name        string         `gorm:"size:100;not null" json:"name"`
    Description string         `gorm:"size:255" json:"description"`
    Service     string         `gorm:"size:50;not null" json:"service"`
    Method      string         `gorm:"size:50;not null" json:"method"`
    Roles       []Role         `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Indexes:**

- Primary Key: `id`
- Unique Index: `code`
- Composite Index: `service, method`

### Junction Tables

#### 4. User Roles Table

```sql
CREATE TABLE user_roles (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id)
);
```

**Indexes:**

- Primary Key: `(user_id, role_id)`
- Foreign Key: `user_id` → `users(id)`
- Foreign Key: `role_id` → `roles(id)`

#### 5. Role Permissions Table

```sql
CREATE TABLE role_permissions (
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id, permission_id)
);
```

**Indexes:**

- Primary Key: `(role_id, permission_id)`
- Foreign Key: `role_id` → `roles(id)`
- Foreign Key: `permission_id` → `permissions(id)`

### Audit Tables

#### 6. User Logs Table

```sql
CREATE TABLE user_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    action VARCHAR(100) NOT NULL,
    description TEXT,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

**GORM Model:**

```go
type UserLog struct {
    ID          uint64    `gorm:"primarykey" json:"id"`
    UserID      *uint64   `gorm:"index" json:"user_id"`
    Action      string    `gorm:"size:100;not null" json:"action"`
    Description string    `gorm:"type:text" json:"description"`
    IPAddress   string    `gorm:"type:inet" json:"ip_address"`
    UserAgent   string    `gorm:"type:text" json:"user_agent"`
    CreatedAt   time.Time `json:"created_at"`
    User        User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
```

**Indexes:**

- Primary Key: `id`
- Index: `user_id`
- Index: `created_at`

## 🔗 Relationships

### Entity Relationship Diagram

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│    Users    │    │    Roles    │    │Permissions  │
├─────────────┤    ├─────────────┤    ├─────────────┤
│ id (PK)     │    │ id (PK)     │    │ id (PK)     │
│ email       │    │ code        │    │ code        │
│ password    │    │ name        │    │ name        │
│ full_name   │    │ description │    │ service     │
│ is_active   │    │ created_at  │    │ method      │
│ is_admin    │    │ updated_at  │    │ created_at  │
│ created_at  │    │ deleted_at  │    │ updated_at  │
│ updated_at  │    └─────────────┘    │ deleted_at  │
│ deleted_at  │                       └─────────────┘
└─────────────┘                                │
       │                                       │
       │                                       │
       └─────────────┐    ┌────────────────────┘
                     │    │
              ┌─────────────┐
              │ User Roles  │
              ├─────────────┤
              │ user_id (FK)│
              │ role_id (FK)│
              │ created_at  │
              └─────────────┘
                     │
                     │
              ┌─────────────┐
              │Role Perms   │
              ├─────────────┤
              │ role_id (FK)│
              │ perm_id (FK)│
              │ created_at  │
              └─────────────┘
```

### Relationship Definitions

#### One-to-Many Relationships

- **User → UserLogs**: One user can have many log entries

#### Many-to-Many Relationships

- **Users ↔ Roles**: Through `user_roles` table
- **Roles ↔ Permissions**: Through `role_permissions` table

## 📊 Data Seeding

### Default Roles

```sql
INSERT INTO roles (code, name, description) VALUES
('admin', 'Administrator', 'Full system access'),
('user', 'User', 'Regular user access'),
('moderator', 'Moderator', 'Content moderation access');
```

### Default Permissions

```sql
INSERT INTO permissions (code, name, description, service, method) VALUES
('user:create', 'Create Users', 'Permission to create new users', 'user', 'create'),
('user:read', 'View Users', 'Permission to view user information', 'user', 'read'),
('user:update', 'Update Users', 'Permission to update user information', 'user', 'update'),
('user:delete', 'Delete Users', 'Permission to delete users', 'user', 'delete'),
('role:create', 'Create Roles', 'Permission to create new roles', 'role', 'create'),
('role:read', 'View Roles', 'Permission to view role information', 'role', 'read'),
('role:update', 'Update Roles', 'Permission to update role information', 'role', 'update'),
('role:delete', 'Delete Roles', 'Permission to delete roles', 'role', 'delete'),
('permission:create', 'Create Permissions', 'Permission to create new permissions', 'permission', 'create'),
('permission:read', 'View Permissions', 'Permission to view permission information', 'permission', 'read'),
('permission:update', 'Update Permissions', 'Permission to update permission information', 'permission', 'update'),
('permission:delete', 'Delete Permissions', 'Permission to delete permissions', 'permission', 'delete');
```

### Default Admin User

```sql
-- Password: admin123 (bcrypt hashed)
INSERT INTO users (email, password, full_name, is_active, is_admin) VALUES
('admin@example.com', '$2a$10$...', 'System Administrator', true, true);

-- Assign admin role to admin user
INSERT INTO user_roles (user_id, role_id) VALUES
(1, (SELECT id FROM roles WHERE code = 'admin'));

-- Assign all permissions to admin role
INSERT INTO role_permissions (role_id, permission_id)
SELECT
    (SELECT id FROM roles WHERE code = 'admin'),
    id
FROM permissions;
```

## 🔄 Migration System

### Migration Files Structure

```
migrations/
├── 000001_create_init_tables.up.sql
├── 000001_create_init_tables.down.sql
├── 000002_add_user_profile.up.sql
├── 000002_add_user_profile.down.sql
└── ...
```

### Migration Commands

```bash
# Run migrations up
make migrate-up

# Run migrations down
make migrate-down

# Create new migration
migrate create -ext sql -dir migrations -seq add_new_table
```

### Example Migration

```sql
-- Up migration
-- 000001_create_init_tables.up.sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    is_admin BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Down migration
-- 000001_create_init_tables.down.sql
DROP TABLE users;
```

## 🔍 Query Optimization

### Indexing Strategy

#### Primary Indexes

- All tables have `id` as primary key
- Unique constraints on business keys (`email`, `code`)

#### Performance Indexes

```sql
-- User queries by email
CREATE INDEX idx_users_email ON users(email);

-- User queries by active status
CREATE INDEX idx_users_is_active ON users(is_active);

-- Log queries by user and date
CREATE INDEX idx_user_logs_user_created ON user_logs(user_id, created_at);

-- Permission queries by service and method
CREATE INDEX idx_permissions_service_method ON permissions(service, method);
```

#### Composite Indexes

```sql
-- Role permission lookups
CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);

-- User role lookups
CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX idx_user_roles_role_id ON user_roles(role_id);
```

### Query Examples

#### User with Roles and Permissions

```sql
SELECT
    u.id,
    u.email,
    u.full_name,
    u.is_active,
    r.code as role_code,
    r.name as role_name,
    p.code as permission_code,
    p.service,
    p.method
FROM users u
LEFT JOIN user_roles ur ON u.id = ur.user_id
LEFT JOIN roles r ON ur.role_id = r.id
LEFT JOIN role_permissions rp ON r.id = rp.role_id
LEFT JOIN permissions p ON rp.permission_id = p.id
WHERE u.id = ? AND u.deleted_at IS NULL;
```

#### Active Users Count

```sql
SELECT COUNT(*) as active_users
FROM users
WHERE is_active = true AND deleted_at IS NULL;
```

#### User Activity Logs

```sql
SELECT
    ul.action,
    ul.description,
    ul.ip_address,
    ul.created_at
FROM user_logs ul
WHERE ul.user_id = ?
ORDER BY ul.created_at DESC
LIMIT 50;
```

## 🛡️ Security Considerations

### Password Security

- Passwords are hashed using bcrypt
- Salt rounds: 10 (configurable)
- Never stored in plain text

### Data Protection

- Soft deletes for data retention
- Audit logging for sensitive operations
- Input validation and sanitization

### Access Control

- Role-based access control (RBAC)
- Permission-based authorization
- Database-level constraints

## 📈 Performance Monitoring

### Query Performance

```sql
-- Slow query analysis
SELECT
    query,
    calls,
    total_time,
    mean_time,
    rows
FROM pg_stat_statements
ORDER BY mean_time DESC
LIMIT 10;
```

### Index Usage

```sql
-- Index usage statistics
SELECT
    schemaname,
    tablename,
    indexname,
    idx_scan,
    idx_tup_read,
    idx_tup_fetch
FROM pg_stat_user_indexes
ORDER BY idx_scan DESC;
```

### Table Statistics

```sql
-- Table size and statistics
SELECT
    schemaname,
    tablename,
    attname,
    n_distinct,
    correlation
FROM pg_stats
WHERE schemaname = 'public'
ORDER BY tablename, attname;
```

## 🔧 Maintenance

### Regular Maintenance Tasks

```sql
-- Update table statistics
ANALYZE;

-- Vacuum tables
VACUUM ANALYZE users;
VACUUM ANALYZE user_logs;

-- Clean up old logs (older than 90 days)
DELETE FROM user_logs
WHERE created_at < NOW() - INTERVAL '90 days';
```

### Backup Strategy

```bash
# Full database backup
pg_dump -h localhost -U postgres -d go_template > backup.sql

# Incremental backup (WAL archiving)
# Configure in postgresql.conf
wal_level = replica
archive_mode = on
archive_command = 'cp %p /path/to/archive/%f'
```

This database schema provides a robust foundation for user management, role-based access control, and audit logging in the Go web application template.
