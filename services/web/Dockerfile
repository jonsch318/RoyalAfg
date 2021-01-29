FROM node:current-alpine

ENV PORT 3000

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Installing dependencies
COPY package.json yarn.lock /usr/src/app/
RUN yarn install

# Copying source files
COPY . /usr/src/app

# Building app
RUN yarn run build
EXPOSE 3000

# Running the app
CMD "yarn" "run" "start"