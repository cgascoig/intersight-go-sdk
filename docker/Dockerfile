FROM ubuntu:focal

RUN apt-get update; \
    DEBIAN_FRONTEND=noninteractive \
    TZ=Australia/Melbourne \
    apt-get install -y openjdk-11-jre-headless python3 git make pipenv wget golang; \
    wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.3.0/openapi-generator-cli-5.3.0.jar -O ~/openapi-generator-cli.jar

RUN go get golang.org/x/tools/cmd/goimports


COPY run-update.sh /run-update.sh

CMD [ "/run-update.sh" ]