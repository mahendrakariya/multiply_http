FROM golang:1.12.9-stretch
EXPOSE 8001

COPY multiply multiply

CMD ["./multiply"]

