FROM  golang:1.16.5-alpine AS dev
# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

COPY . .

#Build app
RUN go build -o /weebquote

CMD [ "/weebquote" ]