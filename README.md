

This project aims to offer a set of tools for designing APIs on multiple protocols (HTTP/REST, gRPC)
and generate base bootstrap server and client code for the API in various languages, for various frameworks.

It follows the philosophy of design first: you first describe API at an abstract level, in a DSL,
then implement the service.

In order to achieve language interop, the API spec needs to be written in a language agnostic DSL.
Tools would be provided to support the custom DSL: parser, IDE plugin, reference code generator.
Other generators can be written for different languages and frameworks. Such generators will start with
the JSON representation of the DSL, and generate server/client code with various level of support
for the DSL features.

These should be the steps:
1. create the API design
  ```hcl
    service MyService {
        method AddNumbers {
            payload {
                field first Integer {
                    description = "first operand"
                    required = true
                }
                field second Integer {
                    description = "second operand"
                    required = true
                }
            }
            response = Integer
        }
    }
  ```
2. Generate code for the desired language (golang, python, scala) and desired framework
3. Write implementation code using the generated classes/interfaces.

This repo will contain:
- The DSL grammar and parser (golang). The parser will export a JSON representation of the AST for language interop.
- IDE language server for the DSL (golang)
- Goa generator: a tool to generate the server bootstrap code using Goa framework. The tool will read the JSON AST,
  generate Goa AST and run goa gen to produce the bootstrap code.
- Maybe other generators - for other languages, other frameworks. Or they can be hosted in different repos.
