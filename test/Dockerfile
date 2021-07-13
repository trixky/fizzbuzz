FROM node:latest

# create the project folder
RUN mkdir /home/node/test
# copy the source
COPY . /home/node/test
WORKDIR /home/node/test

# install dependencies
RUN npm install -g npm@7.19.1
RUN npm install jest -g
RUN npm install

ENTRYPOINT npx jest --runInBand

# --runInBand : Run all tests serially in the current process, rather than creating a worker pool of child processes that run tests. This can be useful for debugging.
