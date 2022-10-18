CREATE TABLE IF NOT EXISTS course (
  id VARCHAR(255) NOT NULL,
  teacher_id VARCHAR(255),
  subject_id VARCHAR(255),
  course_name VARCHAR(255),
  start_day TIMESTAMP,
  course_extend FLOAT,
  credit_amount INT,
  slot INT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
)