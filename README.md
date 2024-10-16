# Gator 🐊
This is a CLI tool for Aggergating RSS feeds. It supports multiple users, feed following and storing posts.

# Setup
You will need to install [postgres](https://www.postgresql.org/download/) and [go](https://go.dev/doc/install).
After that you can install the Gator CLI from the root of this project with
```BASH
go install
```
Before running you need to make sure that the postgres service is online. In Linux or WSL you can achieve this with:
```BASH
sudo service postgresql start
```
The Gator CLI requires an existing database with name `gator`. You can create it via the psql shell with the following SQL statement:
```SQL
CREATE DATABASE gator;
```
You will also need to setup a config file located at `~/.gatorconfig.json`. A default config may look like this:
```JSON
{
    "db_url": "postgres://<user>:<password>@localhost:5432/gator?sslmode=disable",
    "current_user_name": "peter"
}
```

# CLI Usage
```bash
gator register <username>        # registers a new user
gator login <username>           # switches the current user
gator users                      # lists all users
gator feeds                      # lists all feeds
gator browse [<limit>]           # lists all posts of all feeds the current user follows
gator addfeed <name> <url>       # adds a new feed
gator follow <url>               # lets the current user follow an existing feed url
gator unfollow <url>             # lets the current user unfollow a followed feed url
gator following                  # lists all feeds that the current user is following

gator reset                      # resets the database

```
