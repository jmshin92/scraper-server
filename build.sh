BUILD_DIR="./bin"
mkdir $BUILD_DIR

echo "Building linux amd64..."
GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/scraper.linux.amd64

echo "Building darwin amd64..."
GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/scraper.darwin.amd64
