AIR=/Users/claudiojunior/go/bin/air

# Inicia o watch
start:
	@$(AIR) -c .air.toml

# Rodar projeto
run:
	@go run *.go

# Compilar binario (para o seu sistema operacional)
build:
	@mkdir -p bin \
	&& echo "Gerando build..." \
	&& go build -o bin/api *.go \
	&& echo "Build gerado na pasta \"./bin\""

# Compilar binario para Linux
build-linux:
	@mkdir -p bin \
	&& echo "Gerando build (Linux)..." \
	&& GOOS=linux go build -o bin/api-linux *.go \
	&& echo "Build gerado na pasta \"./bin\""

clean:
	@rm -f bin/api
	@rm -f bin/api-linux