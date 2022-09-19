pipeline {
    agent any
    stages{
        stage('GitHub') {
            steps{
                git branch: 'main', url: 'https://github.com/jocielalves/MicroserviceGo-Calc'
            }
        }
        stage('Start Calc') {
            steps{
                sh """
                #!/bin/bash
                cd app
                go run calc.go &
                """
            }
        }
        stage('Gatling Test'){
            steps{
                sh '''
                #!/bin/bash
                chmod +x config.sh
                ./config.sh
                '''
            }
        }

}
}