package mocks

//go:generate mockery --name "Ads" --dir ../domain --output ./ --filename "adsMock.go" --with-expecter
//go:generate mockery --name "Clock" --dir ../domain --output ./ --filename "clockMock.go" --with-expecter
//go:generate mockery --name "IdGenerator" --dir ../domain --output ./ --filename "idGeneratorMock.go" --with-expecter
//go:generate mockery --name "CreateAdService" --dir ../service --output ./ --filename "service/createAdServiceMock.go" --with-expecter
//go:generate mockery --name "FindAdService" --dir ../service --output ./ --filename "service/FindAdServiceMock.go" --with-expecter
