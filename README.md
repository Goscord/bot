# Example bot 
This is the example bot from Goscord. You can test it on [our Discord](https://goscord.dev/discord).

## Environment variable
You must create a `.env` file with the bot token variable (`BOT_TOKEN`).

## Start it
### Without Docker
You must have version 1.18 of GoLang installed on your machine. [See how to install GoLang](https://go.dev/dl/).

Then, run this command on your terminal:
```sh
go run main.go
```

### With Docker
This bot can also be launched via Docker so you don't need to install Golang on your machine. 

To start the bot with docker, here are the commands you need to do:
```sh
# Build the docker image:
docker build -t goscord-example .

# Then, run the docker image:
docker run --name goscord-bot -it --rm goscord-example
```