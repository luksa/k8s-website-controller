FROM scratch
LABEL org.opencontainers.image.authors="marko.luksa@gmail.com"
ADD website-controller /
ADD deployment-template.json /
ADD service-template.json /
CMD ["/website-controller"]
