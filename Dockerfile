# Użyj obrazu Golang
FROM golang:1.23-alpine
RUN apk update && apk add git openssh

# Skopiuj klucze SSH do kontenera (lokalizacja na Windowsie)
COPY ./.ssh /root/.ssh

# Ustawienia uprawnień dla kluczy SSH
RUN chmod 600 /root/.ssh/id_rsa && chmod 600 /root/.ssh/id_rsa.pub

# Skonfiguruj Git do używania SSH zamiast HTTPS
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# Skopiuj pliki go.mod i go.sum oraz pobierz zależności
COPY go.mod go.sum ./
RUN go mod download

# Skopiuj resztę kodu aplikacji
COPY . .

# Zbuduj aplikację
RUN go build -o main .

# Eksponuj port aplikacji
EXPOSE 8080

# Uruchom aplikację
CMD ["./main"]