# stencil
Stencil renders templated text with variables. It supports:

- [Go templates](https://golang.org/pkg/text/template/)
- [Mustache templates](https://mustache.github.io/)
- Colon templates: `:name` -> `Bob` renders the string "Hi, :name!" as "Hi, Bob!"

```
❯ ./stencil --help
Stencil command: Convert templated text using variables

Usage: ./stencil [OPTIONS]
  -f, --file string        path to the template file
  -t, --type string        type of template to use (default "mustache")
  -v, --variables string   comma-separated list of key=value pairs
pflag: help requested

❯ ./stencil -f examples/template.mustache --variables name=Bob,age=35
Name: |Bob|
Age: |35|

❯ ./stencil -f examples/template.mustache --variables name=Bob,age=35 --type mustache
Name: |Bob|
Age: |35|

❯ ./stencil -f examples/template.gotemplate --variables name=Bob,age=35 --type go
Name: |Bob|
Age: |35|

❯ ./stencil -f examples/template.colon --variables name=Bob,age=35 --type colon
Name: |Bob|
Age: |35|

```

## Install

```
gh repo clone willf/stencil
cd stencil
go install ./cmd/stencil

```
