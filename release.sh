#!/bin/bash

# Usage function
usage() {
    echo "Usage: $0 -v <version> -t <title> -n <notes>"
    echo "  -v <version>     Version tag for the release (e.g., v1.0.0)"
    echo "  -t <title>       Title of the release"
    echo "  -n <notes>       Notes for the release"
    exit 1
}

# Parse flags
while getopts "v:t:n:" opt; do
    case $opt in
        v) version="$OPTARG" ;;
        t) release_title="$OPTARG" ;;
        n) release_notes="$OPTARG" ;;
        *) usage ;;
    esac
done

# Ensure all required flags are provided
if [ -z "$version" ] || [ -z "$release_title" ] || [ -z "$release_notes" ]; then
    usage
fi

# Define platforms as "GOOS/GOARCH"
platforms=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "windows/amd64"
)

# Array to store built file names for uploading
built_files=()

# Iterate over the platforms and build/compress binaries
for platform in "${platforms[@]}"; do
    GOOS=${platform%/*}
    GOARCH=${platform#*/}
    output_name="animal-hash-digest-${GOOS}-${GOARCH}"

    # Add .exe extension for Windows binaries
    if [ "$GOOS" == "windows" ]; then
        output_name="${output_name}.exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" || {
        echo "Failed to build for $GOOS/$GOARCH" >&2
        continue
    }

    echo "Compressing $output_name with UPX..."
    upx "$output_name" > /dev/null 2>&1 || {
        echo "Failed to compress $output_name with UPX" >&2
        continue
    }

    echo "Successfully built and compressed $output_name."
    built_files+=("$output_name")
done

# Create a GitHub release and upload binaries
if [ "${#built_files[@]}" -eq 0 ]; then
    echo "No binaries were successfully built. Skipping release creation."
    exit 1
fi

echo "Creating GitHub release ${version}..."
gh release create "$version" "${built_files[@]}" \
    --title "$release_title" \
    --notes "$release_notes" || {
    echo "Failed to create GitHub release." >&2
    exit 1
}

echo "Release ${version} created successfully with uploaded binaries."
