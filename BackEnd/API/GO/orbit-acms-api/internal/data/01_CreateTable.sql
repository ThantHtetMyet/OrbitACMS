-- 01_CreateTable.sql
-- Create core tables for Orbit ACMS.

CREATE TABLE IF NOT EXISTS user_role (
  id              VARCHAR(36) PRIMARY KEY,
  role_name       VARCHAR(100) NOT NULL UNIQUE,
  description     VARCHAR(500),
  is_deleted      BOOLEAN NOT NULL DEFAULT FALSE,
  created_date    TIMESTAMP,
  updated_date    TIMESTAMP,
  created_by      VARCHAR(100),
  updated_by      VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS users (
  id              VARCHAR(36) PRIMARY KEY,
  user_role_id    VARCHAR(36) NOT NULL,
  first_name      VARCHAR(100) NOT NULL,
  last_name       VARCHAR(100) NOT NULL,
  email           VARCHAR(255) NOT NULL UNIQUE,
  mobile_no       VARCHAR(30) NOT NULL,
  login_password  VARCHAR(255) NOT NULL,
  remark          VARCHAR(500),
  last_login      TIMESTAMP,
  is_active       BOOLEAN NOT NULL DEFAULT TRUE,
  is_deleted      BOOLEAN NOT NULL DEFAULT FALSE,
  created_date    TIMESTAMP NOT NULL,
  updated_date    TIMESTAMP NOT NULL,
  created_by      VARCHAR(100),
  updated_by      VARCHAR(100),
  CONSTRAINT fk_users_user_role
    FOREIGN KEY (user_role_id) REFERENCES user_role(id)
);

CREATE TABLE IF NOT EXISTS user_permission (
  id              VARCHAR(36) PRIMARY KEY,
  user_role_id    VARCHAR(36) NOT NULL,
  module_name     VARCHAR(100) NOT NULL,
  can_create      BOOLEAN NOT NULL DEFAULT FALSE,
  can_read        BOOLEAN NOT NULL DEFAULT FALSE,
  can_update      BOOLEAN NOT NULL DEFAULT FALSE,
  can_delete      BOOLEAN NOT NULL DEFAULT FALSE,
  is_deleted      BOOLEAN NOT NULL DEFAULT FALSE,
  created_date    TIMESTAMP,
  updated_date    TIMESTAMP,
  created_by      VARCHAR(100),
  updated_by      VARCHAR(100),
  CONSTRAINT fk_user_permission_user_role
    FOREIGN KEY (user_role_id) REFERENCES user_role(id),
  CONSTRAINT uq_user_permission_role_module
    UNIQUE (user_role_id, module_name)
);

CREATE TABLE IF NOT EXISTS image_type (
  id                VARCHAR(36) PRIMARY KEY,
  image_type_name   VARCHAR(100) NOT NULL UNIQUE,
  is_deleted        BOOLEAN NOT NULL DEFAULT FALSE,
  created_date      TIMESTAMP NOT NULL,
  updated_date      TIMESTAMP NOT NULL,
  created_by        VARCHAR(100),
  updated_by        VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS user_image (
  id                VARCHAR(36) PRIMARY KEY,
  user_id           VARCHAR(36) NOT NULL,
  image_type_id     VARCHAR(36) NOT NULL,
  image_name        VARCHAR(255) NOT NULL,
  stored_directory  VARCHAR(500) NOT NULL,
  is_deleted        BOOLEAN NOT NULL DEFAULT FALSE,
  uploaded_date     TIMESTAMP NOT NULL,
  uploaded_by       VARCHAR(100),
  CONSTRAINT fk_user_image_user
    FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_user_image_type
    FOREIGN KEY (image_type_id) REFERENCES image_type(id)
);

CREATE INDEX IF NOT EXISTS idx_users_user_role_id ON users(user_role_id);
CREATE INDEX IF NOT EXISTS idx_user_permission_user_role_id ON user_permission(user_role_id);
CREATE INDEX IF NOT EXISTS idx_user_image_user_id ON user_image(user_id);
CREATE INDEX IF NOT EXISTS idx_user_image_image_type_id ON user_image(image_type_id);

