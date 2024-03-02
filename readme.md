# Go Lang Practice 🔧

<div align="center">
    <img src="./other/gocute.png" width="200px" />
</div>

```go
$ go mod init example.com/example -> Initializes module in current dir
$ go mod tidy -> Cleans & updates dependencies in go.mod
```

# Code Style: Camel Case or Snake Case?
<p>
In Go, there are some common naming conventions for variables, functions, and other identifiers: [1]

snake_case is commonly used for variable and function names. All lowercase letters separated by underscores. For example: function_name, variable_name.
CamelCase is used for type names (structs, interfaces, etc) and exported (public) identifiers. The first letter of each internal word is capitalized. For example: MyType, ExportedFunction.
leading underscores for private variables/functions - _privateVar. These are not exported.
Some key points:

snake_case is preferred over camelCase for non-exported identifiers by many Go programmers.
Exported identifiers use CamelCase so they are accessible from other packages.
The naming convention helps identify what kind of entity something is (variable, type, exported function) at a glance.
So in summary, snake_case is commonly used for local variables and functions, while CamelCase is used for exported/public types and identifiers that need to be accessible from other packages. Following common conventions helps make code more readable and maintainable. 
</p>
