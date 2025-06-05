hello:
	@echo "Hello, world!"
test:
	@echo "🧪 Running tests..."
	@go test ./... || (echo "❌ Some tests failed. Check above 👆" && exit 1)

bench:
	@echo "🧪 Running tests and benchmarks..."
	@go test -bench=. -benchmem ./... || (echo "❌ Some tests failed. Check above 👆" && exit 1)

all: hello test