pipeline {
  agent {
    dockerfile {
      filename 'Example2/Dockerfile'
    }
    
  }
  stages {
    stage('Say something') {
      steps {
        echo 'Saying hello'
      }
    }
  }
}
