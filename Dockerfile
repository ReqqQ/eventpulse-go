# Użyj obrazu Golang
FROM golang:1.23

# Ustaw katalog roboczy
WORKDIR /app

# Skopiuj lokalny moduł eventpulse-user-go do /app/eventpulse-user-go
COPY ./eventpulse-user-go /app/eventpulse-user-go

# Skopiuj pliki go.mod i go.sum oraz pobierz zależności
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod edit -replace github.com/ReqqQ/eventpulse-user-go=/app/eventpulse-user-go
RUN go mod download

# Skopiuj resztę kodu aplikacji backend
COPY ./backend /app

# Zbuduj aplikację Go
RUN go build -o main .

# Eksponuj port aplikacji Fiber
EXPOSE 8080

# Uruchom aplikację Fiber
CMD ["./main"]
