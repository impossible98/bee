# one
FROM impossible98/golang as one
WORKDIR /src/
COPY ./ ./
RUN make install \
    && make build
# two
FROM debian:10-slim
ARG name=bee
ENV PORT=7226
WORKDIR /app
COPY --from=one /src/dist/$name ./
EXPOSE $PORT
ENTRYPOINT ["/app/bee", "serve"]
