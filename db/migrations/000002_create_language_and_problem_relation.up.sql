CREATE TABLE IF NOT EXISTS Language (
  id VARCHAR(40) PRIMARY KEY,
  value VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS LanguageOfProblem (
  problem_id VARCHAR(50) NOT NULL,
  language_id VARCHAR(40) NOT NULL,
  PRIMARY KEY (problem_id, language_id),
  FOREIGN KEY (language_id) REFERENCES Language(id) ON DELETE CASCADE,
  FOREIGN KEY (problem_id) REFERENCES Problem(id) ON DELETE CASCADE
);