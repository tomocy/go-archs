package usecase

var mockHash = "hash"

type mockHashService struct {
}

func (s mockHashService) GenerateHashFromPassword(plain string) (string, error) {
	return mockHash, nil
}
