FROM node:18-bullseye
WORKDIR "/app"
EXPOSE 80

COPY . .
RUN chown -R node: .
USER node
RUN yarn install && yarn build

CMD [ "sh", "-c", "yarn install && yarn start --host 0.0.0.0 --port 80 --cors" ]
