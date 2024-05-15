FROM golang:1.22-alpine as builder

LABEL authors="brunooliveira"

# Definir diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos de módulo primeiro para aproveitar o cache de camadas Docker
COPY go.mod go.sum ./

# Baixar dependências de forma segura e verificar a integridade
RUN go mod download
RUN go mod verify

# Copiar o resto do código fonte para o diretório de trabalho
COPY . .

# Compilar o aplicativo para o binário específico, otimizando para ambientes alpine
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/main ./cmd

# Uso de uma imagem base limpa para a imagem final
FROM alpine:latest
WORKDIR /root/

# Copiar apenas o binário compilado do estágio de construção para reduzir o tamanho da imagem
COPY --from=builder /app/main .
COPY --from=builder /app/cmd/.env .

# Configurações adicionais podem ser especificadas aqui, como variáveis de ambiente, volumes, etc., se necessário
COPY wait-for.sh .
RUN chmod +x wait-for.sh # make the script executable

# Comando para executar o binário
CMD ["./wait-for.sh", "mysql:3306", "./wait-for.sh", "rabbitmq:5672", "./main"]
