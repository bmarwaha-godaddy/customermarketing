1 Run docker build . e.g docker built . -t market:dev

2.Once docker image is built initiate the container using below command

docker run --env-file env-var.txt --publish 4000:8080 market:dev env

Make sure to set AWS credentials in env-var.txt file.

POST
http://localhost:4000/channel
```json lines
   {
   "name" :"Google",
   "identifier": "asasa",
   "loggedInFrom": "LDH",
   "customerId": "121212121"

}
```

2. GET URL.
http://localhost:4000/channel/?identifier=abc&loggedInCity=LDH

3. Delete URL
http://localhost:4000/channel/?identifier=12121&loggedInCity=Amritsar
