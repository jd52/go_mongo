pipeline {
  agent {
    node {
      label 'master'
    }
    
  }
  stages {
    stage('run local jenkins script') {
      steps {
        sh 'sudo /home/goproject/jenkins_script.sh'
      }
    }
  }
  environment {
    GOPATH = '/home/goproject'
    GOBIN = '/home/goproject/bin'
  }
}