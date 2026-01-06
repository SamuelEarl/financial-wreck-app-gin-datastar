# Common Errors

---

## missing metadata for import

This specific error—"missing metadata for import"—usually means that while the code is there, the Go compiler's "cache" or the gopls (Go Language Server) is out of sync with your `go.mod` file. It’s a common hiccup when adding new dependencies like `templ`.

Here is the "reset" sequence to fix this.

### 1. The "Clean and Sync" Sequence
Run these commands in your project root to force Go to rebuild the metadata for your dependencies:

```bash
# 1. Download and prune dependencies
go mod tidy

# 2. Clean the build cache
go clean -modcache

# 3. Re-download everything
go mod download
```

### 2. Refresh your IDE (VS Code)

If you are using VS Code, the Language Server (`gopls`) sometimes gets "stuck" even after you run `go mod tidy`.

Open the Command Palette (`Ctrl+Shift+P` or `Cmd+Shift+P`).

Type and select: **"Go: Restart Language Server"**.

Wait 5 seconds. The red squiggles should disappear.

### 3. Check for Version Mismatches

If you have a package installed globally but a different version in your `go.mod`, the compiler can get confused. Check your `go.mod` file to see if the version is listed. Take the `templ` package as an example:

```go
// go.mod
require github.com/a-h/templ v0.3.833 // Ensure this line exists
```

If it's missing, run:

```bash
go get github.com/a-h/templ@latest
```

### 4. Why this happens (The "BrokenImport" Logic)

Go keeps a local cache of all downloaded packages. If a download was interrupted or if a dependency was updated but the "checksum" (go.sum) didn't update correctly, the compiler sees a folder for the package but no valid "metadata" to actually use it.

---
