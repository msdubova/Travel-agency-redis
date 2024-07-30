FROM golang

WORKDIR /app

COPY . .

WORKDIR /app/cmd

RUN go build -o /app/travelagency .

CMD ["/app/travelagency"]


