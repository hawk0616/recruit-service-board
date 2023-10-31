FROM node:16

WORKDIR /app

COPY ./recruit-info-service-front/package.json ./recruit-info-service-front/package-lock.json ./
RUN npm install

COPY ./recruit-info-service-front/ . 

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]