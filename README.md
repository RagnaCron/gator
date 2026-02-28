# Gator

This repo follows the course on [boot.dev](https://boot.dev)

Here is a part of the introduction to the course on boot.dev:

# Welcome to the Blog Aggregator

We're going to build an RSS feed aggregator in Go! We'll call it "Gator", you know, because aggreGATOR 🐊. Anyhow, it's a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

# Setup to run `gator`

To run `gator` there are some requirements that must be met.

- Golang has to be installed
- Postgres has to be installed, and depending on the OS, certain configs have to be done
- Install Goose

It is assumed that there is a valid golang installation is available.

## Postgres configs

*macOS* with `brew`

```
brew install postgresql@16

```

*Linux/WSL (Debian)*

```
sudo apt update 
sudo apt install postgresql postgresql-contrib
```

Check if the installation was successful.

```
psql --version
```

*(Linux only)* UPdate postgres password:

```
sudo passwd postgres
```

Start the Postgres server in the background

- Mac: `brew services start postgresql@16`
- Linux: `sudo service postgresql start`

Check the connection to the server:

- Mac: `psql postgres`
- Linux: `sudo -u postgres psql`

Create the `gator` database:

```
CREATEL DATABASE gator;
```

Connect to the new database:

```
\c gator
```

*(Linux only)* Set the user passsword:

```
ALTER USER postgres PASSWORD 'postgres';
```

Enter `exit` to quit the `psql` shell.

## Installing Goose

Goose is a command line tool written in Go, it is usesed for migrations in SQL.

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```




