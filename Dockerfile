FROM tedgkassen/1.16-buster-debug:latest

ARG platform=telegram

COPY . src

WORKDIR /go/src

RUN make setup
RUN make build_$platform
RUN cp bin/$platform /go/app

CMD "/go/app"
