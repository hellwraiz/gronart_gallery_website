# Stage 1: Build frontend
FROM node:25 AS frontend-builder
WORKDIR /build
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend/ ./
RUN npm run build

# Stage 2: Build Go app
FROM golang:1.25-bookworm AS go-builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
# Copy the built frontend from previous stage
COPY --from=frontend-builder /build/build ./frontend/build
RUN go build -v -o run-app .

# Stage 3: Final runtime image
FROM debian:bookworm
WORKDIR /app
COPY --from=go-builder /build/run-app /usr/local/bin/
COPY --from=go-builder /build/migrations ./migrations
COPY --from=go-builder /build/frontend/build ./frontend/build
EXPOSE 8080
CMD ["run-app"]
