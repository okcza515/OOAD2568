### TODO

1. (waiting for CSV, see GH issues as ref) Transform data to SQLite.
2. Controller Interface (API).
3. Update docs

### Comamnd to dump database schema as sql file

```
sqlite3 test.db
.output database_schema.sql
.schema
.output stdout
```
