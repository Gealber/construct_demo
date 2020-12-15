echo "Declaring environment variables"

rem Yous should change this two path for the correct path to the public and private keys

<<<<<<< HEAD
SET PRIVATE_KEY_PATH=E:/Kmada Proyectos/Proyectos Freelancer/Tests/construct_demo/serializer/jwt/key_backup/id_rsa
SET PUBLIC_KEY_PATH=E:/Kmada Proyectos/Proyectos Freelancer/Tests/construct_demo/serializer/jwt/key_backup/id_rsa.pub

set MONGO_URL="localhost:27017"
set REDIS_PASSW="construct_demo.321"
set REDIS_URL="redis://default:$REDIS_PASSW@localhost:6379/0"
=======
set PRIVATE_KEY_PATH=""
set PUBLIC_KEY_PATH=""

set MONGO_URL="localhost"
set REDIS_URL=""
>>>>>>> cc83d9365ee390d8940e78f34228efdd119adab6

go run ../../main.go
