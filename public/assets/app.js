var conn = null
connect();

var msgEl = document.getElementById('count');

function connect() {
  conn = new WebSocket('ws://localhost:3000/ws');

  conn.onconnect = function(e) {
    console.log('connected');
  };

  conn.onclose = function(e) {
    console.log('connection closed');
    connect()
  }

  conn.onmessage = function(e) {
    var parsed = JSON.parse(e.data);
    msgEl.innerText = "uptime: " + parsed.uptime + "s";
    if (parsed.uptime % 20 == 0) {
      console.log(parsed);
    }
    entities = parsed.entities;
  }
}

var canvas = document.getElementById("canvas");
var ctx = canvas.getContext('2d');
var WIDTH = canvas.width, HEIGHT = canvas.height;
var entities = [];
render();

var colors = ["blue", "brown", "green", "yellow", "pink", "orange", "purple", "red"];

function render(timestamp) {
  ctx.clearRect(0, 0, WIDTH, HEIGHT);

  if (entities.length > 0) {

    var i = entities.length;
    while (i--) {
      var ent = entities[i];
      var pos = ent.position;
      ctx.fillStyle = colors[ent.team_id];
      ctx.fillRect(pos.x, pos.y, 20, 20);
    }
  }
  requestAnimationFrame(render);
}
