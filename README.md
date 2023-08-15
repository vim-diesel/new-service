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

 • No code generation
  ○ Don't hide your SQL and database code
  ○ Make things easy to understand and debug/maintain
   § Not easy to do

## Database

Migrations are done using the darwin package from ArdanLabs. This prevents misuse of the sql files.

Sqlx is our chosen database management package, and pgx is our chosen driver.
