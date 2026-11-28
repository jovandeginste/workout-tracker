#!/usr/bin/env bash
#
# Reset the workout-tracker database to a known test configuration.
#
# The database file is replaced with a fresh one. The app creates the admin
# user on first start; the users and workouts from scripts/reset-testdata.json
# are then created through the web and API endpoints:
#
#   - users are registered through /user/register and activated by the admin
#   - profiles (language, theme, timezone) are set through /user/profile
#   - workouts are uploaded through /api/v1/import/generic
#   - metadata (name, type, private) is set through the workouts API
#   - share links are created through the web endpoint
#
# Usage: scripts/reset-testdata.sh [--yes] [--no-build] [--keep-running]
#
# Environment:
#   WT_DB       database path (default: ./data/database.db)
#   WT_BIND     address to bind the temporary server on (default: 127.0.0.1:8080)
#   WT_URL      base URL used by the script (default: http://$WT_BIND)
#   WT_BIN      server binary (default: ./tmp/workout-tracker)
#   WT_MANIFEST manifest file (default: ./scripts/reset-testdata.json)
#
set -euo pipefail

ROOT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)

DB_PATH=${WT_DB:-"$ROOT_DIR/data/database.db"}
BIND=${WT_BIND:-"127.0.0.1:8080"}
BASE_URL=${WT_URL:-"http://$BIND"}
BASE_URL=${BASE_URL%/}
BIN=${WT_BIN:-"$ROOT_DIR/tmp/workout-tracker"}
MANIFEST=${WT_MANIFEST:-"$ROOT_DIR/scripts/reset-testdata.json"}

LOG_FILE="$ROOT_DIR/tmp/reset-testdata-server.log"
BUILD_LOG="$ROOT_DIR/tmp/reset-testdata-build.log"
SUMMARY_FILE="$ROOT_DIR/tmp/reset-testdata-summary.txt"

ASSUME_YES=false
BUILD=true
KEEP_SERVER=false

usage() {
	awk 'NR > 2 && /^set -euo pipefail$/ { exit } NR > 2 { sub(/^# ?/, ""); print }' "${BASH_SOURCE[0]}"
}

fail() {
	echo "error: $*" >&2
	exit 1
}

set_profile() {
	local jar=$1 username=$2 language=$3 theme=$4 timezone=$5
	local status

	status=$(curl -sS -o /dev/null -w '%{http_code}' -b "$jar" \
		--data-urlencode "language=$language" \
		--data-urlencode "theme=$theme" \
		--data-urlencode "timezone=$timezone" \
		--data-urlencode "auto_import_directory=" \
		--data "totals_show=running" \
		"$BASE_URL/user/profile")
	[[ $status == 302 ]] || fail "updating the profile of $username failed (HTTP $status)"
}

while [[ $# -gt 0 ]]; do
	case "$1" in
	-y | --yes)
		ASSUME_YES=true
		;;
	--no-build)
		BUILD=false
		;;
	--keep-running)
		KEEP_SERVER=true
		;;
	-h | --help)
		usage
		exit 0
		;;
	*)
		echo "unknown option: $1" >&2
		usage >&2
		exit 1
		;;
	esac
	shift
done

for cmd in curl jq lsof; do
	command -v "$cmd" >/dev/null || fail "'$cmd' is required"
done

[[ -f $MANIFEST ]] || fail "manifest not found: $MANIFEST"
jq -e '
	.admin.username and .admin.password and .admin.language and .admin.theme and .admin.timezone and
	(.users | length > 0) and
	(.users | all(.language and .theme and .timezone)) and
	(.workouts | length > 0)
' "$MANIFEST" >/dev/null || fail "invalid manifest: $MANIFEST"

mkdir -p "$(dirname "$DB_PATH")" "$ROOT_DIR/tmp"

if [[ -f $DB_PATH ]] && lsof "$DB_PATH" >/dev/null 2>&1; then
	echo "The database is in use by another process:" >&2
	lsof "$DB_PATH" >&2
	fail "stop the running server first"
fi

if curl -fsS --max-time 1 "$BASE_URL/health" >/dev/null 2>&1; then
	fail "something is already answering on $BASE_URL; stop it first"
fi

if ! $ASSUME_YES; then
	echo "This will replace $DB_PATH with a freshly seeded database."
	echo "The current database is kept as a timestamped backup."
	read -r -p "Continue? [y/N] " answer
	[[ $answer =~ ^[Yy]$ ]] || exit 1
fi

if $BUILD; then
	echo "Building server..."
	if ! make -C "$ROOT_DIR" build-server >"$BUILD_LOG" 2>&1; then
		tail -n 50 "$BUILD_LOG" >&2
		fail "failed to build the server (see $BUILD_LOG)"
	fi
fi

[[ -x $BIN ]] || fail "server binary not found or not executable: $BIN"

if [[ -f $DB_PATH ]]; then
	backup="${DB_PATH}.bak-$(date +%Y%m%d-%H%M%S)"
	mv "$DB_PATH" "$backup"
	echo "Previous database moved to $backup"
