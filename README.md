# ecmd
Written in Go, ecmd is a command line tool for common Escenic tasks. It supports piped output to allow for chaining of commands.

## Examples
Print title and body for article 123456
ecmd article -f "title,body"

List contents in news-main of section 121:
ecmd area news-main 121

Print title and body for all articles in news-main area of section 121:
ecmd area news-main 121 | ecmd article -f "title,body"
