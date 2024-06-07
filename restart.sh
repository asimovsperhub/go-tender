git pull origin master
ps -ef | grep 'tender'  | grep -v grep | awk '{print $2}' | xargs kill -9
go build -o tender
nohup ./tender &
