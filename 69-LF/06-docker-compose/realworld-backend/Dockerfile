FROM node:latest
WORKDIR /app
COPY . /app
RUN yarn install && yarn cache clean && yarn boot
CMD ["yarn", "start"]