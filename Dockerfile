# Etapa de compilación
FROM golang:1.23-alpine AS builder

# Habilita bash para debug y más utilidades (opcional)
RUN apk add --no-cache bash

WORKDIR /app

# Copiar primero los archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Forzar compilación completa del binario
RUN go build -a -installsuffix cgo -o parking-status ./cmd/main.go

# Etapa final con imagen mínima
FROM alpine:latest

WORKDIR /app

# Copiar solo el binario desde la etapa anterior
COPY --from=builder /app/parking-status .

# Exponer el puerto usado por el servicio
EXPOSE 8080

# Ejecutar el binario
CMD ["./parking-status"]
