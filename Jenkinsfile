pipeline {
  agent any
  stages {
    stage('Init') {
      steps {
        echo 'Test message'
      }
    }
    stage('git pull') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && git pull'
      }
    }
  }
}