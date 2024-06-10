# Episode 25 of the Backend Masterclass: https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/learn/lecture/25822346#overview
# state the golang version being used from the golang page in Dockerhub
#  https://hub.docker.com/_/golang | Use the most recent alpine version
# Build stage 
# specify the builder for multi-stage deployment to run the binary/executable
FROM golang:1.22.4-alpine3.20 As builder 
# declare the current working directory inside the Docker image
WORKDIR /app
# Copy everything from the current folder (root of this GoSimpleBank project) will be copied, represented by the first .
# Second . represents the location of the copied data
COPY . .
RUN go build -o main main.go
# install curl 
RUN apk add curl
# download go migrate library
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.darwin-amd64.tar.gz | tar xvz
        

# Run stage
FROM alpine:3.20
WORKDIR /app
# copy executable from the builder stage to the run stage env
COPY --from=builder /app/main .
# copy migrate binary to a new migrate directory in the run stage env 
COPY --from=builder /app/migrate ./migrate
# copy the viper config file to the docker run stage env
COPY app.yml .
COPY start.sh .
COPY wait-for.sh .
# copy all migration files from this project to the Docker file in a new migration directory
COPY db/migration ./migration


# inform docker that the container listens on the specified network port at runtime
EXPOSE 8080
# define the default command to run when the container starts
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]

# REPOSITORY          TAG         IMAGE ID       CREATED              SIZE                                     
# simplebank          latest      acab8a99ea88   About a minute ago   590MB
# simplebank image is absolutely huge as it contains the golang elements in every directory
# use a multi-stage build to fix this issue by building the go binary and deploying that only. This will reduce the size of the image significantly 

# --  Docker Commands --
# build Docker image with: $ docker build -t simplebank:latest .
# Run docker app from image (also creates a container that can be reran): 
    # 1. $ docker run --name simplebank -p 8080:8080 simplebank:latest
    # 2. (with envvar) $ docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgrel://root:root@postgres-container:5432/simple_bank?sslmode=disable" simplebank:latest
# View containers with: $ docker ps -a
# Remove existing containers with the command: $ docker rm <container ID>
# Remove existing images with the command: $ docker rmi <image ID>
# Get IP address by viewing the network settings of the postgres db: $ docker container inspect cd21302ab9fd
