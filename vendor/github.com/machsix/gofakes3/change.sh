#!/bin/bash

# Replace all instances of github.com/alist-org with github.com/machsix
# in .go, .mod, and .sum files

echo "Replacing github.com/alist-org with github.com/machsix ..."

# Search and replace in relevant files
find . -type f \( -name "*.go" -o -name "go.mod" -o -name "go.sum" \) -exec \
    sed -i 's|github.com/alist-org|github.com/machsix|g' {} +

echo "Done."

# Optional: tidy up Go modules
echo "Running go mod tidy ..."
go mod tidy

echo "All set."

