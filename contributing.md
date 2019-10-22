# Contributor Guide

## Short Name Limited Scope

Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. 


## Append

Integration tests should have the custom build directive `// +build integration`,
belong to a file called `.*_integration_test.go`, and a a package called `<packagename>_test`


## Right Build Directive Wrong Package

Integration tests should have the custom build directive `// +build integration`, belong to a file called `.*_integration_test.go`, and a a package called `<packagename>_test`

## Right Package Wrong Build Directive

Integration tests should have the custom build directive `// +build integration`, belong to a file called `.*_integration_test.go`, and a a package called `<packagename>_test`

## Right Package Wrong File

Integration tests should have the custom build directive `// +build integration`, belong to a file called `.*_integration_test.go`, and a a package called `<packagename>_test`

## Right File Wrong Build Directive

Integration tests should have the custom build directive `// +build integration`, belong to a file called `.*_integration_test.go`, and a a package called `<packagename>_test`

## Right File Wrong Package

Integration tests should have the custom build directive `// +build integration`, belong to a file called `.*_integration_test.go`, and a a package called `<packagename>_test`

## Camel case constants

See https://golang.org/doc/effective_go.html#mixed-caps. This applies even when it breaks conventions in other languages. For example an unexported constant is maxLength not MaxLength or MAX_LENGTH.


## Caught Generic Exceptions

Find caught generic exceptions.

In most cases, only specific exceptions should be caught. This allows different exceptions to be handled in different ways.
Catching generically will capture all other exception types lower in the hierarchy.
See the full hierarchy of the built in exceptions here: https://docs.python.org/3/library/exceptions.html#exception-hierarchy.


## Go Error Format

Error strings should not be capitalized (unless beginning with proper nouns 
or acronyms) or end with punctuation, since they are usually printed following
other context. That is, use fmt.Errorf("something bad") not fmt.Errorf("Something bad"),
so that log.Printf("Reading %s: %v", filename, err) formats without a spurious 
capital letter mid-message. This does not apply to logging, which is implicitly
line-oriented and not combined inside other messages.


## Context in Struct

Values of the context.Context type carry security credentials, tracing information, deadlines, and cancellation signals across API and process boundaries. Go programs pass Contexts explicitly along the entire function call chain from incoming RPCs and HTTP requests to outgoing requests.

Most functions that use a Context should accept it as their first parameter:

`func F(ctx context.Context, /* other arguments */) {}`

A function that is never request-specific may use context.Background(), but err on the side of passing a Context even if you think you don't need to. The default case is to pass a Context; only use context.Background() directly if you have a good reason why the alternative is a mistake.

Don't add a Context member to a struct type; instead add a ctx parameter to each method on that type that needs to pass it along. The one exception is for methods whose signature must match an interface in the standard library or in a third party library.

Don't create custom Context types or use interfaces other than Context in function signatures.

If you have application data to pass around, put it in a parameter, in the receiver, in globals, or, if it truly belongs there, in a Context value.

Contexts are immutable, so it's fine to pass the same ctx to multiple calls that share the same deadline, cancellation signal, credentials, parent trace, etc.


## Defer In Loop

Finds defer calls in loops.

Defers are only executed when a function returns. Using them inside loops is often a mistake since they will
build up and can be executed later than expected. The defers will never get called if the loop does not exit.


## Update Comment First Word as Subject

Doc comments work best as complete sentences, which allow a wide variety of automated presentations.
The first sentence should be a summary that starts with the name being declared.


## Redundant Defer Wraps

Finds redundant defer wrapping that could be simplified into a single line.

A single line defer is cleaner and easier to read than one wrapped in a function.


## Go Lint Rules

Find all go linter rules.


## Global Variable

Global variables used within functions are not visible in 
the functions signature. This complicates testing, reduces readability 
and increases the complexity of code.


## Shadowed Function Parameter

Avoid shadowing your function parameters to make your code easier to read.
See "https://gist.github.com/lavalamp/4bd23295a9f32706a48f"


## New Package Requires Test

All new packages must come with unit tests.


## test

The loop variable is reused for each iteration, so it is shared across all goroutines. We need to make sure that it is unique for each goroutine. One way to do that, is passing the loop variable as an argument to the closure in the goroutine.
Note: This tenet assumes that loop variables are not shadowed inside goroutine. We need ssa to work to find the right loop variables in that case.


## no Get at the start of Getters name

It's neither idiomatic nor necessary to put Get into the getter's name. If you have a field called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner. 


## JSON API kind compulsory metadata

Every object kind must provide a nested object field called 'metadata' that contains both a 'namespace' and a 'name' field.


