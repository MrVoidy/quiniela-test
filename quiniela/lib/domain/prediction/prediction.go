package prediction

// Prediction is a usuario's score guess for a fixture (quinielas_response_fixture row).
type Prediction struct {
	FixtureID int32
	UsuarioID int32
	PredA     int32
	PredB     int32
}