fi
rm -f "$DB_PATH-wal" "$DB_PATH-shm" "$DB_PATH-journal"

WORK_DIR=$(mktemp -d "${TMPDIR:-/tmp}/wt-reset-testdata.XXXXXX")
SERVER_PID=""

cleanup() {
	if [[ -n $SERVER_PID ]] && ! $KEEP_SERVER; then
		kill "$SERVER_PID" 2>/dev/null || true
		wait "$SERVER_PID" 2>/dev/null || true
	fi

	rm -rf "$WORK_DIR"
}

trap cleanup EXIT

echo "Starting temporary server on $BIND..."
WT_BIND="$BIND" \
	WT_DSN="$DB_PATH" \
	WT_DATABASE_DRIVER=sqlite \
	WT_JWT_ENCRYPTION_KEY="reset-testdata" \
	WT_REGISTRATION_DISABLED=false \
	WT_WEB_ROOT=/ \
	WT_WORKER_DELAY_SECONDS=3600 \
	WT_OFFLINE=true \
	nohup "$BIN" >"$LOG_FILE" 2>&1 &
SERVER_PID=$!

ready=false
for _ in $(seq 1 120); do
	if curl -fsS --max-time 1 "$BASE_URL/health" >/dev/null 2>&1; then
		ready=true
		break
	fi

	if ! kill -0 "$SERVER_PID" 2>/dev/null; then
		break
	fi

	sleep 0.5
done

if ! $ready; then
	tail -n 30 "$LOG_FILE" >&2
	fail "the temporary server did not become ready (see $LOG_FILE)"
fi

ADMIN_USERNAME=$(jq -r '.admin.username' "$MANIFEST")
ADMIN_PASSWORD=$(jq -r '.admin.password' "$MANIFEST")
ADMIN_LANGUAGE=$(jq -r '.admin.language' "$MANIFEST")
ADMIN_THEME=$(jq -r '.admin.theme' "$MANIFEST")
ADMIN_TIMEZONE=$(jq -r '.admin.timezone' "$MANIFEST")
ADMIN_JAR="$WORK_DIR/admin.cookies"

status=$(curl -sS -o "$WORK_DIR/login.html" -w '%{http_code}' \
	-c "$ADMIN_JAR" \
	--data-urlencode "username=$ADMIN_USERNAME" \
	--data-urlencode "password=$ADMIN_PASSWORD" \
	"$BASE_URL/user/signin")
[[ $status == 302 ]] || fail "admin login failed (HTTP $status)"

set_profile "$ADMIN_JAR" "$ADMIN_USERNAME" "$ADMIN_LANGUAGE" "$ADMIN_THEME" "$ADMIN_TIMEZONE"

whoami=$(curl -sS -b "$ADMIN_JAR" "$BASE_URL/api/v1/whoami")
jq -e --arg username "$ADMIN_USERNAME" --arg language "$ADMIN_LANGUAGE" --arg theme "$ADMIN_THEME" --arg timezone "$ADMIN_TIMEZONE" \
	'.results.username == $username and .results.active == true and .results.admin == true and
	 .results.profile.language == $language and .results.profile.theme == $theme and .results.profile.timezone == $timezone' \
	<<<"$whoami" >/dev/null || fail "could not verify the admin profile: $whoami"

echo "Creating users..."
while IFS=$'\t' read -r username name password language theme timezone; do
	[[ $username =~ ^[A-Za-z0-9._@-]+$ ]] || fail "invalid username in manifest: $username"

	status=$(curl -sS -o "$WORK_DIR/register-$username.html" -w '%{http_code}' \
		--data-urlencode "username=$username" \
		--data-urlencode "name=$name" \
		--data-urlencode "password=$password" \
		"$BASE_URL/user/register")
	[[ $status == 302 ]] || fail "registration of $username failed (HTTP $status)"

	admin_page=$(curl -sS -b "$ADMIN_JAR" "$BASE_URL/admin")
	user_id=$(grep -oE "<a href=\"[^\"]*/users/[0-9]+\">${username}</a>" <<<"$admin_page" |
		grep -oE '[0-9]+' | head -n 1 || true)
	[[ -n $user_id ]] || fail "could not find the id of user $username"

	status=$(curl -sS -o /dev/null -w '%{http_code}' \
		-b "$ADMIN_JAR" \
		--data-urlencode "username=$username" \
		--data-urlencode "name=$name" \
		--data "active=on" \
		"$BASE_URL/admin/users/$user_id")
	[[ $status == 302 ]] || fail "activation of $username failed (HTTP $status)"

	user_jar="$WORK_DIR/$username.cookies"
	status=$(curl -sS -o "$WORK_DIR/login-$username.html" -w '%{http_code}' \
		-c "$user_jar" \
		--data-urlencode "username=$username" \
		--data-urlencode "password=$password" \
		"$BASE_URL/user/signin")
	[[ $status == 302 ]] || fail "login of $username failed (HTTP $status)"

	set_profile "$user_jar" "$username" "$language" "$theme" "$timezone"

	whoami=$(curl -sS -b "$user_jar" "$BASE_URL/api/v1/whoami")
	jq -e --arg username "$username" --arg language "$language" --arg theme "$theme" --arg timezone "$timezone" \
		'.results.username == $username and .results.active == true and .results.admin == false and
		 .results.profile.language == $language and .results.profile.theme == $theme and .results.profile.timezone == $timezone' \
		<<<"$whoami" >/dev/null || fail "could not verify user $username: $whoami"

	echo "  $username ($name, $language, $theme, $timezone)"
