import * as fs from "fs-extra";
import * as path from "path";

const IMAGE_DIR = "C:/Users/user/Desktop/result/images/";
const VIDEO_DIR = "C:/Users/user/Desktop/result/videos/";


const IMAGE_FORMATS = [".jpg", ".png"];
const VIDEO_FORMATS = [".mp4"];

let dirToSort = "C:/Users/user/Desktop/Takeout";






function getAllChildFiles(): string[] {
  let children: string[] = [dirToSort];

  for (; ;) {
    if (children.every(file => fs.statSync(file).isFile())) break;

    children.forEach((child: string, index) => {
      if (fs.statSync(child).isDirectory()) {
        delete children[index];

        let child_children = fs.readdirSync(child).map(v => `${child}/${v}`);
        children.push(...child_children);
      }
    })
  }

  return children;
}


function HandleSortFiles(){
  let files = getAllChildFiles();

  let amountFilesMoved = 0;

  updateAmountOfFilesMove(5, amountFilesMoved);
  
  files.forEach((filePath:string) =>{
    let extName = path.extname(filePath);
    extName = extName.toLowerCase();

    if(IMAGE_FORMATS.includes(extName)){
      //handle adding file to images folder
      copyFileToFolder(filePath, IMAGE_DIR);
    }else if(VIDEO_FORMATS.includes(extName)){
      //handle adding file to videos folder
      copyFileToFolder(filePath, VIDEO_DIR);
    }
    amountFilesMoved += 1;
  });

  console.log(`Moved ${amountFilesMoved} file(s)`);
}


function copyFileToFolder(fileToCopy:string, destination:string){
  let fileName = path.basename(fileToCopy);
  let newDestination = destination + fileName;

  fs.copyFileSync(fileToCopy, newDestination);
}


function updateAmountOfFilesMove(updateInterval: number = 5, amountFilesMoved:number){
  setTimeout(()=>{
    console.log(`Moved $${amountFilesMoved} until now...`);
  }, updateInterval * 1000);
}


export const Controllers = {
  HandleSortFiles
}