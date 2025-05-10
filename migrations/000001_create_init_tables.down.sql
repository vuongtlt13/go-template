-- Drop indexes
DROP INDEX IF EXISTS idx_role_permissions_permission_id;
DROP INDEX IF EXISTS idx_role_permissions_role_id;
DROP INDEX IF EXISTS idx_permissions_name;
DROP INDEX IF EXISTS idx_roles_name;
DROP INDEX IF EXISTS idx_users_role_id;
DROP INDEX IF EXISTS idx_users_email;

-- Drop foreign key constraints
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_users_role;

-- Drop tables in reverse order (to handle dependencies)
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
