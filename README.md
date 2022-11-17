# sisco - Lightweight Service Discovery
<!-- vim-markdown-toc GFM -->

* [Configuration](#configuration)
* [Integrated Help](#integrated-help)
* [Database Setup](#database-setup)
* [Starting the Server](#starting-the-server)
* [Areas, Services & Tags](#areas-services--tags)
* [REST API Endpoints](#rest-api-endpoints)
    * [Authentication & Authorization](#authentication--authorization)
    * [Register an Area](#register-an-area)
    * [Register a Service](#register-a-service)
    * [Query a Service](#query-a-service)
    * [List Areas](#list-areas)
    * [List Services](#list-services)
        * [List Services in Area](#list-services-in-area)
        * [List Services with Tag](#list-services-with-tag)
    * [List Tags](#list-tags)
    * [Delete Area](#delete-area)
    * [Delete Service](#delete-service)
    * [Delete Tag](#delete-tag)

<!-- vim-markdown-toc -->

`sisco` is a lightweight server to provide the discovery of services in a network. Using the REST API or gRPC interface, one can register and query services. Services can be also update and deleted. Access to the REST API or gRPC interface requires an authentication. Normal users are only allowed to query services, admin users can also add, modify and delete them. LDAP or Active Directory users and groups are used to get the credentials. The service data is stored in a database, which might be either a Postgres or a MySQL
database.

## Configuration

The configuration for `sisco` is stored in a file called `.sisco.yaml` in the home directory of the user. Use the command line option `-c` or `--config` to specify an alternative configuration file.

Below listed is a sample configuration file:

```yaml
debug: true
ginReleaseMode: false
port: 9999
gRPCPort: 8888
dbType: "postgres"
dbHost: "localhost"
dbPort: 5432
dbName: "<your-db-name>"
dbUser: "<your-db-user>"
dbPassword: "<your-db-password>"
dbSSLMode: "disable"
ldapURL: "ldap://localhost:3893"
ldapBaseDN: "DC=example,DC=com"
ldapBindDN: "CN=JohnDoe,OU=users,DC=example,DC=com"
ldapBindPassword: "<your-ldap-password>"
ldapFilterUsersDN: "(&(objectClass=posixAccount)(memberOf=OU=users,OU=groups,DC=example,DC=com)(uid={user}))"
ldapFilterAdminsDN: "(&(objectClass=posixAccount)(memberOf=OU=admins,OU=groups,DC=example,DC=com)(uid={user}))"
# Active Directory
# ldapFilterUsersDN: "(&(objectClass=person)(memberOf=OU=Users,OU=groups,DC=example,DC=com)(sAMAccountName={user}))"
# ldapFilterAdminsDN: "(&(objectClass=person)(memberOf=OU=Admins,OU=groups,DC=example,DC=com)(sAMAccountName={user}))"
tokenValidInSeconds: 1800
```

## Integrated Help

To show the integrated help just call `sisco` with the `help` command:

```shell
$ sisco --help
sisco is a small and lightweight server providing the possibility to register services and
to query for them.

Usage:
  sisco [command]

Available Commands:
  admin       Administrate sisco
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List components
  login       Login to sisco
  migrate     Support database migration
  serve       Start server
  version     Print the version of sisco

Flags:
  -c, --config string   config file (default is $HOME/.sisco.yaml)
  -d, --debug           enable debug output
  -h, --help            help for sisco
  -p, --pretty          enable pretty output

Use "sisco [command] --help" for more information about a command.
```

## Database Setup

As soon as your configuration is ready, the initial database setup for `sisco` is pretty easy, just execute the following command:

```shell
sisco migrate apply
```

Depending on your configured database type `sisco` is applying the database migrations stored in the database-specific directory in the `migrations` folder.

## Starting the Server

The server can be started using the `serve` command:

```shell
sisco serve
```

## Areas, Services & Tags

**Areas** allow you to define multiple **services** with the same name. Think of it as some sort of *prefix*. **Tags** can be added to **services**, allowing you to group them for example by functionality.

## REST API Endpoints

All REST API endpoints are currently starting with `/api/v1`.

### Authentication & Authorization

Authentication and authorization is done by posting the following JSON data to the REST API endpoint `/api/v1/login`:

Example `login.json`:

```json
{
    "User": "JohnDoe",
    "Password": "MyPassword"
}
```

Example `cURL` call:

```shell
$ curl -X POST --data @login.json localhost:9999/api/v1/login

{"token":"5bc9b49d41ef3477848fd56f8d6eac8e507331898de5fe14ff4bcd86381183d8"}
```

The REST API call is answered with a bearer token which you have to use in any other call to the `sisco` REST API. Depending on your group membership this might grant you also administrative access to the REST API.

### Register an Area

To register a new **area** the following JSON data needs to be posted to the REST API endpoint `/api/v1/admin/register/area/<area-name>`:

Example `register-area.json` file:

```json
{
    "description": "Area description"
}
```

Example `cURL` call:

```shell
$ curl -X POST -H "Bearer: <token>" --data @register-area.json localhost:9999/api/v1/register/area/<area-name>

{"area":{"id":42,"name":"<area-name>","description":"Area description","edges":{}}}
```

### Register a Service

To register a new **service** in an already existing **area** the following JSON data needs to be posted to the REST API endpoint `/api/v1/admin/register/service/<service-name>/in/<area-name>`:

Example `register-service.json` file:

```json
{
    "description": "Service description",
    "protocol": "Service protocol",
    "host": "Service host",
    "port": "Service port",
    "tags": ["foo", "v1"]
}
```

Example `cURL` call:

```shell
$ curl -X POST -H "Bearer: <token>" --data @register-area.json localhost:9999/api/v1/register/service/<service-name>/in/<area-name>

{"area":"<area-name>","service":{"id":42,"name":"<service-name>","description":"Service description","protocol":"Service protocol","host":"Service host","port":"Service port","edges":{}}}%
```

### Query a Service

To query for a service the REST API endpoint `/api/v1/get/service/<service-name>/in/<area-name>` is available:

Example `cURL` call:

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/get/service/<service-name>/in/<area-name>

{"id":42,"name":"<service-name>","description":"Service description","protocol":"Service protocol","host":"Service host","port":"Service port","edges":{}}%
```

### List Areas

To list all known **areas** the REST API endpoint `/api/v1/list/areas` is available:

Example `cURL` call (with prettied output):

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/list/areas

[
  {
    "id": 1,
    "name": "foo",
    "description": "foo description",
    "edges": {
      "services": [
        {
          "id": 1,
          "name": "alice",
          "description": "Alice description",
          "protocol": "Alice protocol",
          "host": "Alice host",
          "port": "Alice port",
          "edges": {}
        },
        {
          "id": 2,
          "name": "bob",
          "description": "Bob description",
          "protocol": "Bob protocol",
          "host": "Bob host",
          "port": "Bob port",
          "edges": {}
        }
      ]
    }
  },
  {
    "id": 2,
    "name": "bar",
    "description": "bar description",
    "edges": {}
  }
]
```

### List Services

#### List Services in Area

To list all known **services** in a specific **area** the REST API endpoint `/api/v1/list/services/in/<area-name>` is available:

Example `cURL` call (with prettied output):

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/list/services/in/<area-name>

[
  {
    "id": 1,
    "name": "alice",
    "description": "Alice description",
    "protocol": "Alice protocol",
    "host": "Alice host",
    "port": "Alice port",
    "edges": {}
  },
  {
    "id": 2,
    "name": "bob",
    "description": "Bob description",
    "protocol": "Bob protocol",
    "host": "Bob host",
    "port": "Bob port",
    "edges": {
      "tags": [
        {
          "id": 1,
          "name": "foo",
          "edges": {}
        },
        {
          "id": 2,
          "name": "v1",
          "edges": {}
        }
      ]
    }
  }
]
```

#### List Services with Tag

To list all known **services** with a specific **tag** the REST API endpoint `/api/v1/list/services/with/<tag-name>` is available:

Example `cURL` call (with prettied output):

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/list/services/in/<area-name>

[
  {
    "id": 2,
    "name": "bob",
    "description": "Bob description",
    "protocol": "Bob protocol",
    "host": "Bob host",
    "port": "Bob port",
    "edges": {
      "tags": [
        {
          "id": 1,
          "name": "foobar",
          "edges": {}
        },
        {
          "id": 2,
          "name": "v1",
          "edges": {}
        }
      ]
    }
  },
  {
    "id": 3,
    "name": "charlie",
    "description": "Charlie description",
    "protocol": "Charlie protocol",
    "host": "Charlie host",
    "port": "Charlie port",
    "edges": {
      "tags": [
        {
          "id": 1,
          "name": "barfoo",
          "edges": {}
        },
        {
          "id": 2,
          "name": "v1",
          "edges": {}
        }
      ]
    }
  }
]
```

### List Tags

To list all known **tags** the REST API endpoint `/api/v1/list/tags` is available:

Example `cURL` call (with prettied output):

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/list/tags

[
  {
    "id": 1,
    "name": "foo",
    "edges": {}
  },
  {
    "id": 2,
    "name": "v1",
    "edges": {}
  },
  {
    "id": 4,
    "name": "adas",
    "edges": {}
  },
  {
    "id": 5,
    "name": "v2",
    "edges": {}
  },
  {
    "id": 6,
    "name": "cpu",
    "edges": {}
  },
  {
    "id": 7,
    "name": "k8s",
    "edges": {}
  }
]
```

### Delete Area

To delete an area the REST API endpoint `/api/v1/admin/delete/area/<area>` is available:

***Note:*** An area can be only deleted if it does not contain any services.

Example `cURL` call (with prettied output):

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/admin/delete/area/foobar
```

### Delete Service

To delete a service the REST API endpoint `/api/v1/admin/delete/service/<service>/in/<area>` is available:

Example `cURL` call (with prettied output):

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/admin/delete/service/<service>/in/<area>
```

### Delete Tag

To delete a tag the REST API endpoint `/api/v1/admin/delete/tag/<tag>` is available:

Example `cURL` call (with prettied output):

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/admin/delete/tag/<tag>
```

## gRPC Endpoints

### Authentication & Authorization

Authentication and authorization is done by executing the following command:

```shell
sisco login <user> <password>
```
The gRPC call is answered with a bearer token which you have to use in any other gRPC call. The token might be provided either as command line option or by using an environment variable called `SISCO_TOKEN``. Depending on your group membership this might grant you also administrative access to the gRPC interface.
### Register an Area

To register a new **area** the following needs to be executed:

```shell
```

### Register a Service

To register a new **service** in an already existing **area** the following JSON data needs to be posted to the REST API endpoint `/api/v1/admin/register/service/<service-name>/in/<area-name>`:

Example `cURL` call:

```shell
```

### Query a Service

To query for a service the REST API endpoint `/api/v1/get/service/<service-name>/in/<area-name>` is available:

Example `cURL` call:

```shell
```

### List Areas

To list all known **areas** the REST API endpoint `/api/v1/list/areas` is available:

Example `cURL` call (with prettied output):

```shell
```

### List Services

#### List Services in Area

To list all known **services** in a specific **area** the REST API endpoint `/api/v1/list/services/in/<area-name>` is available:

Example `cURL` call (with prettied output):

```shell
```

#### List Services with Tag

To list all known **services** with a specific **tag** the REST API endpoint `/api/v1/list/services/with/<tag-name>` is available:

Example `cURL` call (with prettied output):

```shell
```

### List Tags

To list all known **tags** the REST API endpoint `/api/v1/list/tags` is available:

Example `cURL` call (with prettied output):

```shell
```

### Delete Area

To delete an area the REST API endpoint `/api/v1/admin/delete/area/<area>` is available:

***Note:*** An area can be only deleted if it does not contain any services.

Example `cURL` call (with prettied output):

```shell
```

### Delete Service

To delete a service the REST API endpoint `/api/v1/admin/delete/service/<service>/in/<area>` is available:

Example `cURL` call (with prettied output):

```shell
```

### Delete Tag

To delete a tag the REST API endpoint `/api/v1/admin/delete/tag/<tag>` is available:

Example `cURL` call (with prettied output):

```shell
```
