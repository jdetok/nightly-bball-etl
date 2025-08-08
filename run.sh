DIR=/home/jdeto/go/github.com/jdetok/nightly-bball-etl
LOGD=$DIR/z_log
# cd into this dir
cd $DIR

# run nightly etl
export PATH=$PATH:/usr/local/go/bin
go run ./pgins >> gtest.log 2>&1

# TODO: open the log file (most recent in z_log) and append to it
LOGF=$LOGD/$(ls $LOGD -t | head -n 1)
# LOGDF=$LOGD/$LOGF

echo "attempting to run sp_nightly_call() from call.sql at $(date)..." | tee -a $LOGF
# call procedures: change container name as needed
docker exec -i pgbball psql -U postgres -d bball < ./call.sql 2>&1 | tee -a $LOGF
# docker exec -i devpg psql -U postgres -d bball < ./call.sql

echo "finished running sp_nightly_call()" | tee -a $LOGF
echo "script complete at $(date)" | tee -a $LOGF