package mocks

//go:generate mockery --name "Ads" --dir ../domain --output ./ --filename "adsMock.go" --with-expecter
//go:generate mockery --name "Clock" --dir ../domain --output ./ --filename "clockMock.go" --with-expecter
//go:generate mockery --name "IdGenerator" --dir ../domain --output ./ --filename "idGeneratorMock.go" --with-expecter
