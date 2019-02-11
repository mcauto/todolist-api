#!/bin/bash
# TODO: timeout option
while ! docker exec $1 mysql --user=$2 --password=$3 -e "SELECT 1" >/dev/null 2>&1; do
    sleep 1;
	printf "."
done
printf " complete!\n"