import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';


platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(err => console.error(err));


type myFile = string;
type myDirectory = {
  [name:string] : myFile | myDirectory
}

export function pathArrayToDirStructure(files:string[]) : myDirectory {
  let root: myDirectory = {}

  files.forEach(
    file => {
      let path = file.split("/");
      let fNamearr = path.splice(path.length - 1);
      let fName = fNamearr[0];

      let dir = root;
      path.forEach( pEl => {
        let newDir = root[pEl];
        if (typeof newDir == "object") {
          dir = newDir
        }
        else if (!newDir) {
          newDir = {};
          dir[pEl] = newDir;
          dir = newDir;
        }
      })
      dir[fName] = "file";
      root[file] = "file"
    }
  )
  return {}
}