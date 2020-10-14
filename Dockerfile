FROM golang:1.15 as BUILD-BACKEND
RUN apt-get update && apt-get install sqlite3 -y && mkdir /db 
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN rm -fr ./frontend
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /bin/app .
RUN chmod -R 777 db.sh && ./db.sh
#RUN CGO_ENABLED=0 go build -o /bin/backend .

FROM node:12.18.4-stretch-slim as BUILD-FRONTEND
WORKDIR /src
COPY frontend .
RUN npm install
RUN npm run build && cp -fr build /bin/

#FROM node:12.18.4-stretch-slim
FROM nginx
WORKDIR /app
#RUN npm install -g serve && mkdir build && mkdir db
RUN mkdir build && mkdir db
COPY --from=BUILD-BACKEND /bin/app .
COPY --from=BUILD-BACKEND /db/database.db ./db/
COPY --from=BUILD-FRONTEND /bin/build ./build/
COPY entrypoint.sh .
COPY default.conf /etc/nginx/conf.d/default.conf
RUN chmod 777 entrypoint.sh
ENTRYPOINT ./entrypoint.sh