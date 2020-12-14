echo "Declaring environment variables"

rem Yous should change this two path for the correct path to the public and private keys

set PRIVATE_KEY_PATH=""
set PUBLIC_KEY_PATH=""

set MONGO_URL="localhost"
set REDIS_URL=""

go run ..\..\main.go
