CREATE TABLE IF NOT EXISTS fixture_results (
    fixture_id INTEGER PRIMARY KEY REFERENCES fixtures (id),
    score_a INTEGER NOT NULL,
    score_b INTEGER NOT NULL
);
