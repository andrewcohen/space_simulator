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

var stats = new Stats();
stats.domElement.style.position = 'absolute';
stats.domElement.style.left = '0px';
stats.domElement.style.top = '0px';
document.body.appendChild( stats.domElement );

var simulation = { scale: 0.1 };
simulation.fullscreen = function() {
  document.documentElement.webkitRequestFullscreen();
};

var FOV = 75;
var ASPECT = window.innerWidth / window.innerHeight;
var NEAR = 0.1;
var FAR = 10000000;
var scene = new THREE.Scene();
simulation.camera = new THREE.PerspectiveCamera(FOV, ASPECT, NEAR, FAR);

var controls = new THREE.TrackballControls(simulation.camera);
controls.rotateSpeed = 1.0;
controls.zoomSpeed = 1.2;
controls.panSpeed = 0.8;
controls.noZoom = false;
controls.noPan = false;
controls.staticMoving = true;
controls.dynamicDampingFactor = 0.3;
controls.keys = [ 65, 83, 68 ];

var renderer = new THREE.WebGLRenderer();
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

var gui = new dat.GUI();
gui.add(simulation, 'fullscreen');
gui.add(simulation, 'scale', 1, 100);
var cameraFolder = gui.addFolder('Camera');
cameraFolder.add(simulation.camera.position, 'x', -10, 50000).listen();
cameraFolder.add(simulation.camera.position, 'y', -10, 50000).listen();
cameraFolder.add(simulation.camera.position, 'z', -10, 50000).listen();

var entities = [];

var geometry = new THREE.SphereGeometry(5, 32, 32);
var sphere = new THREE.Mesh(geometry, new THREE.MeshNormalMaterial());

simulation.camera.position.z = 300;

function render() {
  requestAnimationFrame(render);
  if (scene.children.length < entities.length) {
    console.log('adding missing 3js entities');
    var i, j = entities.length;

    var cumX = cumY = cumZ = 0;
    for (i = 0; i < j; i++) {
      var s = new THREE.Mesh(
        new THREE.SphereGeometry(entities[i].mass*simulation.scale, 32, 32),
        new THREE.MeshNormalMaterial);
      cumX += entities[i].position.x;
      cumY += entities[i].position.y;
      cumZ += entities[i].position.z;
      scene.add(s);
    }
    simulation.camera.position.set(cumX/j, cumY/j, cumZ/j);
  } else {
    var i, j = scene.children.length;
    for (i = 0; i < j; i++) {
      scene.children[i].position.x = entities[i].position.x;
      scene.children[i].position.y = entities[i].position.y;
      scene.children[i].position.z = entities[i].position.z;
    }
  }
  renderer.render(scene, simulation.camera);
  stats.update();
  controls.update();
}
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

window.onkeydown = function(e) {
  switch(e.which) {
    case 32: // SPACEBAR
      e.preventDefault();
      sendMessage([{CommandType: "add_planet"}]);
      break;
  }
};

window.addEventListener('resize', onWindowResize, false);
function onWindowResize() {
  simulation.camera.aspect = window.innerWidth/window.innerHeight;
  simulation.camera.updateProjectionMatrix();
  renderer.setSize(window.innerWidth, window.innerHeight);
}

var sendMessage = function(msg) {
  var json = JSON.stringify(msg);
  console.log("send msg:", json);
  conn.send(json);
}
