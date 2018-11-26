pipeline {
    agent {
        docker {
            image 'harbor.enncloud.cn/enncloud/ceres-build:jdk8' 
        }
    }
    stages {
        stage('Build') { 
            steps {
                    sh 'cd OpenUrl/src/pyrmont/'
                    sh 'javac *.java'
                    sh 'pwd'
                    sh 'ls'
                    sh 'cd /var/jenkins_home/workspace/cicdtest/OpenUrl/'
                    sh 'echo ====================/openurl======================='
                    sh 'ls'
                    sh 'echo ====================/openurl======================='
                    sh 'jar cvfm  pyrmont.jar mymanifest -C src/ .'
                    sh 'ls'
            }
        }
    }
}
