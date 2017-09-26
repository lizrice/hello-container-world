pipeline {
  agent none
  
  stages {
    
    def app 
    
    stage('Say something') {
      steps {
        echo 'Saying hello'
      }
    }
    
    stage('Clone repository') {
        checkout scm
    }

    stage('Build image') {
        app = docker.build("lizrice/hello-container-world/Example1")
    }

    stage('Test image') {
        app.inside {
            sh 'echo "Tests passed"'
        }
    }

    /* stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
       /*  docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("${env.BUILD_NUMBER}")
            app.push("jenkins-test")
        }
    } */
  }
}
