FROM node:18

WORKDIR /usr/src/app

COPY ./node/package*.json ./

RUN npm install

COPY ./node .

EXPOSE 3000

CMD ["npm", "run", "dev"]
