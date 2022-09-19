pipeline {
    agent any
    stages{
        stage('GitHub') {
            steps{
                git branch: 'main', url: 'https://github.com/'
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
                sh """
                #!/bin/bash
                cp TestSimulation.scala /tmp/gatling/user-files/simulations/computerdatabase
                /tmp/gatling/bin/gatling.sh -s computerdatabase.TestSimulation --run-mode local
                """
            }
        }

}
}