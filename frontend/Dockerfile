FROM node:latest AS builder
WORKDIR /gochat

# Copy the Frontend Directory
COPY .env .env
COPY ./frontend/src ./src
COPY ./frontend/public ./public
COPY ./frontend/nginx ./nginx
COPY ./frontend/package.json package.json
COPY ./frontend/tsconfig.json tsconfig.json
COPY ./frontend/yarn.lock yarn.lock

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

# Copy the Nginx Configuration
COPY --from=builder ./gochat/nginx/frontend.conf /etc/nginx/conf.d/default.conf

# Run the Nginx Daemon
CMD ["nginx", "-g", "daemon off;"]