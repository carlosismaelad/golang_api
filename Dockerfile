# Usa a versão mais recente estável do Go
FROM golang:1.23 AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos do projeto para dentro do container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila o binário do projeto
RUN go build -o main ./src

# Usa uma imagem menor para execução
FROM golang:1.21 AS runner
WORKDIR /app

# Copia apenas o binário compilado
COPY --from=builder /app/main .

# Expõe a porta usada pelo servidor
EXPOSE 8080

# Comando para rodar o projeto
CMD ["./main"]
