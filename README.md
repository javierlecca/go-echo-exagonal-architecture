# DOMAIN-DRIVEN DESIGN WITH GO
This is a project to explain hexagonal architecture in go

## What is the Domain?
To define domain-driven design we should first establish what we mean by domain in this context (and in development in general). The common dictionary definition of domain is: “A sphere of knowledge or activity.” Drilling down a bit from that, domain in the realm of software engineering commonly refers to the subject area on which the application is intended to apply. In other words, during application development, the domain is the “sphere of knowledge and activity around which the application logic revolves.”
Another common term used during software development is the domain layer or domain logic which may be better known to many developers as the business logic. The business logic of an application refers to the higher-level rules for how business logic interact with one another to create and modify modelled data.


## Communication flow

![myimage-alt-tag](https://miro.medium.com/max/810/1*b75xN3W9mQzta37pT-siRQ.png) 


## Dependencies

### GO version

- [GO v1.15](https://golang.org)

### Base Framework
- [Echo](https://github.com/labstack/echo/v4)

### DB Connection
- [Gorm](https://github.com/jinzhu/gorm)

### Other
- [Configor v1.2.0](https://github.com/jinzhu/configor)


```sh
$ git clone https://github.com/javierlecca/go-echo-exagonal-architecture.git
$ go run main.go
```