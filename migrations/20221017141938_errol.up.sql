CREATE TABLE IF NOT EXISTS course (
  id VARCHAR(255) NOT NULL,
  course_id VARCHAR(255),
  student_id VARCHAR(255),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
)