FROM library/golang

ADD main.go /opt/

RUN  go get github.com/sevlyar/go-daemon \
  && cd /opt/ && go build main.go

ENTRYPOINT ["/opt/main"]
