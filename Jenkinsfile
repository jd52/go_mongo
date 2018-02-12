pipeline {
  agent {
    node {
      label 'master'
    }
    
  }
  stages {
    stage('run jenkins script') {
      steps {
        sh 'sudo /home/goproject/jenkins_script.sh'
      }
    }
  }
}