pipeline {
    agent {
        docker {
            image 'harbor.enncloud.cn/enncloud/ceres-build:jdk8' 
        }
        dockerfile {
           filename 'Dockerfile'
           label 'cicdtest-label'
           registryUrl 'https://harbor.enncloud.cn'
       }
    }
    stages {
        stage('Build') { 
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
       age('deploy') { 
            steps {
                   sh 'ls'
                   sh 'pwd'
                   sh 'docker build -t harbor.enncloud.cn/create-cicd-hub/cicd-java:v1126 .'
                   sh 'docker push harbor.enncloud.cn/create-cicd-hub/cicd-java:v1126'

            }
        }
    }
}
