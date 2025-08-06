DIR=~/go/github.com/jdetok/nightly-bball-etl
LOGD=z_log
# cd into this dir
cd $DIR
# cd ~/go/github.com/jdetok/nightly-bball-etl

# run nightly etl
go run ./pgins

# TODO: open the log file (most recent in z_log) and append to it
LOGF=$(ls z_log -t | head -n 1)
LOGDF=$LOGD/$LOGF

echo "attempting to run sp_nightly_call() from call.sql..." | tee -a $LOGDF
# call procedures: change container name as needed
docker exec -i devpg psql -U postgres -d bball < ./call.sql 2>&1 | tee -a $LOGDF
# docker exec -i devpg psql -U postgres -d bball < ./call.sql

echo "finished running sp_nightly_call()" | tee -a $LOGDF
echo "script complete at $(date)" | tee -a $LOGDF