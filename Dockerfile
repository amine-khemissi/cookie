FROM centos 
RUN yum install -y git wget
RUN mkdir -p /app/cookie
WORKDIR /app

RUN wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && tar -zxvf go1.12.7.linux-amd64.tar.gz 
RUN cd /usr/bin && ln -s /app/go/bin/go
ENV GOBIN=/app/go/bin 
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN cd /usr/bin && ln -s /app/go/bin/dep

WORKDIR /app/src/awesomeProject
COPY . /app/src/awesomeProject/ 
ENV GOPATH=/app/
RUN dep ensure
WORKDIR /app/src/awesomeProject/main
RUN go build   

CMD ["/app/src/awesomeProject/main/main"]

