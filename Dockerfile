FROM node:14
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install && npm install express
COPY VERSION /
COPY . .
EXPOSE 3000
CMD [ "node", "server.js" ]