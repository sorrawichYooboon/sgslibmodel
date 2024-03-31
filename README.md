# How to create shared library and model and use it in another project

### Step to create shared library and model

1. Create a new project (if you have dependencies, you can use go mod init)

```bash
go mod init <module name>
```

2. Create a new shared library (not create main.go)

```bash
package <shared library name>
```

3. Create a new model and function (PascalCase)

4. Push the shared library to the repository

5. Release the shared library with tag version

---

### Step to use the shared library and model that create by another project

1. Get the shared library and model

```bash
go get <path to the shared library>
```

or

```bash
go get <path to the shared library>@<tag version>
```

2. Import the shared library

```go
import "<path to the shared library>"
```

or

```go
import somemodule "<path to the shared library>/<some module>"

```

3. Use the shared library

```go
// Create a new instance of the shared library
sharedLibrary := shared.NewSharedLibrary()

// Use the shared library
sharedLibrary.DoSomething()
```

4. (optional) vendor the shared library for offline build and deterministic build

```bash
go mod vendor
```
