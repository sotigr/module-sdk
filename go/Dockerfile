FROM golang:1.21

# Installing NodeJs to use nodemon
RUN apt-get update && apt-get install curl -y
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - 
RUN apt-get install -y nodejs  

# Installing nodemon
RUN npm i -g nodemon

WORKDIR /src 

CMD [ "nodemon" ] 