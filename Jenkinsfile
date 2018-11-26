pipeline {
    agent {
        docker {
            image 'harbor.enncloud.cn/enncloud/golang:1.8.3' 
        }
    }
    stages {
        stage('Build') { 
            steps {
                sh 'pwd'
                sh 'go env'
                sh 'cp -r /var/jenkins_home/workspace/cicdtest/* /go/src/'
                sh 'cd /go/src/beeblog/'
                sh 'echo "==========" > test.txt'
                sh 'go build -o mytest'
                sh 'cp mytest /var/jenkins_home/workspace/cicdtest/beeblog/' 
                sh 'ls /var/jenkins_home/workspace/cicdtest/beeblog/'
            }
        }
    }
}
