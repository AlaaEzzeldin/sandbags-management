FROM node:14.15.3-alpine as develop

WORKDIR /client

COPY package*.json ./

RUN yarn install
COPY . .

FROM develop as build
RUN yarn build

FROM nginx:stable-alpine as production

RUN mkdir /app
COPY frontend-nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /client/dist /app
#COPY --from=build /client/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

