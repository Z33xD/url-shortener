# Go URL Shortener

A simple URL shortener built with Go. This project has a basic REST API that generates short links, stores URL mappings in Redis, and redirects users to the original URL.

> This project was built by following a tutorial series by Eddy Wilkinson, with personal experimentation and minor modifications along the way.

## Features

* Shorten long URLs
* Fast redirects to original URLs
* Redis-based storage
* Written in Go

## Tech Stack

* Go (Gin)
* Redis

## Getting Started

### Prerequisites

* Go 1.26+
* Redis

### Clone the repository

```bash
git clone https://github.com/Z33xD/url-shortener.git
cd url-shortener
```

### Install dependencies

```bash
go mod tidy
```

### Start Redis

If you have Redis installed locally:

```bash
redis-server
```

Or use Docker:

```bash
docker run --name redis -p 6379:6379 redis
```

### Run the application

```bash
go run .
```

The server will start on the configured port.

## API

### Create a short URL

**POST**

```
http://localhost:9808/create-short-url
```

Example request:

```json
{
  "long_url": "https://github.com/Z33xD",
  "user_id": "e0dba740-fc4b-4977-872c-d360239e6b10"
}
```

Example response:

```json
{
    "message": "Short URL created successfully",
    "short_url": "http://localhost:9808/cUmP6yWo"
}
```


### Redirect

**GET** (Visiting):

```
http://localhost:9808/cUmP6yWo
```

redirects to the original URL.

## Project Structure

```
.
├── handler/
├── shortener/
├── store/
├── main.go
├── go.mod
├── go.sum
└── README.md
```


## Learning Resources

This project closely follows the "Let's Build a URL Shortener in Go" series by Eddy Wilkinson:

- [Part 1 - Project setup](https://www.eddywm.com/lets-build-a-url-shortener-in-go/)
- [Part 2 - Storage Layer](https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/)
- [Part 3 - Short Link Generator](https://www.eddywm.com/lets-build-a-url-shortener-in-go-part-3-short-link-generation/)
- [Part 4 - Forwarding](https://www.eddywm.com/lets-build-a-url-shortener-in-go-part-iv-forwarding/)

The tutorials provide an in-depth walkthrough of building a URL shortener from scratch using Go and Redis.