FROM heroiclabs/nakama-pluginbuilder:3.21.1 AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend

COPY go.mod .
COPY *.go .
COPY configs/ ./configs/
COPY vendor/ vendor/
RUN ls -la /backend

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM registry.heroiclabs.com/heroiclabs/nakama:3.21.1

COPY --from=go-builder /backend/backend.so /nakama/data/modules/
COPY local.yml /nakama/data/
COPY configs/ /nakama/data/configs/