CREATE TABLE IF NOT EXISTS quinielas_response_fixture (
    id SERIAL PRIMARY KEY,
    fixture_id INTEGER NOT NULL REFERENCES quinielas_fixtures (id),
    usuario_id INTEGER NOT NULL REFERENCES usuarios (id),
    prediccion_a INTEGER NOT NULL,
    prediccion_b INTEGER NOT NULL
);
