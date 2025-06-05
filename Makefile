hello:
	@echo "👋 Hello, world!"

test:
	@echo "🧪 Running tests..."
	@go test ./... || (echo "❌ Some tests failed. Check above 👆" && exit 1)

testv:
	@echo "🧪 Running verbose tests..."
	@go test -v ./... || (echo "❌ Some tests failed. Check above 👆" && exit 1)

bench:
	@echo "🏋️ Running benchmarks..."
	@go test -bench=. -benchmem ./... || (echo "❌ Some benchmarks failed. Check above 👆" && exit 1)

coverage:
	@echo "📊 Running test coverage..."
	@go test -cover ./...

all: hello test
