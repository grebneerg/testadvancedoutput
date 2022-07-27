FROM golang:1.18

COPY . /bundle
WORKDIR /bundle
RUN go build -o testao
RUN chmod +x testao

ENTRYPOINT [ "/bundle/testao" ]
