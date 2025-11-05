# syntax=docker/dockerfile:1
FROM golang:tip-trixie

# Make the Working Directory
WORKDIR /app

# Copy across source code
COPY . ./

# Download Go modules
RUN make setup

# Expose the HTTP Port
EXPOSE 8080

# Run
CMD ["make", "run"]
