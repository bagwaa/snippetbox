The internal directory

It’s important to point out that the directory name internal carries a special meaning and behavior in Go: any packages
which live under this directory can only be imported by code inside the parent of the internal directory. In our case,
this means that any packages which live in internal can only be imported by code inside our snippetbox project
directory.

Or, looking at it the other way, this means that any packages under internal cannot be imported by code outside of our
project.

This is useful because it prevents other codebases from importing and relying on the (potentially unversioned and
unsupported) packages in our internal directory — even if the project code is publicly available somewhere like GitHub.