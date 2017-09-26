pipeline {
  agent any 
  
  stages {
    
    stage('Say something') {
      steps {
        echo 'Saying hello'
      }
    }
    
    /* 
    stage ('Verify Tools'){
      steps {
        parallel (
          docker: { sh "docker -v" }
        )
      }
    }  */   
    
    stage ('Checkout Code') {
      steps {
        checkout scm
      }
    }

    
    stage('Build image') {
      agent { dockerfile { dir 'Example1' } }  
      steps {
        echo 'Built something'
      }
    }

    /* 
    stage('Test image') {
        app.inside {
            sh 'echo "Tests passed"'
        }
    }
    */
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
