# GoChat

GoChat is a microservice-based chat application build in Go and Typescript.

The application demonstrates the following:

- Microservice Architecture using gRPC
- React front-end application design
- Golang back-end application design
- Client / Server application design

## Platform

The platform is a microservice-driven application built in Go (Golang). It is divided into `services` and `endpoints`. Services are internal microservices that handle data ingress and egress. Endpoints are publicly-available access-points for the data (for example: sending and receiving messages).

## Frontend

The frontend is a web-based application built in Typescript and React. It allows users to create accounts, log in, and send and receive messages on the chat server.

## Running the Application

### Environment Variables

This application needs a few environment variables in order to run properly. The easiest way is to create a copy of the `.env.dist` file called `.env`. Please be sure to change the passwords and authentication keys.

**Note** The `REACT_APP_API_URL` and `REACT_APP_WS_URL` specified in `docker-compose.yml` should point to your network IP address (if you're running on a local network) or the domain name / server IP (if you deploy to the web).

### Running in Docker

To run the application, you'll need [Docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/) installed. Once you have these, just go into the directory where you cloned this repository and run:

    docker-compose up -d

## Donations

I enjoy writing these sample applications as they allow me to stretch my engineering brain a little. If you have found my code useful, please consider making a donation to my Paypal account (link shown below).

[Click here to donate](https://www.paypal.me/reverendandrewmills)
