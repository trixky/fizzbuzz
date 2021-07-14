FROM node:latest

# create the project folder
RUN mkdir /home/node/supertest
# copy the source
COPY . /home/node/supertest
WORKDIR /home/node/supertest

# install dependencies
RUN npm install -g npm@7.19.1
RUN npm install jest -g
RUN npm install

ENTRYPOINT npx jest --runInBand

# --runInBand : Run all tests serially in the current process, rather than creating a worker pool of child processes that run tests. This can be useful for debugging.
