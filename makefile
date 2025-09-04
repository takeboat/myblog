.Phony: t, clean,generate

t:
	@go mod tidy
generate:
	@goctl api go -api ./api/blog.api -dir ./api

