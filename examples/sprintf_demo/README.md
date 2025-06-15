# Sprintf Demo

This example demonstrates the new `fmt.Sprintf` functionality in gexe, which allows you to use Go's string formatting alongside variable expansion.

## Features Demonstrated

1. **Basic string formatting**: Using `%s`, `%d`, `%t` format verbs
2. **Multiple arguments**: Passing multiple values for formatting
3. **Combined functionality**: Using both sprintf and variable expansion together
4. **File operations**: Creating files with formatted paths
5. **Variable setting**: Setting variables with formatted values
6. **Backward compatibility**: Showing that unused args are ignored when no format verbs exist

## Run the Example

```bash
go run main.go
```

## Key Benefits

- **Cleaner code**: No need for manual string concatenation or `fmt.Sprintf` calls
- **Type safety**: Go's type checking for format arguments
- **Flexibility**: Combine with existing variable expansion features
- **Backward compatible**: Existing code continues to work unchanged