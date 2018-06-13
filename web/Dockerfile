FROM node as build-deps
ENV NPM_CONFIG_LOGLEVEL warn
WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . ./
RUN yarn build

FROM nginx
COPY --from=build-deps /usr/src/app/build /usr/share/nginx/html
