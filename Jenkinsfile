pipeline {
  agent none
  stages {
    stage('git pull') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && sudo git pull'
      }
    }
    stage('go install') {
      steps {
        sh 'cd /home/goproject/src/go_mongo && sudo go install'
      }
    }
  }
}