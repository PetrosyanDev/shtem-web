# Erik Petrosyan ©
FROM debian:bullseye-slim as main

LABEL version="1.0"
LABEL author="ERIK_PETROSYAN"
LABEL description="Shtemaran.am web server"

RUN addgroup --system erik && adduser --system --ingroup erik erik
COPY --chown=erik:erik ./build/web/ /home/erik/cmd/web/
WORKDIR /home/erik/
RUN chmod +x ./cmd/web/app
USER erik
EXPOSE 2222