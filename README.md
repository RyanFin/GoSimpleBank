# Simple Bank

### DBDiagram.io

[DBDiagram.io](https://dbdiagram.io/d/simple-bank-660a849437b7e33fd733241e)

## SQLC

- [SQLC Installation](https://docs.sqlc.dev/en/latest/overview/install.html)
  - run sqlc init to generate `sqlc.yaml` file
- [SQLC Getting Started](https://docs.sqlc.dev/en/v1.26.0/tutorials/getting-started-postgresql.html#schema-and-queries)
- [Setting up sqlc.yaml old](https://docs.sqlc.dev/en/v1.8.0/reference/config.html)
- [Setting up sqlc.yaml latest](https://docs.sqlc.dev/en/v1.26.0/reference/config.html)

- `transfer.sql` and `entry.sql` implementations:
  https://github.com/techschool/simplebank/blob/master/db/query/transfer.sql
  https://github.com/techschool/simplebank/blob/master/db/query/entry.sql

-[lib/pq is required for postgres unit testing process](https://github.com/lib/pq)
