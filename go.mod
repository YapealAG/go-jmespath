module github.com/YapealAG/go-jmespath

go 1.14

replace github.com/YapealAG/go-jmespath/internal/testify => ./internal/testify

require (
	github.com/jmespath/go-jmespath v0.4.0
	github.com/jmespath/go-jmespath/internal/testify v1.5.1
)
