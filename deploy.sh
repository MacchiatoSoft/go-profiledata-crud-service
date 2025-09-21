echo "Stopping and removing containers..."
docker-compose down -v

echo "Building and starting containers..."
docker-compose up --build -d