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
        sh '/home/goproject/jenkins_script.sh'
      }
    }
  }
  environment {
    GOPATH = '/home/goproject'
    GOBIN = '/home/goproject/bin'
  }
}