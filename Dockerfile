# Użyj obrazu Golang
FROM golang:1.23

# Ustaw katalog roboczy
WORKDIR /app

# Skopiuj pliki go.mod i go.sum oraz pobierz zależności
COPY go.mod go.sum ./
RUN go mod download

# Skopiuj resztę kodu aplikacji
COPY . .

# Zbuduj aplikację Go
RUN go build -o main .

# Eksponuj port aplikacji Fiber
EXPOSE 8080

# Uruchom aplikację Fiber
CMD ["./main"]
