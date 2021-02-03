FROM golang:1.16rc1-buster

RUN mkdir -p /usr/src/GoWebProject/
WORKDIR /usr/src/GoWebProject/

COPY . /usr/src/GoWebProject/

#RUN go get github.com/jackc/pgx
#RUN go get github.com/bmizerany/pat
#RUN go get github.com/golangcollege/sessions
#RUN go get github.com/justinas/alice

#EXPOSE 8000

CMD ["go", "run", "./cmd/web"]