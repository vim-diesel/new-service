# Build the Go Binary.
FROM golang:1.21 as build_sales-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /service

# Build the service binary.
WORKDIR /service/app/services/sales-api
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Build the admin binary.
WORKDIR /service/app/tooling/admin
RUN go build


# Run the Go Binary in Alpine.
FROM alpine:3.18
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S sales && \
    adduser -u 1000 -h /service -G sales -S sales
COPY --from=build_sales-api --chown=sales:sales /service/app/services/sales-api/sales-api /service/sales-api
COPY --from=build_sales-api --chown=sales:sales /service/.env /service/.env
COPY --from=build_sales-api --chown=sales:sales /service/key.pem /service/key.pem

EXPOSE 3000

WORKDIR /service
USER sales
CMD ["./sales-api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.title="sales-api" \
    org.opencontainers.image.authors="Ian Clark <email>" \
    org.opencontainers.image.source="https://github.com/vim-diesel/new-service/app/sales-api" \
    org.opencontainers.image.revision="${BUILD_REF}" \
    org.opencontainers.image.vendor="Vim Diesel"