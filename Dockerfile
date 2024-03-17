FROM golang:1.21-alpine 

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copie o código fonte da aplicação para o contêiner
COPY . .

# Baixe e instale as dependências da aplicação
RUN go mod download

# Compile a aplicação
RUN go build -o api .

# Especifique como a aplicação deve ser executada
CMD ["./api"]
