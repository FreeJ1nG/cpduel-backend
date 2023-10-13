CREATE TABLE IF NOT EXISTS Problem (
  id VARCHAR(50) PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  time_limit VARCHAR(20) NOT NULL,
  memory_limit VARCHAR(20) NOT NULL,
  input_type VARCHAR(30) NOT NULL,
  output_type VARCHAR(30) NOT NULL,
  difficulty INT NOT NULL,
  body text NOT NULL
);