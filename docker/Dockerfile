FROM library/golang

ADD main.go /opt/

RUN  go get github.com/sevlyar/go-daemon \
  && go build /opt/main.go

ENTRYPOINT ["/opt/main"]
