# syntax=docker/dockerfile:1

ARG NODE_VERSION=22.13.1
ARG GO_VERSION=1.23.4

FROM node:${NODE_VERSION}-alpine as node-base
WORKDIR /assets/vue

FROM node-base as deps

RUN --mount=type=bind,source=/assets/vue/package.json,target=package.json \
    --mount=type=bind,source=/assets/vue/package-lock.json,target=package-lock.json \
    --mount=type=cache,target=/root/.npm \
    npm ci --omit=dev

FROM deps as node-build

COPY assets/vue .
RUN npm run build

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS go-build
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

FROM alpine:latest AS final

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
    ca-certificates \
    tzdata \
    && \
    update-ca-certificates

RUN mkdir -p assets/vue

ENV NODE_ENV production

# Create a non-privileged user that the app will run under.
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

COPY assets/vue/package.json ./assets/vue
COPY --from=deps /assets/vue/node_modules ./assets/vue/node_modules
COPY --from=node-build /assets/vue/dist ./assets/vue/dist

COPY --from=go-build /bin/server /bin/
COPY .env .env

EXPOSE 8000

ENTRYPOINT [ "/bin/server" ]
