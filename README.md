# SQLcoin
Tool writen in Go that parses the bitcoin blockchain to MySQL.
## How to run
Make sure you've executed the SQL create statement found [HERE](sqldiagram/createsqlcoin3.sql).

In order to run the script, run the following command:

    go run cmd/main.go -dbcreds=user:password -dbname=sqlcoin -blockfile=/path/to/sqlcoin/testfiles/blk00004.dat
## Roadmap

**Needs performance upgrade. Database inserts batching.**

- ~~Add extra startup parameter for database credentials~~ :heavy_check_mark:
- Code cleanup
- Comment all the things!
- Write tests
- Write benchmarks
- ~~Convert ScriptPubKey to address~~ :heavy_check_mark:
- ~~Create database schema~~ :heavy_check_mark:
    - ~~Write service that prepares data for SQL insert~~ :heavy_check_mark:
    - ~~Make MySQL connection~~ :heavy_check_mark:
    - Switch to PostgreSQL
    - Change schema to [THIS](https://github.com/Blockchair/Blockchair.Support/blob/master/SQL.md#-database-schema)
    - Refactor database inserts
## Future add ons
Might distribute the following in different repos
- UTXO and coin age distribution
- Wallet identifier
- How much bitcoin of tx X is used in tx Y
