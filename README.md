# new-service

This is built using code from [ardanlabs/service](https://github.com/ardanlabs/service)

## Description

This project has been modified to use [Neon.tech](https://neon.tech/) database, as this is just a hobby project we don't need to host our own database, or run the service in Kubernetes.

Instead it will be run using a cheaper alternative such as [Fly.io](https://fly.io) or Google Cloud Run.

This project is for learning and demonstration purposes only.

What follows are some general ideas and philosophies that drive the development of this service.

## Design Philosophies

Prototype driven development.

Integrity, Simplicity, Readability, and finally Performance.

Readability is about not hiding things. Make them obvious.
Simplicity is about hiding complexity. Always start with readability and refactor into simplicity if you can find it

Don't add complexity until you need to

Precision. :)

No code generation. Don't hide your SQL and database code

No ORMs. Writing plain SQL allows us to fully debug and maintain.

Make things easy to understand and debug/maintain, not easy to do

Packaging allows us to have firewalls between the different domains of our program.

We write packages that __provide__, not packages that contain. For example, any packages that are named "util, helper" are containment packages and are avoided.

Similarly, we don't share models or types across packages/domains. So there are no files or packages named "models" or "types". App layer types stay in the app layer, and Business layer types stay in the business layer. 

Every package is its own API/module. Every line of code you write either Allocates, Reads, or Writes memory (integers). Every function you write does a data transformation. Every API is a function. 

Purpose of type system is that it allows for Input and Output through an API. So every API has its own type system, and allows us to maintain strict firewalls. Our APIs do not share types. 

This maintaining of strict firewalls between APIs helps to avoid cascading breaking code changes.
 
## App layers

![Is our app an ogre?](https://media.tenor.com/TXJmqbUeyO8AAAAC/shrek-ogres-have-layers.gif)

Our code is organized in layers. The lower in layer you are the stricter the policies are.

### App layer

Our interface to the outside world. This is the layer where data validation takes place. Any layer below takes advantage of the Go type system and compiler to ensure data is already validated.

### Business Layer

Migrations are done using [ArdanLabs/darwin](https://github.com/ardanlabs/darwin), a database schema evolution api for Go. 

Sqlx is our chosen database management package, and pgx is our chosen driver.

### Foundation layer