## Local variables with limited scope are prefered to be short


Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. Prefer c to lineCount. Prefer i to sliceIndex.

The basic rule. the further from its declaration that a name is used, the more descriptive the name must be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (i, r). More unusual things and global variables need more descriptive names.


## No Custom Context

Values of the context.Context type carry security credentials, tracing information, deadlines, and cancellation signals across API and process boundaries. Go programs pass Contexts explicitly along the entire function call chain from incoming RPCs and HTTP requests to outgoing requests.

Most functions that use a Context should accept it as their first parameter:

func F(ctx context.Context, /* other arguments */) {}
A function that is never request-specific may use context.Background(), but err on the side of passing a Context even if you think you don't need to. The default case is to pass a Context; only use context.Background() directly if you have a good reason why the alternative is a mistake.

Don't add a Context member to a struct type; instead add a ctx parameter to each method on that type that needs to pass it along. The one exception is for methods whose signature must match an interface in the standard library or in a third party library.

Don't create custom Context types or use interfaces other than Context in function signatures.

If you have application data to pass around, put it in a parameter, in the receiver, in globals, or, if it truly belongs there, in a Context value.

Contexts are immutable, so it's fine to pass the same ctx to multiple calls that share the same deadline, cancellation signal, credentials, parent trace, etc.


## Comment First Word as Subject

Doc comments work best as complete sentences, which allow a wide variety of automated presentations.
The first sentence should be a summary that starts with the name being declared.


## Flags Have Underscores

Command-line flags should use dashes, not underscores. See "https://github.com/kubernetes/community/blob/master/contributors/guide/coding-conventions.md"


## Use Appropriate Verbs in Format

Prefer the most specific verb for your use. In other words, 
prefer to avoid %v when possible. However, %v is to be used 
when formatting bindings which might be nil and which do not 
already handle nil formatting. Notably, nil errors formatted 
as %s will render as "%!s()" while nil errors formatted as %v 
will render as "". Therefore, prefer %v when formatting errors 
which are not known to be non-nil.


## Function Arguments Inline Comments

A code reader encountering a function call should be able to 
intuit what all the arguments to the call represent. Whenever
it wouldn't be otherwise clear what the value used as an argument
represents (for example, from the variable's name if a variable
is used or from the type name if a struct literal is used), 
consider annotating it with an inline comment specifying the
respective parameter's name. Particularly, consider doing this
for literals of "basic" types (boolean, numeric, string types,
whether the type is predeclared or not) and for nil identifiers, 
as they are frequently not suggestive enough of what they represent.


## Debug Prints

Find debug print calls.

Debug print calls shouldn't be left in production code.
They make the program unnecessarily verbose and can output sensitive information.


## Context as First Argument

Values of the context.Context type carry security credentials, tracing information, 
deadlines, and cancellation signals across API and process boundaries. Go programs 
pass Contexts explicitly along the entire function call chain from incoming RPCs 
and HTTP requests to outgoing requests.

Most functions that use a Context should accept it as their first parameter.


## Avoid Meaningless Package Names

Avoid meaningless package names like util, common, misc, api, types, and interfaces. See http://golang.org/doc/effective_go.html#package-names and http://blog.golang.org/package-names for more.


## Println Format Strings

Find Print and Println using format strings.


## Todo Comments

Find comments containing TODO.

TODOs should be tracked in a place outside of the code to ensure they aren't forgotten.


## Unconvert

Identify unnecessary type conversions


## Unnecessary Parentheses

Avoid sunnecessary parentheses to make your code easier to read.


## Bool Arg

Find functions with one or more boolean arguments.


## Goto Statement

Find uses of goto.


## Unsafe Go Routine Variables

Example tenet that finds unsafe variables in goroutines.

## Empty Slice

Find empty slice assignments.


## Go Lint Rules

Find all go linter rules.


## Reallocated Slice

Find slices that may be subject to reallocation pointer changes. Designed to catch problems like this https://github.com/juju/juju/commit/8ff9d72ebc07c0f1d2f048e5d0486335e637b313

## Sprintf Error

Find instances of 'errors.New(fmt.Sprintf(...))'.


## Global Variable

Global variables used within functions are not visible in 
the functions signature. This complicates testing, reduces readability 
and increases the complexity of code.


## Nil Only Functions

Find functions that only return nil.


## Remove Break Statement

The break statement that is needed at the end of each case is provided automatically in Go

## NewConcat

Example tenet that finds all exported functions that do not have a corresponding test function.

## Init

Check that no inits functions are present in Go code.


## Mixed Marshalling

