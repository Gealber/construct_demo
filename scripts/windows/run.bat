echo "Declaring environment variables"

rem Yous should change this two path for the correct path to the public and private keys

set PRIVATE_KEY_PATH="/home/gealber/go/src/github.com/Gealber/construct_demo/serializer/jwt/key_backup/id_rsa"
set PUBLIC_KEY_PATH="/home/gealber/go/src/github.com/Gealber/construct_demo/serializer/jwt/key_backup/id_rsa.pub"

set MONGO_URL="localhost"
set REDIS_PASSW="construct_demo.321"
set REDIS_URL="redis://default:$REDIS_PASSW@localhost:6379/0"

go run ..\..\main.go
