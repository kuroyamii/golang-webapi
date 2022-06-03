<!-- ALL-CONTRIBUTORS-BADGE:START -->

[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)

<!-- ALL-CONTRIBUTORS-BADGE:END -->

# Overview

## Prerequisites

- Have installed MySQL
- Have installed Go
- Have installed Database Migrations `https://github.com/golang-migrate/migrate`
- Have learned the basics of Go

## Get Started

- First of all, you must make an empty database in your MySQL
- Create a `.env` file like below (change if it's needed):
  ```bash
  SERVER_ADDRESS=127.0.0.1:8080
  DB_ADDRESS=localhost:3309
  DB_USERNAME=root
  DB_PASSWORD=
  DB_NAME=inugami_db
  WHITELISTED_URLS=http://localhost:3000
  ```
- Run database migration. Example:</br></br>
  To run a migration

```bash
migrate -source file://directory-to-migration-folder -database "database-connection-query" up <number>

//example
migrate -source file://./db/migration -database "mysql://root:@tcp(localhost:3307)/inugami_db" up
```

To rollback a migration

```bash
migrate -source file://directory-to-migration-folder -database "database-connection-query" down <number>

//example
migrate -source file://./db/migration -database "mysql://root:@tcp(localhost:3307)/inugami_db" down 1
```

To create a new migration file

```bash
migrate create -ext sql -dir migrations <migration-name>

//example
migrate create -ext sql -dir migrations create_tables
```

- Run main.go in cmd/main

## Contributors ✨

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center">
        <a href="https://github.com/kuroyamii">
            <img src="https://avatars.githubusercontent.com/u/76874550?v=4?s=100" width="100px;" alt=""/>
            <br />
            <sub>
                <b>
                Gede Gery Sastrawan
                </b>
            </sub>
        </a>
        <br />
        <a href="https://github.com/kuroyamii/golang-webapi/commits?author=kuroyamii" title="Code">💻</a>
        <a href="https://github.com/kuroyamii/golang-webapi/commits?author=kuroyamii" title="Documentation">📖</a>
        <a href="#infra-kuroyamii" title="Infrastructure">🚇</a>
        <a href="https://github.com/kuroyamii/golang-webapi/pulls?q=is%3Apr+reviewed-by%3Akuroyamii" title="Reviewed Pull Requests">👀</a>
    </td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
