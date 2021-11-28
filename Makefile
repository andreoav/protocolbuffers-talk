compile: compile-go compile-js
	@echo "Compilando arquivos .proto."

compile-go:
	@protoc -I=src/ --go_out=types/go --go_opt=paths=source_relative --go-grpc_out=types/go --go-grpc_opt=paths=source_relative ./src/*.proto

compile-js:
	@protoc -I=src/ --js_out=import_style=commonjs,binary:types/js ./src/*.proto
