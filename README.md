String escape parser
===

It can be useful to decompose strings into quoted argument arrays similiar to how you expect
command line args to be processed. This string escape parser will decompose a string into its 
subsequent segments preserving quoted sections as single strings and escaped quotes within them.

For example:

```
foo bar biz
```
becomes
```
[]string { "foo", "bar", "biz" }
```

And

```
foo bar "biz \"baz\" boo"
```
will be

```
[]string { "foo", "bar", "biz \"baz\" boo" }
``` 
