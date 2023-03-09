FROM centos

#WORKDIR /httpserver

ADD module2/httpserver /httpserver

EXPOSE 8000

ENTRYPOINT ["/httpserver"]
