compile: clean
	@echo "Compilando arquivos .proto."
	@protoc -I=src/ --js_out=import_style=commonjs,binary:types/js --go_out=types/go --go_opt=paths=source_relative ./src/*.proto

clean:
	@echo "Limpando compilações anteriores."
	@rm -rf ./src/*.go
