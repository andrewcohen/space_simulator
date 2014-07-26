var conn = null
connect();

var msgEl = document.getElementById('count');

function connect() {
  conn = new WebSocket('ws://localhost:3000/ws');

  conn.onopen = function(e) {
    console.log('connected');
    sendMessage([{CommandType: "join", kind: "join"}]);
  };

  conn.onclose = function(e) {
    console.log('connection closed');
    connect()
  }

  conn.onmessage = function(e) {
    var parsed = JSON.parse(e.data);
    lastMsg = parsed;
    entities = parsed.entities;
  }
}

var canvas = document.getElementById("canvas");
var ctx = canvas.getContext('2d');
var WIDTH = canvas.width = document.body.clientWidth, HEIGHT = canvas.height = document.body.clientHeight;
var entities = [];
render();

var colors = ["blue", "brown", "green", "yellow", "pink", "orange", "purple", "red"];

var STATIC_ENTITY = 0;
var DYNAMIC_ENTITY = 1;

function render(timestamp) {
  ctx.clearRect(0, 0, WIDTH, HEIGHT);

  if (entities && entities.length > 0) {

    var i = entities.length;
    while (i--) {
      var ent = entities[i];
      var pos = ent.position;
      ctx.fillStyle = colors[ent.team_id];
      ctx.fillRect(pos.x, pos.y, ent.size.x, ent.size.y);
    }
  }
  requestAnimationFrame(render);
}

window.onkeydown = function(e) {
  switch(e.which) {
    case 32:
      e.preventDefault();
      sendMessage([{CommandType: "direct", kind: "jump"}]);
      break;

    case 65: // A
      sendMessage([{CommandType: "direct", kind: "move", direction: -1}]);
      break;
    case 68: // D
      sendMessage([{CommandType: "direct", kind: "move", direction: 1}]);
      break;
  }
};

canvas.onmouseup = function(e) {
  var data = {kind: "move", x: e.offsetX, y: e.offsetY};
  sendMessage([data]);
};

var sendMessage = function(msg) {
  var json = JSON.stringify(msg);
  console.log("send msg:", json);
  conn.send(json);
}
