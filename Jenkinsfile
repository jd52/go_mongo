pipeline {
  agent any
  stages {
    stage('Init') {
      steps {
        echo 'Test message'
      }
    }
    stage('run local script') {
      steps {
        sh 'cd /home/goproject && sudo ./jenkins_script.sh'
      }
    }
  }
}