CREATE TABLE IF NOT EXISTS fixture_predictions (
    id SERIAL PRIMARY KEY,
    fixture_id INTEGER NOT NULL REFERENCES fixtures (id),
    user_id UUID NOT NULL REFERENCES users (id),
    pred_a INTEGER NOT NULL,
    pred_b INTEGER NOT NULL
);
