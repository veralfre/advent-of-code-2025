#!/bin/bash

# Generate README.md automatically based on available days

cat > README.md << 'EOF'
# advent-of-code-2025
Advent of Code 2025 in Golang!

EOF

# Add day list
echo "## Days" >> README.md
echo "" >> README.md

for day_dir in cmds/day*/; do
    if [ -d "$day_dir" ]; then
        day_num=$(basename "$day_dir" | sed 's/day//')
        # Try to find a title comment in main.go
        title=""
        if [ -f "$day_dir/main.go" ]; then
            # Look for a comment like "// Day 1: Secret Entrance" at the top of the file
            title=$(head -20 "$day_dir/main.go" | grep -E "^//.*Day.*:" | head -1 | sed 's|^// *||')
        fi
        
        if [ -z "$title" ]; then
            echo "- Day $day_num" >> README.md
        else
            echo "- $title" >> README.md
        fi
    fi
done

# Add how to run section
cat >> README.md << 'EOF'

## How to run

```shell
EOF

for day_dir in cmds/day*/; do
    if [ -d "$day_dir" ]; then
        day=$(basename "$day_dir")
        echo "go run ./cmds/$day -filename ./cmds/$day/input.txt" >> README.md
    fi
done

echo '```' >> README.md

echo "✅ README.md updated successfully!"
