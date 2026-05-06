CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL
);

CREATE TABLE quinielas_fixtures (
    id INT AUTO_INCREMENT PRIMARY KEY,
    home_team VARCHAR(100),
    away_team VARCHAR(100)
);

CREATE TABLE quinielas_fixtures_result (
    fixture_id INT PRIMARY KEY,
    score_a INT NOT NULL,
    score_b INT NOT NULL,
    FOREIGN KEY (fixture_id) REFERENCES quinielas_fixtures(id)
);

CREATE TABLE quinielas_response_fixture (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fixture_id INT NOT NULL,
    usuario_id INT NOT NULL,
    prediccion_a INT NOT NULL,
    prediccion_b INT NOT NULL,
    FOREIGN KEY (fixture_id) REFERENCES quinielas_fixtures(id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);
