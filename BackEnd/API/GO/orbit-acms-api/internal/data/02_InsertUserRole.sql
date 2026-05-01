-- 02_InsertUserRole.sql
-- Seed user_role table.

INSERT INTO user_role (id, role_name, description, is_deleted, created_date, updated_date, created_by, updated_by)
SELECT '11111111-1111-1111-1111-111111111111', 'ADMIN', 'System administrator role', FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'SYSTEM', 'SYSTEM'
WHERE NOT EXISTS (
  SELECT 1 FROM user_role WHERE role_name = 'ADMIN'
);

INSERT INTO user_role (id, role_name, description, is_deleted, created_date, updated_date, created_by, updated_by)
SELECT '22222222-2222-2222-2222-222222222222', 'USER', 'Standard application user role', FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'SYSTEM', 'SYSTEM'
WHERE NOT EXISTS (
  SELECT 1 FROM user_role WHERE role_name = 'USER'
);
