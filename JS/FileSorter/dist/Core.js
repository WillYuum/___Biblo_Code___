"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
var fs = __importStar(require("fs-extra"));
var path = __importStar(require("path"));
var IMAGE_DIR = "C:/Users/user/Desktop/result/images/";
var VIDEO_DIR = "C:/Users/user/Desktop/result/videos/";
var IMAGE_FORMATS = [".jpg", ".png"];
var VIDEO_FORMATS = [".mp4"];
var dirToSort = "C:/Users/user/Desktop/Takeout";
function getAllChildFiles() {
    var children = [dirToSort];
    for (;;) {
        if (children.every(function (file) { return fs.statSync(file).isFile(); }))
            break;
        children.forEach(function (child, index) {
            if (fs.statSync(child).isDirectory()) {
                delete children[index];
                var child_children = fs.readdirSync(child).map(function (v) { return child + "/" + v; });
                children.push.apply(children, child_children);
            }
        });
    }
    return children;
}
function HandleSortFiles() {
    var files = getAllChildFiles();
    var amountFilesMoved = 0;
    updateAmountOfFilesMove(5, amountFilesMoved);
    files.forEach(function (filePath) {
        var extName = path.extname(filePath);
        extName = extName.toLowerCase();
        if (IMAGE_FORMATS.includes(extName)) {
            //handle adding file to images folder
            copyFileToFolder(filePath, IMAGE_DIR);
        }
        else if (VIDEO_FORMATS.includes(extName)) {
            //handle adding file to videos folder
            copyFileToFolder(filePath, VIDEO_DIR);
        }
        amountFilesMoved += 1;
    });
    console.log("Moved " + amountFilesMoved + " file(s)");
    console.log(files.length);
}
function copyFileToFolder(fileToCopy, destination) {
    var fileName = path.basename(fileToCopy);
    var newDestination = destination + fileName;
    fs.copyFileSync(fileToCopy, newDestination);
}
function updateAmountOfFilesMove(updateInterval, amountFilesMoved) {
    if (updateInterval === void 0) { updateInterval = 5; }
    setTimeout(function () {
        console.log("Moved $" + amountFilesMoved + " until now...");
    }, updateInterval * 1000);
}
exports.Controllers = {
    HandleSortFiles: HandleSortFiles
};
//# sourceMappingURL=Core.js.map