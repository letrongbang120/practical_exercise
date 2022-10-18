CREATE TABLE IF NOT EXISTS subjects (
  id VARCHAR(255) NOT NULL,
  subject_name VARCHAR(255),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
)