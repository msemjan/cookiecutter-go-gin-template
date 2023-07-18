# Compile stage
FROM golang:1.17 AS build-env

ADD . /app
WORKDIR /app
RUN make build

# Final stage
FROM scratch

EXPOSE 8000
WORKDIR /
COPY --from=build-env /bin/{{ cookiecutter.binary_name }} /

CMD ["/{{ cookiecutter.binary_name }}"]

HEALTHCHECK --interval=60s --retries=5 --timeout=6s CMD curl --fail http://localhost:8000/api/v1/health || exit 1
