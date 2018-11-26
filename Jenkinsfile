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
                sh 'cp -r /builds/weihongwei/cicdproject/* /go/src/'
                sh 'cd /go/src/beeblog/'
                sh 'echo "==========" > test.txt'
                sh 'go build -o mytest'
                sh 'cp mytest /builds/weihongwei/cicdproject/beeblog/' 
                sh 'ls /builds/weihongwei/cicdproject/beeblog/'
            }
        }
    }
}
