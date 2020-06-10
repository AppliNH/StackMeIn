# easyStack or "StackMeIn" project

<div align="center">
<img src="https://i.kym-cdn.com/photos/images/original/001/465/006/fe5.gif">
</div>

# What is it

easyStack is a project I've started with the main objective to dive deeper into Go and Docker.
The goal of this project is to provide a GUI (through the WebApp) which allows you to graphically create a docker-compose file and to start it.

I wanna give this project an "OpenStack" aspect.

For this, I'm working with the [Docker Golang SDK](https://godoc.org/github.com/docker/docker/client).


# What is done, for now

- Back-End (Golang) :
    - [x] Accept POST requests in the aim to registrer a docker-compose file with a specific id.
    - [x] Accept GET requests to start a docker-compose configuration, that a call a "stack".
    - [x] Accepte DELETE requests to stop a stack.
    - [x] Store, save and access stack's configurations in a database ([FireGo](https://github.com/applinh/firego)).

# What has to be done

- Back-End (Golang) :
    - A more concise way to retrieve running stacks & containers, with the right name.
    - A better correspondence between containers' IDs and their names.


- Front-End (React) : 
    - Not planned yet.


# How does GoCompose work for now ?

If you `clone` , `cd stackmein/GoCompose`, and run `go run main.go`, somewhat of a REST API will run on **port 1997**.

## `/stack/{id}`
### GET

Will start the stack which has been registered with the given ID.

### DELETE

Will stop the stack which has been registered with the given ID.

## `/dockercompose`

### POST

Receives your docker-compose config as JSON in the body, and saves it in the FireGO database with a specific id.
It's now a stack.

Example of a payload that is handled : 

```json
{
	"version": "2.1",
	"services": {
		"debian": {
			"command":"tail -f /dev/null",
			"image": "debian",
			"ports": ["4000:4000"],
			"networks": ["main_network"]
		},
		"ubuntu": {
			"command":"tail -f /dev/null",
			"image": "ubuntu",
			"ports": ["3000:3000"],
			"networks": ["main_network"]
		}
	},
	"networks": {
		"main_network":""
	}
}
```

# How does FireGo database is setup for this ?

Stacks' configurations go in the `/stacks` resource.

As FireGo runs by default on port 5000, you can make a **GET** request to `localhost:5000/stacks` to see the registered stacks.

You'll find objects like this : 

```json
{
    "id":{
        "containers":[],
        "dockercompose": {
                "networks": {"main_network":""},
                "services":{
                    "debian":{
                        "command":"tail -f /dev/null","image":"debian",
                        "networks":["main_network"],"ports":["4000:4000"]
                    },
                    "ubuntu":{
                        "command":"tail -f /dev/null","image":"ubuntu",
                        "networks":["main_network"],"ports":["3000:3000"]
                    }
                },
                "version":"2.1"
            }
    }
}
```

# How do I run this ?

- `git clone` the project
- `cd` inside of it
- run `docker-compose build` && `docker-compose up`