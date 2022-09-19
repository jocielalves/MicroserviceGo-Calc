
version=2.3.1

cd ../../..
wget https://repo1.maven.org/maven2/io/gatling/highcharts/gatling-charts-highcharts-bundle/$version/gatling-charts-highcharts-bundle-$version-bundle.zip
unzip gatling-charts-highcharts-bundle-$version-bundle.zip
cp TestSimulation.scala gatling-charts-highcharts-bundle-$version/user-files/simulations/computerdatabase/
cd gatling-charts-highcharts-bundle-$version/bin/
./gatling.sh -s computerdatabase.TestSimulation