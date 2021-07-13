FROM node:latest

RUN mkdir /home/node/test
COPY . /home/node/test
WORKDIR /home/node/test

RUN npm install -g npm@7.19.1
RUN npm install jest -g
RUN npm install

ENTRYPOINT npx jest --runInBand
