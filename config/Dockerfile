FROM golang:1.23-alpine
RUN apk update && apk add --no-cache gcc musl-dev
WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
EXPOSE 8001
EXPOSE 11001
CMD ["go", "run", ".", "-web"]
