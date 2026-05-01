-- 05_InsertUserPermission.sql
-- Seed user_permission table.

INSERT INTO user_permission (
  id,
  user_role_id,
  module_name,
  can_create,
  can_read,
  can_update,
  can_delete,
  is_deleted,
  created_date,
  updated_date,
  created_by,
  updated_by
)
SELECT
  '55555555-5555-5555-5555-555555555551',
  '11111111-1111-1111-1111-111111111111',
  'USER',
  TRUE,
  TRUE,
  TRUE,
  TRUE,
  FALSE,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP,
  'SYSTEM',
  'SYSTEM'
WHERE NOT EXISTS (
  SELECT 1
  FROM user_permission
  WHERE user_role_id = '11111111-1111-1111-1111-111111111111'
    AND module_name = 'USER'
);

INSERT INTO user_permission (
  id,
  user_role_id,
  module_name,
  can_create,
  can_read,
  can_update,
  can_delete,
  is_deleted,
  created_date,
  updated_date,
  created_by,
  updated_by
)
SELECT
  '55555555-5555-5555-5555-555555555552',
  '22222222-2222-2222-2222-222222222222',
  'USER',
  FALSE,
  TRUE,
  FALSE,
  FALSE,
  FALSE,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP,
  'SYSTEM',
  'SYSTEM'
WHERE NOT EXISTS (
  SELECT 1
  FROM user_permission
  WHERE user_role_id = '22222222-2222-2222-2222-222222222222'
    AND module_name = 'USER'
);
