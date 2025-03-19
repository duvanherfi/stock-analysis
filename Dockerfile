FROM node:22.14.0 AS build-node

WORKDIR /usr/app
COPY ./ui /usr/app
RUN npm install
RUN npm run build

FROM golang:1.24.1 AS build-go

ARG API_URL=""
ARG API_KEY=""
ARG DATABASE_URL=""

WORKDIR /usr/app
COPY . .

RUN echo $(ls -la)

RUN go build -o ./stock-analysis

FROM golang:1.24.1 AS deploy

WORKDIR /usr/app
RUN mkdir -p /usr/app/ui/dist
COPY --from=build-go /usr/app/stock-analysis /usr/app/stock-analysis
COPY --from=build-node /usr/app/dist /usr/app/ui/dist

EXPOSE 3000

CMD ["./stock-analysis"]