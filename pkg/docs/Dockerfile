FROM golang:1.15.0

WORKDIR $GOPATH/src/github.com/JohnnyS318/RoyalAfgInGo/docs

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 9000

CMD [ "docs" ]
