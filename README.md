# SQLcoin
Tool writen in Go that parses the bitcoin blockchain to MySQL.
## How to run
In order to run it, run the following command:

    go run cmd/main.go -blockfile /path/to/sqlcoin/testfiles/blk00004.dat
## Roadmap
- Add extra startup parameter for database credentials
- Code cleanup
- Comment all the things!
- Write tests
- Write benchmarks
- Convert ScriptPubKey to address
- Create database schema
    - Write service that prepares data for SQL insert
    - Make MySQL connection
## Future add ons
Might distribute the following in different repos
- UTXO age distribution tool
- Wallet identifier tool
