var sparkles = {};

var HttpClient = function() {
    this.get = function(aUrl, aCallback) {
        var anHttpRequest = new XMLHttpRequest();
        anHttpRequest.onreadystatechange = function() {
            if (anHttpRequest.readyState == 4 && anHttpRequest.status == 200)
                aCallback(anHttpRequest.responseText);
        }
        anHttpRequest.open( "GET", aUrl, true );
        anHttpRequest.send( null );
    }
}

// Get sparkles from the server
function getSparkles() {
  sClient = new HttpClient();
  sClient.get('/sparkles', function(response) {
    sparkles = JSON.parse(response);
    buildSparkleGraphEdges();
  });
}

// Given sparkles "s", build a graph. This is called from getSparkles.
var edges = [];

// Build a directed graph of who sparkled whom. Only edges are required.
function buildSparkleGraphEdges() {
  _.forEach(sparkles, function(sparkle) {
    var idx = _.findIndex(edges, function(e) { return e.a == sparkle.sparklee && e.b == sparkle.sparkler });
    if (idx > 0) {
      edges[idx].score++;
    } else {
      edges = edges.concat({"a": sparkle.sparklee, "b": sparkle.sparkler, "score": 1});
    }
  })
}

function graphSparklesOverTime(s) {
  // Graph sparkles <s> over time
}

function getStatsForUser(user) {
  // Get stats for a user 
}

// Display for everyone:
// Sparkles over time (line graph)
// Groups (cluster graph)
// Interesting stats?
// * Likelihood someone sparkles another user after being sparkled
// * Sparkle word cloud (minus sparkle party)
// * Party prevalence

// Display for a user:
// * Sparkles over time (line graph)
// * First sparkle received
// * First sparkle given
// * Top sparkled (given and received)
// * Trait: more sparkles given or received, partier, popular words?
// * Similar users



function init() {
  getSparkles();
}

init();
