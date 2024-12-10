# Usar una imagen base oficial de Go
FROM golang:1.22-alpine

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar los archivos Go al contenedor
COPY . .

# Descargar las dependencias
RUN go mod tidy

# Instalar swag para generar la documentación
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generar la documentación Swagger
RUN swag init

# Copiar el directorio docs generado por swag
COPY docs/ /app/docs/

# Compilar la aplicación
RUN go build -o main .

# Exponer los puertos en los que los servidores HTTP y RPC escuchan
EXPOSE 8080 1234

# Comando para ejecutar la aplicación
CMD ["./main"]