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
       stage('deploy') {
         withDockerRegistry(url:'https://harbor.enncloud.cn'){
         def dockerfile = 'Dockerfile'
         def customImage = docker.build("harbor.enncloud.cn/create-cicd-hub/cicdtest-java:v1126","-f $(dockerfile) .")
         customImage.push()
         }
       }
}
