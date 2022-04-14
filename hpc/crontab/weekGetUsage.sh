#!/bin/bash

nowdate=`date +%Y-%m-%d`

endDate=`date -d "${nowdate} -1 days" +%Y-%m-%d`
startDate=`date -d "${nowdate} -7 days" +%Y-%m-%d`

# execute main
./build/main -dt_end=$endDate -dt_start=$startDate
