FROM scratch

COPY devyunmu.com /
COPY templates /templates
COPY static /static

EXPOSE 8080
ENTRYPOINT ["/devyunmu.com"]
