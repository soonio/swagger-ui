FROM node:22-alpine
WORKDIR /app
COPY . .
RUN npm install

ENV PORT=8934
ENV SECRET=123456

CMD ["node", "server.js"]

EXPOSE 8934