Find cases where a server marshals data into a client data type.

Separates client and server versions of a request by ensuring the server never unmarshalls to the client's version. Use the types of your own structs to create a project specific review.


## Avoid Annotations in Comments

Comments do not need extra formatting such as banners of stars. The generated output
may not even be presented in a fixed-width font, so don't depend on spacing for alignment—godoc, 
like gofmt, takes care of that. The comments are uninterpreted plain text, so HTML and other 
annotations such as _this_ will reproduce verbatim and should not be used. One adjustment godoc 
does do is to display indented text in a fixed-width font, suitable for program snippets. 
The package comment for the fmt package uses this to good effect.


## test

The loop variable is reused for each iteration, so it is shared across all goroutines. We need to make sure that it is unique for each goroutine. One way to do that, is passing the loop variable as an argument to the closure in the goroutine.
Note: This tenet assumes that loop variables are not shadowed inside goroutine. We need ssa to work to find the right loop variables in that case.


## Package Comment

Every package should have a package comment, a block comment preceding the package clause. 
For multi-file packages, the package comment only needs to be present in one file, and any one will do. 
The package comment should introduce the package and provide information relevant to the package as a 
whole. It will appear first on the godoc page and should set up the detailed documentation that follows.


## Single Method Interface Name

By convention, one-method interfaces are named by the method name plus an -er suffix 
or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

There are a number of such names and it's productive to honor them and the function names they capture. 
Read, Write, Close, Flush, String and so on have canonical signatures and meanings. To avoid confusion, 
don't give your method one of those names unless it has the same signature and meaning. Conversely, 
if your type implements a method with the same meaning as a method on a well-known type, give it the 
same name and signature; call your string-converter method String not ToString.


## Update Comment First Word as Subject

Doc comments work best as complete sentences, which allow a wide variety of automated presentations.
The first sentence should be a summary that starts with the name being declared.


## Defer Close File

Deferring a call to a function such as Close has two advantages. First, it guarantees that you will never forget to close the file, a mistake that's easy to make if you later edit the function to add a new return path. Second, it means that the close sits near the open, which is much clearer than placing it at the end of the function.
TODO if a function returns the open file, follow it and check if it is closed


## Good Package Name

It's helpful if everyone using the package can use the same name 
to refer to its contents, which implies that the package name should 
be good: short, concise, evocative. By convention, packages are 
given lower case, single-word names; there should be no need for 
underscores or mixedCaps. Err on the side of brevity, since everyone 
using your package will be typing that name. And don't worry about 
collisions a priori. The package name is only the default name for 
imports; it need not be unique across all source code, and in the 
rare case of a collision the importing package can choose a different 
name to use locally. In any case, confusion is rare because the file 
name in the import determines just which package is being used.


## no Get at the start of Getters name

It's neither idiomatic nor necessary to put Get into the getter's name. If you have a field called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner. 


## Unnecessary Else

When an if statement doesn't flow into the next statement—that is, the body ends in break, continue, goto, or return—the unnecessary else is omitted.


## Comment First Word as Subject

Doc comments work best as complete sentences, which allow a wide variety of automated presentations.
The first sentence should be a summary that starts with the name being declared.


## Initialize instance using composite literal

Sometimes the zero value isn't good enough and an initializing constructor is necessary. We can simplify the code using a composite literal, which is an expression that creates a new instance each time it is evaluated.


## Reuse the variable name in a type switch

If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. It's also idiomatic to reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case.


## Avoid Meaningless Package Names

Avoid meaningless package names like util, common, misc, api, types, and interfaces. See http://golang.org/doc/effective_go.html#package-names and http://blog.golang.org/package-names for more.


## Avoid Renaming Imports

Avoid renaming imports except to avoid a name collision; good package names
should not require renaming. In the event of collision, prefer to rename the
most local or project-specific import.


## Short Reader variables


Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. Prefer c to lineCount. Prefer i to sliceIndex.

The basic rule. the further from its declaration that a name is used, the more descriptive the name must be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (i, r). More unusual things and global variables need more descriptive names.


## Context as First Argument

Values of the context.Context type carry security credentials, tracing information, 
deadlines, and cancellation signals across API and process boundaries. Go programs 
pass Contexts explicitly along the entire function call chain from incoming RPCs 
and HTTP requests to outgoing requests.

Most functions that use a Context should accept it as their first parameter.


## Declaring Empty Slices

When declaring an empty slice, prefer

```go
var t []string
```

over

```go
t := []string{}
```

