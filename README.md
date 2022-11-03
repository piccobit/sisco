# sisco - Lightweight Service Discovery

`sisco` is a lightweight tool to provide the discovery of services in a network. Using the REST API, one can register and query services. Services can be also update and deleted. Access to the REST API requires an authentication. Normal users are only allowed to query services, admin users can also add, modify and delete them. LDAP or Active Directory users and groups are used to get the credentials. The service data is stored in a database, which might be either a Postgres or MySQL
database.

## Configuration

The configuration for `sisco` is stored in a file called `.sisco.yaml` in the home directory of the user. Use the command line option `-c` or `--config` to specify an alternative configuration file.

## Database Setup

As soon as your configuration is ready, the initial database setup for `sisco` is pretty easy, just execute the following command:

```shell
sisco migrate apply
```

Depending on your configured database type `sisco` is applying the database migrations stored in the specified `migrations` folder.

## Authentication & Authorization

Authentication and authorization is done by posting the following JSON data to the REST API endpoint `/login`:

```json
{
    "User": "JohnDoe",
    "Password": "MyPassword"
}
```

The REST API call is answered with a bearer token which you have to use in any other call to the `sisco` REST API. Depending on your group membership this might grant you also administrative access to the REST API.

## Areas, Services & Tags

**Areas** allow you to define multiple **services** with the same name. Think of it as some sort of *prefix*. **Tags** can be added to **services**, allowing you to group them for example by functionality.

## Register an Area

To register a new **area** the following JSON data needs to be posted to the REST API endpoint `/register/area/<area-name>`:

```json
{
    "description": "Area description"
}

```

## Register a Service

To register a new **service** the following JSON data needs to be posted to the REST API endpoint `/register/service/<area-name>/<service-name>`:

```json
{
    "description": "Service description",
    "protocol": "Service protocol",
    "host": "Service host",
    "port": "Service port",
    "tags": ["foo", "v1"]
}

```

## Query a service

To query for a service
