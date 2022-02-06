FROM node:17.4.0-bullseye-slim
WORKDIR /app
COPY package.json /app
COPY web/package.json web/package-lock.json /app/web/
COPY server/package.json server/package-lock.json /app/server/
RUN npm run setup
COPY temp /app/server/public
COPY server/src /app/server/src
COPY prd.env /app/server/.env
CMD [ "npm", "run", "start-server" ]