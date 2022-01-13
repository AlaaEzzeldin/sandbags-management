FROM node:14.15.3-alpine as develop

WORKDIR /client

COPY package*.json ./

RUN yarn install
COPY . .

FROM develop as build
RUN yarn build

FROM nginx:stable-alpine as production

COPY --from=build /client/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

