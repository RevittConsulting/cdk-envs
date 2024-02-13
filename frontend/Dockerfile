FROM node:latest AS builder

WORKDIR /frontend

COPY . .

RUN npm install

RUN npm run build

FROM nginx:alpine

COPY --from=builder /frontend/out /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]