FROM golang:1.19-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o server cmd/coffeeshops-server/main.go

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim

RUN mkdir -p /etc/apt/keyrings/

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates \
    wget \
    gnupg

# Install Grafana Agent
RUN wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | tee /etc/apt/keyrings/grafana.gpg > /dev/null
RUN echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | tee /etc/apt/sources.list.d/grafana.list
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y grafana-agent-flow

# Remove apt lists to reduce image size
RUN rm -rf /var/lib/apt/lists/*

# Start the Grafana Agent
RUN --mount=type=secret,id=grafana-agent-flow_river,dst=/etc/secrets/grafana-agent-flow.river \
    cat /etc/secrets/grafana-agent-flow.river > /etc/grafana-agent-flow.river

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server
COPY start.sh start.sh

# Set the GIN_MODE environment variable to production.
ENV GIN_MODE=release

# Run the web service on container startup.
CMD ["./start.sh"]
