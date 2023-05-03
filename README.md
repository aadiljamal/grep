# grep
Command line grep like program in golang with regix feature in pattern 
-pattern=[regix_pattern]
-file=file_to_be_searched
Example case below
Run Below command in terminal where grep.go program is or you can use build as well
go run grep.go -pattern="(main|import)" -file=grep.go
Output will look like this when we search the current program file grep.go
package main
import (
func main() {

