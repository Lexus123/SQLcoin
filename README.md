# SQLcoin
Tool writen in Go that parses the bitcoin blockchain to MySQL.
## How to run
In order to run it, run the following command:

    go run cmd/main.go -blockfile /path/to/sqlcoin/testfiles/blk00004.dat
## Roadmap
- ~~Add extra startup parameter for database credentials~~ :heavy_check_mark:
- Code cleanup
- Comment all the things!
- Write tests
- Write benchmarks
- ~~Convert ScriptPubKey to address~~ :heavy_check_mark:
- ~~Create database schema~~ :heavy_check_mark:
    - ~~Write service that prepares data for SQL insert~~ :heavy_check_mark:
    - ~~Make MySQL connection~~ :heavy_check_mark:
## Future add ons
Might distribute the following in different repos
- UTXO age distribution tool
- Wallet identifier tool
