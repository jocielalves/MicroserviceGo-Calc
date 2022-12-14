
import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class TestSimulation extends Simulation {
    
     val httpConf = http
    .baseUrl("http://localhost:8082") 
    .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    .doNotTrackHeader("1")
    .acceptLanguageHeader("en-US,en;q=0.5")
    .acceptEncodingHeader("gzip, deflate")
    .disableFollowRedirect
    .userAgentHeader("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0")

  val scn = scenario("Test History Calculator")
          .exec(http("History") 
            .get("/calc/history")
            .check(status.is(200))
          )
    
  setUp(
      scn.inject(constantUsersPerSec(100).during(60 seconds))
  ).protocols(httpConf)

}