# new-service

This is built using code from [ardanlabs/service](https://github.com/ardanlabs/service)

## Description

This project has been modified to use [Neon.tech](https://neon.tech/) database, as this is just a hobby project we don't need to host our own database, or run the service in Kubernetes.

Instead it will be run using a cheaper alternative such as [Fly.io](https://fly.io) or Google Cloud Run.

This project is for learning and demonstration purposes only.

What follows are some general ideas and philosophies that drive the development of this service.

## Design Philosophies

Prototype driven development.

Data Oriented design. 

Integrity, Simplicity, Readability, and finally Performance(if necessary).

Readability is about not hiding things. Make them obvious.
Simplicity is about hiding complexity. Always start with readability and refactor into simplicity if you can find it

Don't add complexity until you need to

Precision. :)

If you don't understand that data, you don't understand the problem. If you don't understand the problem, get some code down and work your way to understanding the data/problem.

Prototype driven development is where we learn what the data is and the transformations are. Then we engineer that back into our main product.

The average human being cannot maintain more than 5 things in their head at once. 3 is more realistic and optimal for any average person. This project has 5 layers of depth.

No code generation. Don't hide your SQL and database code

No ORMs. Writing plain SQL allows us to fully debug and maintain.

Make things easy to understand and debug/maintain, not easy to do

Packaging allows us to have firewalls between the different domains of our program.

We write packages that __provide__, not packages that contain. For example, any packages that are named "util, helper" are containment packages and are avoided.

Similarly, we don't share models or types across packages/domains. So there are no files or packages named "models" or "types". App layer types stay in the app layer, and Business layer types stay in the business layer.

Every package is its own API/module. Every line of code you write either Allocates, Reads, or Writes memory (integers). Every function you write does a data transformation. Every API is a function.

Purpose of type system is that it allows for Input and Output through an API. So every API has its own type system, and allows us to maintain strict firewalls. Our APIs do not share types.

This maintaining of strict firewalls between APIs helps to avoid cascading breaking code changes. This is not a monolithic project. This is Go, not Java. We want firewalls between the different parts of our program.

Every line of code is an integer read/write. Every function is a data transformation. 

## Load Testing on Fly.io with a Neon.tech Database

Our fly.io config allows for 1 instance and a hard limit of 25 concurrent requests. Our logs are saying it takes ~40ms for a request to be completed once it hits the server. 

The distribution of the load: 

![load distribution](https://i.imgur.com/eGLBToV.png)


## App layers


[![Watch the video](https://img.youtube.com/vi/-FtCTW2rVFM/hqdefault.jpg)](https://www.youtube.com/embed/-FtCTW2rVFM)

Our code is organized in layers. The lower the layer the stricter the policy.

It is structured so that each layer can only import code from below it.

For example, the business layer can import from foundation layer, itself, and the vendor folder. It cannot import from the app layer. (Ideally business packages should not import each other, but sometimes they have to)

Foundation layer only imports from the vendor folder (dependencies).

### App layer

Our interface to the outside world. This is the layer where data validation takes place. Any layer below takes advantage of the Go type system and compiler to ensure data is already validated.

This is where our main binaries are stored. Such as Services, Tooling (like cli tools), scratch programs.

App layer is startup, shutdown, external input/output, and web handlers. 

It has its own data models for data going in and out, and does not share models with the business layer (maintaining strict firewalls). Our lower level code can change without breaking our app layer. 

### Business Layer

Packages specific to solving the business problem. Lower level stuff, such as database access.

This is not shareable across projects (same with our app layer), since no two repos should be solving the same problem. But these packages can be reused by different app layers programs (for example you might have services and cli tools using the same business APIs).

When you are building a business API, you might have multiple app types consuming, so don't build specifically for web calls or command line (for example, don't hide things in the context).

### Foundation layer

These are packages not specific to the business problem. They could be 3rd party, have their own repo, or be a standard library for our project. 

Here is where the strictest policies are set. There is no logging at this level, so that future developers can use a logger of their choice. So any packages that use a logger must be built in another layer.

## Logging

	- What is the purpose of the logs you are building
		○ Know the purpose otherwise they get messy
		○ Logging everything as insurance policy:
			§ Logging is expensive
		○ Logging in prod should be same as dev
			§ No logging levels
		○ Two biggest:
			§ 1 : ability to maintain, manage, and debug
			§ 2 : store data in logs
				□ If you cant write data to log your service must stop
				□ Downfall of storing data in logs
High ratio of signal to noise

## Convenience packages


		○ Do not build abstraction layers over dependencies
		○ Lots of work for something you will probably not need
			§ "I might want to replace this database, I will build a layer on top to switch out databases later" -> probably not going to happen
		○ Convenience package is not an abstraction layer:
		○ Average developer can maintain 10,000 lines of code
			§ Abstraction layer eats into this number
			§ Just use the dependencies API's and if you do need to rip it out later you can

## Packaging

	- Packages should have clear purpose
		○ Exist within scope of domain
		○ Don't want packages that "contain" code
			§ You can do it a bit, at app layer absolutely, business layer maybe, foundation layer never
			§ "utils, helpers, common" ... are containment packages
		○ Do not define package of common types
			§ "models", "types"
			§ All packages will depend on it
				□ Change one type and you have to change all other packages
		○ This is not a monolithic project, this is go not java
		○ We need firewalls between different parts of our package
			§ Package of common types destroys the firewall
		○ Every package should define its own data models for input and output
		○ 	Considering a package:
			□ Doesn't "contain" code
			□ We want to think about what the API is
			□ Domain of problem
			□ What layer it is
			□ How does it input/output

	- If a package provides or contains:
		○ We want provision, not containing
		○ If it doesn't make sense to name one of source code files as package name, it probably contains
		○ Ie. Common package does not have common.go, (it has stuff like marshall.go, this.go, that.go…) so you know it is containing code, not providing anything
		○ If you want to know what's going on, look in the file names after the package


## API input/output 

		○ Type system allows you to define input and output of an API
		○ You can accept input in one of two ways:
			§ Concrete type: Data is This (what it is)
				□ Most APIs should start out this way
			§ Interface type: Data does this (what it does)
				□ Polymorphic: a piece of code changes its behaviour depending on the concrete data it's operating on 
		○ APIs should only return concrete types
			§ Except error interface or empty interface
		○ Side note
			§ Generics give you ability to write polymorphic functions
			§ -> Code changes behavior depending on data its operating on
			§ Difference: When using interface type, you don't know what concrete data type is until runtime
				□ When using generics, determining what concrete type is at compile type
				□ Generics leverage go syntax 
				□ Empty interface uses Reflect package (an API)

			§ Don't develop with interfaces, discover them:
				□ Prototype driven development 
				□ data oriented design
			§ Every program we write is a data transformation
				□ Understand the data, understand the problem
			§ prototype driven development with our concrete types
				□ Once we have more than one concrete type being processed by a package, discover behaviours, then we refactor for interfaces


## Database

Migrations are done using [ArdanLabs/darwin](https://github.com/ardanlabs/darwin), a database schema evolution api for Go.

Sqlx is our abstraction layer, and pgx is our chosen driver.

## Comments

	- Comments are code
		○ Proper sentence structure, grammar, punctuation
		○ Code you're writing can produce docs for free if you follow guidelines, especially using staticcheck
		○ Any comments above package name will be part of overview in the go doc
			§ File that is named for the package ("logger.go") should be only file that has comments above package name
			§ If you have a large overview for a package, don't pollute the main file, make a "doc.go" and put the overview in there above package directive

## Configuration


		○ Only place config allowed to be read from is main.go
		○ All configuration should have a default value that at bare minimum works in dev environment
			§ Cloned repo should run on its own (unless they need a key, which should be clear where they go for a key and where to store it)
		○ Service should allow for --help
			§ Operator can see all configurable values, their default values, and how to override defaults
			§ Any config should be overridable by Env variable or commandline flags
		○ when app starts up, we should dump config we are using into the logs, and have ability to hide/mask any config that needs to maintain privacy
			§ Hiding credentials from logs is crucial
	- ArdanLabs conf/v3 package does all this for us.

## Error handling

## Logging errors

## HTTP Routing and Load shedding

## Handlers, Web Framework, Middleware

## 