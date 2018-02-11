pipeline {
  agent {
    node {
      label 'master'
    }
    
  }
  stages {
    stage('Init') {
      steps {
        tool 'go'
      }
    }
    stage('git pull') {
      steps {
        dir(path: '/home/goproject/src/go_mongo') {
          sh 'git pull'
        }
        
      }
    }
    stage('go install') {
      steps {
        dir(path: '/home/goproject/src/go_mongo') {
          sh 'go install'
        }
        
      }
    }
  }
  environment {
    GOPATH = '/home/goproject'
    GOBIN = '/home/goproject/bin'
  }
}