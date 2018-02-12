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
        sh 'cd /home/goproject/src/go_mongo && sudo pkill -f go_mongo && echo success  '
      }
    }
    stage('start program') {
      steps {
        sh '''sudo
nohup ./go_mongo &'''
      }
    }
  }
  environment {
    GOPATH = '/home/goproject'
    GOBIN = '/home/goproject/bin'
  }
}