# Book Service

A simple Go service to fetch book information by ISBN using the Google Books API.

## Features
- GET `/book/{isbn}`: Returns book data (title, authors, description) from Google Books.
- Docker and Docker Compose ready.
- CI/CD pipeline with GitHub Actions (self-hosted runner).

## Requirements
- Go 1.21+
- Docker (optional, for containerized usage)
- Google Books API Key

## Environment Variables
- `GOOGLE_BOOKS_API_KEY`: Your Google Books API key (required).

## Running Locally
```sh
go run main.go
```

## Running with Docker
Build and run:
```sh
docker build -t bookservice .
docker run -e GOOGLE_BOOKS_API_KEY=your_key -p 8080:8080 bookservice
```

## Running with Docker Compose
Create a `.env` file with:
```
GOOGLE_BOOKS_API_KEY=your_key
```
Then run:
```sh
docker-compose up --build
```

## API Example
```
GET /book/9780140449136
Response:
{
  "title": "The Odyssey",
  "authors": ["Homer"],
  "description": "..."
}
```

## Running Tests
```sh
go test ./...
```

## CI/CD
- On every push to `master`, GitHub Actions will:
  - Run tests
  - Build Docker image
  - Deploy using Docker Compose on a self-hosted runner
  - The Google Books API key is injected via repository secrets

## License
MIT 