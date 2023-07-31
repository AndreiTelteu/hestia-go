FROM node:18-bullseye
WORKDIR "/app"
EXPOSE 80

COPY . .
RUN chown -R node: .
# RUN su node -c "yarn install"
# RUN su node -c "yarn build"

CMD [ "su", "node", "-c", "yarn install && yarn start --host 0.0.0.0 --port 80 --cors" ]