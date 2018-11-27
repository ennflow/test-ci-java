pipeline {
  agent none
     stages {
        stage('Build') { 
           agent {
               docker {
                    image 'harbor.enncloud.cn/enncloud/ceres-build:jdk8' 
                      }
                  }
            steps {
                    sh 'cd /var/jenkins_home/workspace/cicdtest/OpenUrl/src/pyrmont/'
                    sh 'pwd'
                    sh 'ls'
                    sh 'javac /var/jenkins_home/workspace/cicdtest/OpenUrl/src/pyrmont/*.java'
                    sh 'cd /var/jenkins_home/workspace/cicdtest/OpenUrl/'
                    sh 'echo ====================/openurl======================='
                    sh 'ls /var/jenkins_home/workspace/cicdtest/OpenUrl/src/pyrmont/'
                    sh 'echo ====================/openurl======================='
                    sh 'jar cvfm  /var/jenkins_home/workspace/cicdtest/OpenUrl/pyrmont.jar /var/jenkins_home/workspace/cicdtest/OpenUrl/mymanifest -C /var/jenkins_home/workspace/cicdtest/OpenUrl/src/ .'
                    sh 'ls /var/jenkins_home/workspace/cicdtest/OpenUrl/'
                  }
             }
       stage('Deploy') {
          agent {
             dockerfile {
                   filename 'Dockerfile'
                   label 'image-builder'
                   additionalBuildArgs  '--build-arg version=1.0.2'
               }
           }
       }
     } 
}
