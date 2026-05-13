CREATE TABLE IF NOT EXISTS quinielas_fixtures_result (
    fixture_id INTEGER PRIMARY KEY REFERENCES quinielas_fixtures (id),
    score_a INTEGER NOT NULL,
    score_b INTEGER NOT NULL
);
