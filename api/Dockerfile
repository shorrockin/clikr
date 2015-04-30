FROM google/golang

RUN go get github.com/gorilla/mux
RUN go get github.com/skelterjohn/rerun

RUN mkdir -p /gopath/src/github.com/shorrockin/clikr
ADD ./src /gopath/src/github.com/shorrockin/clikr
RUN go get github.com/shorrockin/clikr

EXPOSE 3000
CMD ["/gopath/bin/clikr"]
