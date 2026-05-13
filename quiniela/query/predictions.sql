-- name: CreatePrediction :exec
INSERT INTO quinielas_response_fixture (fixture_id, usuario_id, prediccion_a, prediccion_b)
VALUES ($1, $2, $3, $4);

-- name: GetUserScore :one
SELECT COUNT(*)::bigint AS total_points
FROM quinielas_response_fixture p
JOIN quinielas_fixtures_result r ON p.fixture_id = r.fixture_id
WHERE p.usuario_id = $1
AND (
    (r.score_a > r.score_b AND p.prediccion_a > p.prediccion_b) OR
    (r.score_a < r.score_b AND p.prediccion_a < p.prediccion_b) OR
    (r.score_a = r.score_b AND p.prediccion_a = p.prediccion_b)
);
