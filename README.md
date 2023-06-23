# stencil
Stencil renders templated text with variables. It supports:

- [Go templates](https://golang.org/pkg/text/template/)
- [Mustache templates](https://mustache.github.io/)
- Colon templates: `:name` -> `Bob` renders the string "Hi, :name!" as "Hi, Bob!"

```
❯ ./stencil -template examples/template.gotemplate -variables name=Bob,age=35 -type gotemplate
Name: |Bob|
Age: |35|


❯ ./stencil -template examples/template.mustache -variables name=Bob,age=35 -type mustache
Name: |Bob|
Age: |35|


❯ ./stencil -template examples/template.colon -variables name=Bob,age=35 -type colon
Name: |Bob|
Age: |35|

```
