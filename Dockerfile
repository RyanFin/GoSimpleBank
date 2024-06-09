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

# Run stage
FROM alpine:3.20
WORKDIR /app
# copy executable from the builder stage to the run stage
COPY --from=builder /app/main .

# inform docker that the container listens on the specified network port at runtime
EXPOSE 8080
# define the default command to run when the container starts
CMD [ "/app/main" ]

# # build Docker image with: $ docker build -t simplebank:latest .
# REPOSITORY          TAG         IMAGE ID       CREATED              SIZE                                     
# simplebank          latest      acab8a99ea88   About a minute ago   590MB
# simplebank image is absolutely huge as it contains the golang elements in every directory
# use a multi-stage build to fix this issue by building the go binary and deploying that only. This will reduce the size of the image significantly 