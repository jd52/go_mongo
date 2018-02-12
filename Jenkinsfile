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
    stage('stop start program') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && sudo pskill -f go_mongo && sudo nohup ./go_mongo &'
      }
    }
  }
  environment {
    GOPATH = '/home/goproject'
    GOBIN = '/home/goproject/bin'
  }
}