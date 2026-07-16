#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
MANIFEST="${SCRIPT_DIR}/modules.tsv"
REMOTE="${REMOTE:-origin}"

usage() {
	cat <<'EOF'
Usage:
  module-release.sh list
  module-release.sh check <module> <version>
  module-release.sh tag <module> <version>
  module-release.sh push <module> <version>

Examples:
  module-release.sh check kit v3.0.0
  module-release.sh tag kit v3.0.0
  module-release.sh push kit v3.0.0
EOF
}

die() {
	printf 'error: %s\n' "$*" >&2
	exit 1
}

load_module() {
	local requested="$1"
	local name directory module_path tag_prefix

	while IFS=$'\t' read -r name directory module_path tag_prefix; do
		[[ -z "${name}" || "${name}" == \#* ]] && continue
		if [[ "${name}" == "${requested}" ]]; then
			MODULE_NAME="${name}"
			MODULE_DIR="${directory}"
			MODULE_PATH="${module_path}"
			TAG_PREFIX="${tag_prefix}"
			return 0
		fi
	done < "${MANIFEST}"

	die "unknown module '${requested}'; run '$0 list'"
}

validate_version() {
	local version="$1"
	[[ "${version}" =~ ^v3\.[0-9]+\.[0-9]+([+-][0-9A-Za-z.-]+)?$ ]] || \
		die "version must be a v3 semantic version, got '${version}'"
}

release_tag() {
	local version="$1"
	if [[ "${TAG_PREFIX}" == "-" ]]; then
		printf '%s\n' "${version}"
	else
		printf '%s/%s\n' "${TAG_PREFIX}" "${version}"
	fi
}

declared_module_path() {
	sed -n 's/^module[[:space:]]\+//p' "${REPO_ROOT}/${MODULE_DIR}/go.mod"
}

check_release() {
	local version="$1"
	local tag declared
	tag="$(release_tag "${version}")"

	[[ -f "${REPO_ROOT}/${MODULE_DIR}/go.mod" ]] || \
		die "missing ${MODULE_DIR}/go.mod"
	declared="$(declared_module_path)"
	[[ "${declared}" == "${MODULE_PATH}" ]] || \
		die "${MODULE_DIR}/go.mod declares '${declared}', expected '${MODULE_PATH}'"
	[[ -z "$(git -C "${REPO_ROOT}" status --porcelain)" ]] || \
		die "working tree is not clean; commit and push the release changes first"
	if git -C "${REPO_ROOT}" rev-parse --verify --quiet "refs/tags/${tag}" >/dev/null; then
		die "tag '${tag}' already exists locally"
	fi

	printf 'module:  %s\n' "${MODULE_NAME}"
	printf 'path:    %s\n' "${MODULE_PATH}"
	printf 'version: %s\n' "${version}"
	printf 'tag:     %s\n' "${tag}"
	printf 'commit:  %s\n' "$(git -C "${REPO_ROOT}" rev-parse HEAD)"
}

list_modules() {
	local name directory module_path tag_prefix pattern
	local -a tags

	printf '%-20s %-62s %s\n' "MODULE" "MODULE PATH" "LATEST TAG"
	while IFS=$'\t' read -r name directory module_path tag_prefix; do
		[[ -z "${name}" || "${name}" == \#* ]] && continue
		if [[ "${tag_prefix}" == "-" ]]; then
			pattern='v3.*'
		else
			pattern="${tag_prefix}/v3.*"
		fi
		mapfile -t tags < <(git -C "${REPO_ROOT}" tag --list "${pattern}" --sort=-version:refname)
		printf '%-20s %-62s %s\n' "${name}" "${module_path}" "${tags[0]:--}"
	done < "${MANIFEST}"
}

main() {
	local command="${1:-}"
	case "${command}" in
	list)
		[[ "$#" -eq 1 ]] || die "list does not accept arguments"
		list_modules
		;;
	check | tag | push)
		[[ "$#" -eq 3 ]] || {
			usage
			exit 1
		}
		load_module "$2"
		validate_version "$3"
		local tag
		tag="$(release_tag "$3")"
		case "${command}" in
		check)
			check_release "$3"
			;;
		tag)
			check_release "$3"
			git -C "${REPO_ROOT}" tag -a "${tag}" -m "release ${MODULE_PATH} ${3}"
			printf 'created tag %s\n' "${tag}"
			;;
		push)
			git -C "${REPO_ROOT}" rev-parse --verify --quiet "refs/tags/${tag}" >/dev/null || \
				die "tag '${tag}' does not exist locally"
			git -C "${REPO_ROOT}" push "${REMOTE}" "refs/tags/${tag}"
			;;
		esac
		;;
	-h | --help | help)
		usage
		;;
	*)
		usage
		exit 1
		;;
	esac
}

main "$@"
