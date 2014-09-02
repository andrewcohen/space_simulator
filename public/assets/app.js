var conn = null
connect();

var msgEl = document.getElementById('count');

function connect() {
  conn = new WebSocket('ws://localhost:3000/ws');
  conn.messageCount = 0;

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
    entities = parsed.entities || [];
    updateGUI(entities);

    conn.messageCount++;
    if (conn.messageCount > 250) {
      //console.log(parsed);
      conn.messageCount = 0;

      var i = entities.length;
      console.log(i + " entities");
      //while (i--) {
        //console.log(i + 1 + "| mass: "+ entities[i].mass+ ", x: "+ entities[i].position.x+ ", y: "+ entities[i].position.y);
      //}
      //console.log("\n \n \n");
    }
  }
}


var canvas = document.getElementById("canvas");
var gui = document.getElementById("gui");
var ctx = canvas.getContext('2d');
var WIDTH = canvas.width = document.body.clientWidth, HEIGHT = canvas.height = document.body.clientHeight;
var entities = [];
render();

var colors = ["#AAFF00", "#FFAA00", "#FF00AA", "#AA00FF", "#00AAFF"];

var STATIC_ENTITY = 0;
var DYNAMIC_ENTITY = 1;
var SCALE = 1;

function clamp(val, min, max) {
  if (val < min) return min;
  else if (val > max) return max;
  else return val;
}

function biggestEntity(entities) {
  var i = entities.length;
  var biggest = entities[0];
  while (i--) {
    if (biggest.mass < entities[i].mass) biggest = entities[i];
  }
  return biggest;
}

function render(timestamp) {
  ctx.setTransform(1,0,0,1,0,0);
  if (!streak) {
    ctx.fillStyle = "black";
    ctx.fillRect(0, 0, WIDTH, HEIGHT);
  }

  if (entities && entities.length > 0) {
    var biggest = entities[0];//biggestEntity(entities)
    var camX = clamp(-biggest.position.x + WIDTH/2, 0 - WIDTH, 10000 - WIDTH);
    var camY = clamp(-biggest.position.y + HEIGHT/2, 0 - HEIGHT, 10000 - HEIGHT);
    ctx.translate(camX, camY);

    var i = entities.length;
    while (i--) {
      var ent = entities[i];
      var pos = ent.position;
      ctx.fillStyle = colors[i % colors.length];
      ctx.beginPath();
      ctx.arc(pos.x, pos.y, clamp(ent.mass * SCALE / 200 , 1, 200), 0, Math.PI * 2, true);
      ctx.fill();
    }
  }
  requestAnimationFrame(render);
}

window.onkeydown = function(e) {
  switch(e.which) {
    case 32: // SPACEBAR
      e.preventDefault();
      sendMessage([{CommandType: "add_planet"}]);
      break;

    case 65: // A
      sendMessage([{commandType: "direct", kind: "move", direction: -1}]);
      break;
    case 68: // D
      sendMessage([{commandType: "direct", kind: "move", direction: 1}]);
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

var updateGUI = function(entities) {
  gui.innerText = "# Planets: " + entities.length;
};

var streak = false;
var toggleStreaks = document.getElementById("toggle-streaks").onclick = function(e) {
  e.preventDefault();
  streak = !streak;
}
