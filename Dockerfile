FROM node:16

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install
RUN ls -la

COPY . .

RUN ls -la

CMD [ "node", "main.js" ]