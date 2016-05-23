#!/bin/bash

# Add this file to crontab to enable auto sync
# MUST call with absolute path.

#source ~/.bashrc # required in crontab

#SYNCDIR=$1

#SCM='svn --non-interactive'
SCM='git'
#URL="jd master:master"

# bypass possible invalid ssl cert in https git
export GIT_SSL_NO_VERIFY=1

BASENAME=`basename $0`
LOGFILE=/tmp/$BASENAME-`whoami`.log
#NOTIFYFLAG=/tmp/$BASENAME.notified
CURDIR=`dirname $0`
#SYNCDIR=$CURDIR/disk/
SYNCDIR=$CURDIR
SCMDIR=$SYNCDIR/../../

# create a signature file in sync disk for each host
\rm $SYNCDIR/iAm-* 2>/dev/null
touch $SYNCDIR/iAm-`hostname`

#echo $SYNCDIR >> /tmp/dbg.log

# test
#echo $CURDIR
#exit

XCOWSAY=/usr/games/xcowsay
MSG_TIMEOUT=2
#if [ -e $XCOWSAY ]; then
#    DISPMSG="DISPLAY=:0 $XCOWSAY --cow-size=small"
#else
#    # DISPMSG="export DISPLAY=:0 ; xmessage -timeout 2 -file -"
## doesn't work quit well under centos. giving it a dummy now.
#    DISPMSG="echo"
#fi

#DISPMSG="DISPLAY=:0 /usr/games/xcowsay --cow-size=small"

#date >| /tmp/autosync.log
#id >> /tmp/autosync.log

# kill possible previous svn instance and cleanup, in case svn hangs.
# WARN: side-effect is could kill an interactive svn session die, e.g. when commiting manually.
# DONE update this part for git.
# Another fix: don't simply kill process containing "git" in its name.
procpid=`ps -ef | \grep '[g]it ' | awk '{print $2}'`
if [ -n "$procpid" ]; then
    kill -9 $procpid &>/dev/null
fi
#svn cleanup $SCMDIR

#pwd | write `whoami`

#cd /home/yicfu.ws/baidu_bae/yicfudisk
#cd $SYNCDIR
cd $CURDIR/../
# (for git, still need to cd)

#pwd | write `whoami`

#$SCM up $SCMDIR 2>|$LOGFILE 1>/dev/null
$SCM pull $URL >|$LOGFILE 2>&1
#date >> $LOGFILE # test, fake an error message

if [ $? -ne 0 ]; then
#if [ -s $LOGFILE ]; then
    #if [ ! -f $NOTIFYFLAG ]; then
    #    $DISPMSG < $LOGFILE
    #    touch $NOTIFYFLAG
    #fi
    #DISPLAY=:0 /usr/games/xcowsay --cow-size=small < $LOGFILE

    #$DISPMSG < $LOGFILE
    ##if [ -x $XCOWSAY ]; then
    ##    DISPLAY=:0 $XCOWSAY -t $MSG_TIMEOUT --cow-size=small < $LOGFILE 2>/dev/null
    ##else
    ##    DISPLAY=:0 xmessage -timeout $MSG_TIMEOUT -file - < $LOGFILE 2>/dev/null
    ##fi
    # in case all GUI notification fails, use text
    #write `whoami` < $LOGFILE 2>/dev/null
    # The above may contain sensitive info, so say it briefly.
    echo "Possible autosync error. See $LOGFILE: `cat $LOGFILE`" | write `whoami` 2>/dev/null
    #pwd | write `whoami` 
else
    #svn commit -m "auto commit in Ubuntu `date`"
    #svn commit -m "auto commit `lsb_release -d` $HOSTTYPE by $USER "
    #$SCM commit -m "auto commit `lsb_release -d` $HOSTTYPE" # unable to get $USER in crontab
    #COMMITCMT=`hostname`
    #$SCM commit -m "auto commit $HOSTTYPE $COMMITCMT" # unable to get $USER in crontab
    #$SCM commit $SCMDIR -m "auto commit `whoami` `hostname` $HOSTTYPE" 1>/dev/null # unable to get $USER in crontab
    $SCM commit -a -m "auto commit `whoami`@`hostname` $HOSTTYPE" 1>/dev/null # unable to get $USER in crontab
    #$SCM commit -a -m "auto commit `whoami` `hostname` $HOSTTYPE" 1>/dev/null # unable to get $USER in crontab
    $SCM push $URL >|$LOGFILE 2>&1
    # whoami also works, but not quit useful
fi

# do one more update, in case some commands require update
#$SCM up $SCMDIR 2>|$LOGFILE 1>/dev/null
$SCM pull $URL >|$LOGFILE 2>&1

