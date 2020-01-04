function setup() {
    createCanvas(window.innerWidth, window.innerHeight, WEBGL);
}
let angle = 0
function draw() {
    background(175);
    rotateX(angle);
    rotateY(angle);
    rotateZ(angle * 0.75);
    // box(w, h, d) 
    // if only one arg, then use the value for all three dimensions
    box(175);
    angle += 0.025;
}