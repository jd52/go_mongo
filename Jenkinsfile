pipeline {
  agent {
    node {
      label 'master'
    }
    
  }
  stages {
    stage('git pull') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && git pull'
      }
    }
    stage('go build') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && go build'
      }
    }
    stage('stop program') {
      steps {
        sh '''cd /home/goproject/src/go_mongo
go_mongo_pid = ps -aef | grep \'go_mongo\' | grep -v grep | awk \'{print $2}\'
echo $go_mongo_pid
if ![ -z "$go_mongo_pid" ]; then
	kill -9 $go_mongo_pid
else
	echo not running
fi'''
      }
    }
    stage('start program') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && ./go_mongo disown &'
      }
    }
  }
  environment {
    GOPATH = '/home/goproject'
    GOBIN = '/home/goproject/bin'
  }
}