FROM node:latest AS builder
WORKDIR /gochat

# Copy the Frontend Directory
COPY ./frontend/src ./src
COPY ./frontend/public ./public
COPY ./frontend/nginx ./nginx
COPY ./frontend/package.json package.json
COPY ./frontend/tsconfig.json tsconfig.json
COPY ./frontend/yarn.lock yarn.lock

# Download the Config File
ARG REACT_APP_API_URL
ARG REACT_APP_WS_URL

RUN echo "REACT_APP_API_URL=$REACT_APP_API_URL\r\n" > .env
RUN echo "REACT_APP_WS_URL=$REACT_APP_WS_URL" >> .env

# Run the Install
RUN yarn install

# Run the Build
RUN yarn build

# Now create the Nginx Server
FROM nginx:alpine

# Create an App Directory to work fro
RUN mkdir -p /gochat/frontend
WORKDIR /gochat/frontend

# Copy the Build Directory
COPY --from=builder ./gochat/build .
COPY --from=builder ./gochat/.env .env

# Copy the Nginx Configuration
COPY --from=builder ./gochat/nginx/frontend.conf /etc/nginx/conf.d/default.conf

# Run the Nginx Daemon
CMD ["nginx", "-g", "daemon off;"]