The former declares a nil slice value, while the latter is non-nil but zero-length. They are functionally equivalent-their `len` and `cap` are both zero `tbut the nil slice is the preferred style.

Note that there are limited circumstances where a non-nil but zero-length slice is preferred, such as when encoding JSON objects (a `nil` slice encodes to `null`, while `[]string{}` encodes to the JSON array `[]`).

When designing interfaces, avoid making a distinction between a nil slice and a non-nil, zero-length slice, as this can lead to subtle programming errors.

For more discussion about nil in Go see Francesc Campoy's talk [Understanding Nil](https://www.youtube.com/watch?v=ynoY2xz-F8s).


## No Custom Context

Values of the context.Context type carry security credentials, tracing information, deadlines, and cancellation signals across API and process boundaries. Go programs pass Contexts explicitly along the entire function call chain from incoming RPCs and HTTP requests to outgoing requests.

Most functions that use a Context should accept it as their first parameter:

func F(ctx context.Context, /* other arguments */) {}
A function that is never request-specific may use context.Background(), but err on the side of passing a Context even if you think you don't need to. The default case is to pass a Context; only use context.Background() directly if you have a good reason why the alternative is a mistake.

Don't add a Context member to a struct type; instead add a ctx parameter to each method on that type that needs to pass it along. The one exception is for methods whose signature must match an interface in the standard library or in a third party library.

Don't create custom Context types or use interfaces other than Context in function signatures.

If you have application data to pass around, put it in a parameter, in the receiver, in globals, or, if it truly belongs there, in a Context value.

Contexts are immutable, so it's fine to pass the same ctx to multiple calls that share the same deadline, cancellation signal, credentials, parent trace, etc.


## Method's receiver name

The name of a method's receiver should be a reflection of its identity; often a one or two letter 
abbreviation of its type suffices (such as "c" or "cl" for "Client"). Don't use generic names such as 
"me", "this" or "self", identifiers typical of object-oriented languages that gives the method a 
special meaning. In Go, the receiver of a method is just another parameter and therefore, should 
be named accordingly. The name need not be as descriptive as that of a method argument, as its 
role is obvious and serves no documentary purpose. It can be very short as it will appear on 
almost every line of every method of the type; familiarity admits brevity. 


## Short Name Limited Scope

Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. 


## Context in Struct

Values of the context.Context type carry security credentials, tracing information, deadlines, and cancellation signals across API and process boundaries. Go programs pass Contexts explicitly along the entire function call chain from incoming RPCs and HTTP requests to outgoing requests.

Most functions that use a Context should accept it as their first parameter:

`func F(ctx context.Context, /* other arguments */) {}`

A function that is never request-specific may use context.Background(), but err on the side of passing a Context even if you think you don't need to. The default case is to pass a Context; only use context.Background() directly if you have a good reason why the alternative is a mistake.

Don't add a Context member to a struct type; instead add a ctx parameter to each method on that type that needs to pass it along. The one exception is for methods whose signature must match an interface in the standard library or in a third party library.

Don't create custom Context types or use interfaces other than Context in function signatures.

If you have application data to pass around, put it in a parameter, in the receiver, in globals, or, if it truly belongs there, in a Context value.

Contexts are immutable, so it's fine to pass the same ctx to multiple calls that share the same deadline, cancellation signal, credentials, parent trace, etc.


## Do Not Discard Errors

Do not discard errors using _ variables. If a function returns an error, check it to make sure the function succeeded. Handle the error, return it, or, in truly exceptional situations, panic.


## Go Error Format

Error strings should not be capitalized (unless beginning with proper nouns 
or acronyms) or end with punctuation, since they are usually printed following
other context. That is, use fmt.Errorf("something bad") not fmt.Errorf("Something bad"),
so that log.Printf("Reading %s: %v", filename, err) formats without a spurious 
capital letter mid-message. This does not apply to logging, which is implicitly
line-oriented and not combined inside other messages.


## Camel case constants

See https://golang.org/doc/effective_go.html#mixed-caps. This applies even when it breaks conventions in other languages. For example an unexported constant is maxLength not MaxLength or MAX_LENGTH.


## Local variables with limited scope are prefered to be short


Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. Prefer c to lineCount. Prefer i to sliceIndex.

The basic rule. the further from its declaration that a name is used, the more descriptive the name must be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (i, r). More unusual things and global variables need more descriptive names.


## Use Crypto Rand

Do not use package math/rand to generate keys, even 
throwaway ones. Unseeded, the generator is completely predictable. 
Seeded with time.Nanoseconds(), there are just a few bits of entropy. 
Instead, use crypto/rand's Reader, and if you need text, print to 
hexadecimal or base64


