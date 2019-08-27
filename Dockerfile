FROM golang:1.12.9-stretch

COPY multiply multiply

CMD ["./multiply"]

