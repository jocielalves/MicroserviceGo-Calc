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
		            cd app
                    cp TestSimulation.scala /home/ilegra/gatling-charts-highcharts-bundle-3.8.4/user-files/simulations/computerdatabase/
		            cd  /home/ilegra/gatling-charts-highcharts-bundle-3.8.4/user-files/simulations/computerdatabase/
		            ./gatling.sh -s TestSimulation --run-mode local
                '''
            }
        }

}
}