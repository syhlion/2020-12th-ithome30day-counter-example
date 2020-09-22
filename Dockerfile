FROM golang:1.15.2-alpine3.12
WORKDIR /ithome30day-counter
ADD . /ithome30day-counter
RUN cd /ithome30day-counter && go build 
EXPOSE 8080
ENTRYPOINT ./ithome30day-counter
