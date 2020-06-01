# build image
docker build -t go-rest-api .

# run application
docker run -it -p 8085:8085 --name="go-rest-api" go-rest-api