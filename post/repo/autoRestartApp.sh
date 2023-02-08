# 一个简单的自动重启 Java 进程的脚本

pwd=`pwd`
retries=3
restartScript=restart.sh
appName=test.jar
sleepSec=3

function restart() {
    $pwd/$restartScript
}

function jpsCount() {
    jps -l | grep $appName | grep -v grep | wc -l
}

echo start auto restart app
echo ''

# do restart
restart

echo restart done
echo start check

# do check

for ((i = 0; i <= $retries; i++))
do
    if [ `jpsCount` -lt 1 ]; then
        echo restart failed
        sleep $sleepSec
        echo try restart app for $i
        restart
    else
        echo app restrt success
        break
    fi
done

echo exist restart....