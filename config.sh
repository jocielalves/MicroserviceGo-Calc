cd ../../..
wget https://repo1.maven.org/maven2/io/gatling/highcharts/gatling-charts-highcharts-bundle/3.8.4/gatling-charts-highcharts-bundle-3.8.4-bundle.zip
unzip gatling-charts-highcharts-bundle-3.8.4-bundle.zip
cp -a TestSimulation.scala gatling-charts-highcharts-bundle-3.8.4/user-files/simulations/computerdatabase/
cd gatling-charts-highcharts-bundle-3.8.4/bin/
./gatling.sh -s computerdatabase.TestSimulation