##//https://tutorialedge.net/golang/go-docker-tutorial/

FROM golang:1.12.0-alpine3.9
## adding git to the Alpine variant can currently be accomplished with a very trivial Dockerfile similar to the following:

 RUN apk add --no-cache git

# create a working directory
WORKDIR /app
# add source code
ADD . /app

# Expose port 8085 to the outside world
EXPOSE 8085


# run all go files 
#CMD ["go", "run", "."]
CMD ["./main"]