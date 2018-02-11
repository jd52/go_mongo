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
    stage('go install') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && go install'
      }
    }
  }
}