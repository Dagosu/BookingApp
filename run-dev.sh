set -o allexport
for ENV_FILE in \
    .booking-api.env \
;do
    if [ -f $ENV_FILE ]
    then
        source $ENV_FILE
    fi
done
set +o allexport

go run .
