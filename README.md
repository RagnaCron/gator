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

# Clone this repo

```
git clone git@github.com/ragnacron/gator.git
```

# Setup to run `gator`

To run `gator` there are some requirements that must be met.

- Golang has to be installed
- Postgres has to be installed, and depending on the OS, certain configs have to be done
- Install Goose

It is assumed that there is a valid golang installation available.

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

*(Linux only)* Update postgres password:

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

To run the goose migration first test the connection to the postgres server:

- macOS (no password, your username): `postgres://<username>:@localhost:5432/gator`
- Linux (password set before, postgres user): `postgres://postgres:postgres@localhost:5432/gator`

To test the connection, in my case `username` == `ragnacron` (macOS):

```
psql "postgres://ragnacron:@localhost:5432/gator"
```

This should connect you to the `gator` database directly. If it is working, move on.

Time to run the migration.

Change the working directory (it is assumed that you are in the gator repo):

```
cd sql/schema
```

Migration time:

```
goose postgres <connection_string> up

# example:
# goose postgres "postgres://ragnacron:@localhost:5432/gator" up
```

# Install gator

In the base directory of the `gator` repository run:

```
go install
```

Change to your home directory and create `.gatorconfig.json` file. The `<connection_string>` needs an extra to work 
as we de not want the application code to try using SSL locally.

```
protocol://username:password@host:port/database?sslmode=disable
```

In my case (macOS) the `.gatorconfig.json` file has the following content:

```
{
    "db_url": "postgres://ragnacron:@localhost:5432/gator?sslmode=disable"
}
```

# gator commands


