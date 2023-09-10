FROM node:18-bullseye
COPY --from=oven/bun /usr/local/bin/bun /usr/local/bin/bun
WORKDIR "/app"
EXPOSE 80

COPY . .
RUN chown -R node: .
USER node
# RUN yarn install && yarn build

CMD [ "sh", "-c", "bun install && bun run start --host 0.0.0.0 --port 80 --cors" ]
