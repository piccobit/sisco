# sisco - Lightweight Service Discovery
<!-- vim-markdown-toc GFM -->

* [Configuration](#configuration)
* [Database Setup](#database-setup)
* [Areas, Services & Tags](#areas-services--tags)
* [API Endpoints](#api-endpoints)
    * [Authentication & Authorization](#authentication--authorization)
    * [Register an Area](#register-an-area)
    * [Register a Service](#register-a-service)
    * [Query a Service](#query-a-service)
    * [List Areas](#list-areas)
    * [List Services](#list-services)
        * [List Services in Area](#list-services-in-area)
        * [List Services with Tag](#list-services-with-tag)
    * [List Tags](#list-tags)

<!-- vim-markdown-toc -->

`sisco` is a lightweight tool to provide the discovery of services in a network. Using the REST API, one can register and query services. Services can be also update and deleted. Access to the REST API requires an authentication. Normal users are only allowed to query services, admin users can also add, modify and delete them. LDAP or Active Directory users and groups are used to get the credentials. The service data is stored in a database, which might be either a Postgres or a MySQL
database.

## Configuration

The configuration for `sisco` is stored in a file called `.sisco.yaml` in the home directory of the user. Use the command line option `-c` or `--config` to specify an alternative configuration file.

Below listed is a sample configuration file:

```yaml
debug: true
ginReleaseMode: false
port: 9999
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

## Database Setup

As soon as your configuration is ready, the initial database setup for `sisco` is pretty easy, just execute the following command:

```shell
sisco migrate apply
```

Depending on your configured database type `sisco` is applying the database migrations stored in the database-specific directory in the `migrations` folder.

## Areas, Services & Tags

**Areas** allow you to define multiple **services** with the same name. Think of it as some sort of *prefix*. **Tags** can be added to **services**, allowing you to group them for example by functionality.

## API Endpoints

All API endpoints are currently starting with `/api/v1`.

### Authentication & Authorization

Authentication and authorization is done by posting the following JSON data to the API endpoint `/api/v1/login`:

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

The API call is answered with a bearer token which you have to use in any other call to the `sisco` API. Depending on your group membership this might grant you also administrative access to the API.

### Register an Area

To register a new **area** the following JSON data needs to be posted to the API endpoint `/api/v1/register/area/<area-name>`:

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

To register a new **service** in an already existin **area** the following JSON data needs to be posted to the API endpoint `/api/v1/register/service/<service-name>/in/<area-name>`:

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

To query for a service the API endpoint `/api/v1/get/service/<service-name>/in/<area-name>` is available:

Example `cURL` call:

```shell
$ curl -H "Bearer: <token>" localhost:9999/api/v1/get/service/<service-name>/in/<area-name>

{"id":42,"name":"<service-name>","description":"Service description","protocol":"Service protocol","host":"Service host","port":"Service port","edges":{}}%
```

### List Areas

To list all known **areas** the API endpoint `/api/v1/list/areas` is available:

Example `cURL` call (with prettyfied output):

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

To list all known **services** in a specifix **area** the API endpoint `/api/v1/list/services/in/<area-name>` is available:

Example `cURL` call (with prettyfied output):

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

To list all known **services** with a specific **tag** the API endpoint `/api/v1/list/services/with/<tag-name>` is available:

Example `cURL` call (with prettyfied output):

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

To list all known **tags** the API endpoint `/api/v1/list/tags` is available:

Example `cURL` call (with prettyfied output):

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

