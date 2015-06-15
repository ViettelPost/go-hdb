# go-hdb - A native Go HANA driver for Go's sql package

## Install

	go get github.com/SAP/go-hdb/driver
	
## Features

* Native Go implementation (no C libraries, CGO).
* Go <http://golang.org/pkg/database/sql> package compliant.
* Support of databse/sql/driver Execer and Queryer interface for parameter free statements and queries.
* Support of bulk inserts.
* Support of UTF-8 to / from CESU-8 encodings for HANA Unicode types.
* Build-in support of HANA decimals as Go rational numbers <http://golang.org/pkg/math/big>.
* Support of Large Object streaming.
* Support of Stored Procedures with table output parameters. 

## Dependencies

* <http://golang.org/x/text/transform>

## Todos

* Support of additional Go HANA datatype conversions.
* Additional Authentication Methods (actually only basic authentication is supported).
* Further performance improvements.
* Increase test code coverage and test documentation.