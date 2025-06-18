# Seat Arrangements Application

A full-stack application for managing aircraft seat arrangements and bookings.

## Project Structure

The project is divided into two main parts:

- **Backend**: A Go-based REST API using Fiber framework
- **Frontend**: A React application using TypeScript and Vite

## Features

- Seat arrangement visualization

## Technologies Used

### Backend

- Go (Golang)
- Fiber web framework
- PostgreSQL database
- Cobra CLI
- Viper configuration
- Zap logger

### Frontend

- React
- TypeScript
- Vite
- Tailwind CSS

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.21 or higher (for local development)
- Node.js 20 or higher (for local development)
- PostgreSQL (for local development without Docker)

### Running with Docker

1. Clone the repository

```bash
git clone https://github.com/evaizee/plane-seat-arrangements.git
cd plane-seat-arrangements
```

2. Start the application using Docker Compose

```bash
docker-compose up -d
```

3. Access the application
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080

### Local Development

#### Backend

1. Navigate to the backend directory

```bash
cd backend
```

2. Install dependencies

```bash
go mod download
```

3. Run database migrations

```bash
go run main.go migrate up
```

4. Start the server

```bash
go run main.go serve
```

#### Frontend

1. Navigate to the frontend directory

```bash
cd frontend
```

2. Install dependencies

```bash
npm install
```

3. Start the development server

```bash
npm run dev
```

## API Documentation

The API documentation is available at `/api/docs` when the backend server is running.

## Database Schema

The database schema includes the following main tables:

- users
- itineraries
- segments
- aircraft
- cabins
- seat_rows
- seats
- seat_prices
- bookings
- booking_seats

## License

This project is licensed under the MIT License - see the LICENSE file for details.