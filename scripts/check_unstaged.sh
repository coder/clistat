#!/bin/bash

set -euo pipefail

FILES="$(git ls-files --other --modified --exclude-standard)"
if [[ "$FILES" != "" ]]; then
	mapfile -t files <<<"$FILES"

  echo
	echo "The following files contain unstaged changes:"
  echo
	for file in "${files[@]}"; do
		echo "  - $file"
	done

	echo
	echo "These are the changes:"
	echo
	for file in "${files[@]}"; do
		git --no-pager diff "$file" 1>&2
	done

	echo
	echo "ERROR: Unstaged changes, see above for details."
  exit 1
fi