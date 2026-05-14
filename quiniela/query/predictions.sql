-- name: CreatePrediction :exec
INSERT INTO fixture_predictions (fixture_id, user_id, pred_a, pred_b)
VALUES ($1, $2, $3, $4);

-- name: GetUserScore :one
SELECT COUNT(*)::bigint AS total_points
FROM fixture_predictions p
JOIN fixture_results r ON p.fixture_id = r.fixture_id
WHERE p.user_id = $1
AND (
    (r.score_a > r.score_b AND p.pred_a > p.pred_b) OR
    (r.score_a < r.score_b AND p.pred_a < p.pred_b) OR
    (r.score_a = r.score_b AND p.pred_a = p.pred_b)
);