done < <(jq -r '.users[] | [.username, .name, .password, .language, .theme, .timezone] | @tsv' "$MANIFEST")

total=$(jq '.workouts | length' "$MANIFEST")
echo "Uploading $total workouts..."

index=0
declare -A counts=()
declare -A private_counts=()
declare -A share_counts=()
share_lines=()

while IFS=$'\t' read -r username file name type private share; do
	index=$((index + 1))
	source_file="$ROOT_DIR/$file"
	[[ -f $source_file ]] || fail "workout file not found: $source_file"

	user_jar="$WORK_DIR/$username.cookies"
	[[ -f $user_jar ]] || fail "no session for user $username"

	encoded_name=$(jq -rn --arg value "$(basename "$file")" '$value | @uri')
	response=$(curl -sS -b "$user_jar" -X POST \
		--data-binary "@$source_file" \
		"$BASE_URL/api/v1/import/generic?name=$encoded_name")

	jq -e '.error == null and (.results | type == "array") and (.results | length == 1)' \
		<<<"$response" >/dev/null || fail "upload of $file failed: $response"

	workout_id=$(jq -r '.results[0].ID' <<<"$response")

	patch=$(jq -cn --arg name "$name" --arg type "$type" --argjson private "$private" \
		'{name: $name, type: $type, private: $private}')
	response=$(curl -sS -b "$user_jar" -X PATCH -H "Content-Type: application/json" \
		--data "$patch" "$BASE_URL/api/v1/workouts/$workout_id")
	jq -e --arg name "$name" --arg type "$type" --argjson private "$private" \
		'.error == null and .results.name == $name and .results.type == $type and .results.private == $private' \
		<<<"$response" >/dev/null || fail "updating $file failed: $response"

	share_url=""
	if [[ $share == true ]]; then
		status=$(curl -sS -o /dev/null -w '%{http_code}' -b "$user_jar" \
			-X POST "$BASE_URL/workouts/$workout_id/share")
		[[ $status == 302 ]] || fail "creating a share link for $file failed (HTTP $status)"

		uuid=$(curl -sS -b "$user_jar" "$BASE_URL/api/v1/workouts/$workout_id" |
			jq -r '.results.publicUUID // empty')
		[[ -n $uuid ]] || fail "no share link found for $file"

		share_url="$BASE_URL/share/$uuid"
		share_lines+=("$username  $name  $share_url")
	fi

	counts[$username]=$((${counts[$username]:-0} + 1))
	if [[ $private == true ]]; then
		private_counts[$username]=$((${private_counts[$username]:-0} + 1))
	fi
	if [[ $share == true ]]; then
		share_counts[$username]=$((${share_counts[$username]:-0} + 1))
	fi

	printf '  [%2d/%2d] %-4s %-42.42s' "$index" "$total" "$username" "$name"
	if [[ -n $share_url ]]; then
		printf ' ok (shared)\n'
	else
		printf ' ok\n'
	fi
done < <(jq -r '.workouts[] | [.user, .file, .name, .type, (.private | tostring), (.share | tostring)] | @tsv' "$MANIFEST")

{
	echo "Workout Tracker test data"
	echo
	echo "Admin: $ADMIN_USERNAME / $ADMIN_PASSWORD [$ADMIN_LANGUAGE, $ADMIN_THEME, $ADMIN_TIMEZONE]"
	while IFS=$'\t' read -r username password language theme timezone; do
		echo "User:  $username / $password [$language, $theme, $timezone] (${counts[$username]:-0} workouts, ${private_counts[$username]:-0} private, ${share_counts[$username]:-0} shared)"
	done < <(jq -r '.users[] | [.username, .password, .language, .theme, .timezone] | @tsv' "$MANIFEST")
	echo
	echo "Workouts: $total"
	echo
	if ((${#share_lines[@]} > 0)); then
		echo "Share links:"
		printf '  %s\n' "${share_lines[@]}"
	fi
} | tee "$SUMMARY_FILE"

echo
echo "Summary written to $SUMMARY_FILE"

if $KEEP_SERVER; then
	echo "The server keeps running on $BASE_URL (pid $SERVER_PID, log $LOG_FILE)."
else
	echo "The temporary server has been stopped; start the app as usual."
fi
