# stencil
Stencil renders templated text with variables. It supports:

- [Go templates](https://golang.org/pkg/text/template/)
- [Mustache templates](https://mustache.github.io/)
- Colon templates: `:name` -> `Bob` renders the string "Hi, :name!" as "Hi, Bob!"

```
❯ ./stencil --help
Usage: stencil [OPTIONS]
Stencil command: Convert templated text using variables

Options:
  -f --file <file>   path to a template file (default: stdin)
  -g --go            use Go template syntax
  -m --mustache      use Mustache template syntax (default)
  -c --colon         use colon template syntax
  -h --help          print this help message
Other flags are passed as key=value pairs for use in the template

Stencil tries to be forgiving about whether keys get dashes or values have a =

❯ stencil -f examples/template.mustache name=Bob age=35
Name: |Bob|
Age: |35|

> stencil -f examples/template.mustache name="Bob Smith" age=35 --mustache
Name: |Bob Smith|
Age: |35|

❯ stencil -f examples/template.mustache name=Bob age=35 --mustache
Name: |Bob|
Age: |35|

❯ stencil -f examples/template.gotemplate name=Bob age=35 -g
Name: |Bob|
Age: |35|

❯ ./stencil -f examples/template.colon name=Bob age=35 --colon
Name: |Bob|
Age: |35|

```

## Install

```
gh repo clone willf/stencil
cd stencil
go install ./cmd/stencil

```
