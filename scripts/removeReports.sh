#!/bin/bash

# Remove all json reports from the reports folder
# Usage: bash scripts/remove_reports.sh

for file in ./reports/*_report.json
do
  if [ -f "$file" ]; then
    rm "$file"
  fi
done
