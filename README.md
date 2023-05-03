# grep
Command line grep like program in golang with regix feature in pattern <br />
-pattern=[regix_pattern] <br />
-file=file_to_be_searched <br />
Example case below <br />
Run Below command in terminal where grep.go program is or you can use build as well <br />
go run grep.go -pattern="(main|import)" -file=grep.go <br />
Output will look like this when we search the current program file grep.go <br />
package main <br />
import ( <br />
func main() {

