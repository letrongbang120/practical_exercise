CREATE TABLE IF NOT EXISTS student (
  id VARCHAR(255) NOT NULL,
  student_id VARCHAR(255),
  name VARCHAR(255),
  email VARCHAR(255),
  pass VARCHAR(255),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
)