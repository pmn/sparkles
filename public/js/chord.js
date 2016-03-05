var graph = {};

var all_nodes = function(edges) {
    var nodes = {}
    $(edges).map(function(i,e) {
        nodes[e.sparklee] = i;
        nodes[e.sparkler] = i;
    });
    return Object.keys(nodes).sort();
}

var nodes = [
    "azizshamim",
    "johnagan",
    "gpadak",
    "jessla",
    "pmn",
    "dsorkin",
    "sachinr",
    "davideg",
    "joewadcan",
];

function mArray(size) {
  column = new Array();
  for (var i=0; i<size; i++) {
      column.push(new Array(size).fill(0));
  }
  return column;
}

var buildChord = function (data) {
  var chord = d3.layout.chord()
    .padding(.01)
    .sortSubgroups(d3.descending)
    .matrix(data.matrix);

  // set some magic paramters to size the chord.
  var width = 960,
    height = 500,
    innerRadius = Math.min(width, height) * .381,
    outerRadius = innerRadius * 1.025;

  // create the fill ordering
  var fill = d3.scale.ordinal()
    .domain(d3.range(5))
    .range([].concat(
        colorbrewer.Set1[5]
    ));

  // create the graph and position it
  var svg = d3.select("body").append("svg")
    .attr("width", width)
    .attr("height", height)
    .append("g")
    .attr("transform", "translate(" + width / 2 + "," + height / 2 + ")");

  svg.append("g").selectAll("path")
    .data(chord.groups)
  .enter().append("path")
    .style("fill", function(d) { return fill(d.index); })
    .style("stroke", function(d) { return fill(d.index); })
    .attr("d", d3.svg.arc().innerRadius(innerRadius).outerRadius(outerRadius))
    .on("mouseover", fade(.1))
    .on("mouseout", fade(1));

  var ticks = svg.append("g").selectAll("g")
    .data(chord.groups)
    .enter().append("g").selectAll("g")
      .data(groupTicks)
    .enter().append("g")
      .attr("transform", function(d) {
         return "rotate(" + (d.angle * 180 / Math.PI - 90) + ")"
            + "translate(" + outerRadius + ",0)";
    });

  ticks.append("line")
    .attr("x1", 1)
    .attr("y1", 0)
    .attr("x2", 5)
    .attr("y2", 0)
    .style("stroke", "#000");

  ticks.append("text")
    .attr("x", 8)
    .attr("dy", ".35em")
    .attr("transform", function(d) { return d.angle > Math.PI ? "rotate(180)translate(-16)" : null; })
    .style("text-anchor", function(d) { return d.angle > Math.PI ? "end" : null; })
    .text(function(d) { return d.label; });

  svg.append("g")
    .attr("class", "chord")
    .selectAll("path")
    .data(chord.chords)
    .enter().append("path")
    .attr("d", d3.svg.chord().radius(innerRadius))
    .style("fill", function(d) { return fill(d.target.index); })
    .style("opacity", 1);

  // Returns an array of tick angles and labels, given a group.
  function groupTicks(d) {
    var k = (d.endAngle - d.startAngle) / 2
    return [{
      angle: k + d.startAngle,
      label: data.nodes[d.index]
    }];
  }

  // Returns an event handler for fading a given chord group.
  function fade(opacity) {
    return function(g, i) {
      svg.selectAll(".chord path")
        .filter(function(d) { return d.source.index != i && d.target.index != i; })
        .transition()
        .style("opacity", opacity);
    };
  }
}

var grouped = function(edges) {
    return values = edges.reduce(function(p,c,ci,a) {
        //c.sparkler c.sparklee c.weight
        if (!p[c.sparkler]) { p[c.sparkler] = {} };
        p[c.sparkler][c.sparklee] = c.weight;
        return p;
    },{});
};

var matrix = function(sparkles) {
    var m2 = mArray(sparkles.nodes.length);
    _.each(sparkles.nodes, function(sparkler,i) {
        _.each(sparkles.nodes, function(sparklee,j) {
            if (!sparkles.grouped[sparkler]) {
            } else {
                if (!sparkles.grouped[sparkler][sparklee]) {
                } else {
                    m2[i][j] = sparkles.grouped[sparkler][sparklee];
                }
            }
        });
    });
    return m2;
};

$.get('/graph.json','',function(data) {
    data["nodes"] = nodes;
    //console.log("Total nodes:" + data.nodes.length);
    //console.log("Matrix size:" + Math.pow(data.nodes.length,2));
    data["grouped"] = grouped(data.edges);
    data["matrix"] = matrix(data);
    //console.log("Total nodes: " + data.matrix.length);
    //console.log("Total matrix: "+ _.flatten(data.matrix).length);
    //console.log(data);
    buildChord(data);
});
