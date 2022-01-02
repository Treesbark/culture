# docker build -f Dockerfile -t culture .
# docker run -p 8089:8089 -t culture  
# docker tag culture treesbark/culture:simple
# docker push treesbark/culture:simple  


# Build Stage

FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-culture"
LABEL REPO="https://github.com/Treesbark/culture"

ENV PROJPATH=/go/src/github.com/Treesbark/culture

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/Treesbark/culture
WORKDIR /go/src/github.com/Treesbark/culture

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/Treesbark/culture"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/culture/bin

WORKDIR /opt/culture/bin

COPY --from=build-stage /go/src/github.com/Treesbark/culture/bin/culture /opt/culture/bin/
RUN chmod +x /opt/culture/bin/culture

# Create appuser
RUN adduser -D -g '' culture
USER culture

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/culture/bin/culture"]