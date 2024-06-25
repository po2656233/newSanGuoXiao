const fs = require("fs");
const path = require("path");
const { exec } = require('child_process');


//接口：遍历目录
function readDirectory (dir, obj, rootSrc) {
	var stat = fs.statSync(dir);
	if (!stat.isDirectory()) {
		return;
	}
	var subpaths = fs.readdirSync(dir), subpath, size, md5, compressed, relative;
	for (var i = 0; i < subpaths.length; ++i) {
		if (subpaths[i][0] === '.') {
			continue;
		}
		subpath = path.join(dir, subpaths[i]);
		stat = fs.statSync(subpath);
		if (stat.isDirectory()) {
			readDirectory(subpath, obj, rootSrc);
		}
		else if (stat.isFile()) {
			// Size in Bytes
			size = stat['size'];
			md5 = crypto.createHash('md5').update(fs.readFileSync(subpath, 'binary')).digest('hex');
			compressed = path.extname(subpath).toLowerCase() === '.zip';

			relative = path.relative(rootSrc, subpath);
			relative = relative.replace(/\\/g, '/');
			relative = encodeURI(relative);
			obj[relative] = {
				'size' : size,
				'md5' : md5
			};
			if (compressed) {
				obj[relative].compressed = true;
			}
		}
	}
}

//接口：创建文件夹
function createDirectory(path) {
	try {
		fs.mkdirSync(path);
	} catch(e) {
		if ( e.code != 'EEXIST' ) throw e;
	}
}

function createDir(dirName) {
	if(!fs.existsSync(dirName)) {
		fs.mkdirSync(dirName);
	}
}

function travelDirectory(dir, func) {
	if(!fs.existsSync(dir)) {
		console.error("dir not exist: ", dir);
		return;
	}
	var files = fs.readdirSync(dir);
	files.forEach(function(item){
		var item_path = path.join(dir, item);
		// console.log(item_path);
		if (fs.statSync(item_path).isDirectory()) {
			travelDirectory(item_path, func);
		} else {
			func(item_path, item);
		}
	});
}

//接口：删除文件夹
function deleteDirectory(dir) {
	if (fs.existsSync(dir) == true) {
		var files = fs.readdirSync(dir);
		files.forEach(function(item){
			var item_path = path.join(dir, item);
		   // console.log(item_path);
			if (fs.statSync(item_path).isDirectory()) {
				deleteDirectory(item_path);
			}
			else {
				fs.unlinkSync(item_path);
			}
		});
		fs.rmdirSync(dir);
	}
}

function deleteFile(fname) {
	if(!fname || fname == "") { return; }
	if(fs.existsSync(fname)) {
		fs.unlinkSync(fname);
		console.log("delete file: ", fname);
	} else {
		console.log("delete unexist file: ", fname);
	}
}

function delFilesByFilter(dstDir, filterFunc) {
	if (fs.existsSync(dstDir) == false) {
		return;
	}
	var dirs = fs.readdirSync(dstDir);
	dirs.forEach(function(item){
		var item_path = path.join(dstDir, item);
		var temp = fs.statSync(item_path);
		if (temp.isFile()) { // 是文件
			if(filterFunc) {
				if(filterFunc(item_path)) {
					fs.unlinkSync(item_path);
				}
			} else {
				fs.unlinkSync(item_path);
			}
		}
	});
}

//接口：拷贝文件夹
function copyDirectory(fromDir, toDir, filterFunc) {
	if (fs.existsSync(toDir) == false) {
		fs.mkdirSync(toDir);
	}
	if (fs.existsSync(fromDir) == false) {
		return false;
	}
	// console.log("fromDir:" + fromDir + ", toDir:" + toDir);
	// 拷贝新的内容进去
	var dirs = fs.readdirSync(fromDir);
	dirs.forEach(function(item){
		var item_path = path.join(fromDir, item);
		var temp = fs.statSync(item_path);
		if (temp.isFile()) { // 是文件
			// console.log("Item Is File:" + item);
			if(filterFunc) {
				if(filterFunc(item_path)) {
					fs.copyFileSync(item_path, path.join(toDir, item));
				}
			} else {
				fs.copyFileSync(item_path, path.join(toDir, item));
			}
		} else if (temp.isDirectory()){ // 是目录
			// console.log("Item Is Directory:" + item);
			copyDirectory(item_path, path.join(toDir, item));
		}
	});
}


function write2file(filepath, str) {
	try {
		fs.writeFileSync(filepath, str, 'utf8');
	} catch (err) {
		console.log("fail write file: ", filepath);
		console.log(err);
	}
}

//获取文件名字
//例如：传入"projects/client1/abcefg.txt"，将返回"abcefg"
function getFileName(pathstr="") {
	var xg = pathstr.lastIndexOf("/") + 1;
	var pt = pathstr.indexOf("\.");
	return pathstr.slice(xg, pt);
}

function getSufix(str) {
	var idx = str.lastIndexOf(".")
	if(idx === null || idx === undefined || idx <= 0) {
		return "";
	}
	return str.substring(idx+1);
}

//判断文件后缀是否为"proto"
function isProtoFile(pathstr) {
	if(pathstr===null || pathstr===undefined) {
		return false;
	}
	return getSufix(pathstr) == "proto";
}

function renameFile(filepath, newpath) {
	if(filepath == newpath) { return; }
	if(fs.existsSync(filepath)) {
		//console.log("renamefile: ", filepath, newpath);
		const data = fs.readFileSync(filepath, 'utf8');
		fs.writeFileSync(newpath, data, 'utf8');
		fs.unlinkSync(filepath);
	}
}

function isNil(v) {
	return v===null || v===undefined;
}

function firstToUpper(str){
    return str.toLowerCase().replace(/( |^)[a-z]/g,(L)=>L.toUpperCase());
}

//---------------------------------------------------------------------------------

//将proto文件的包名修改为wantName
function fixPackageName(filepath, wantName) {
	const data = fs.readFileSync(filepath, 'utf8').split('\n');

	var dependMuds = getImportMud(data);

	var pkgName = "";
	var needWrite = false;

	for(var i=0, len=data.length; i<len; i++) {
		var str = data[i];
		str = str.replace(';', ' ');  //去分号
		var arr = str.split(/\s+/); 	//去空白字符

		if(arr[0].match("//") || (arr[0]=="" && arr[1] && arr[1].match("//"))){
			//是一个注释
		} else {
			if( arr[0]=="package" && arr[1].indexOf("\.")<0 ) {
				pkgName = arr[1];
				if(pkgName !== wantName) {
					data[i] = "package "+wantName+";";
					needWrite = true;
				}
				break;
			}
			else if( arr[0]=="" && arr[1]=="package" && arr[2].indexOf("\.")<0 ){
				pkgName = arr[2];
				if(pkgName !== wantName) {
					data[i] = "package "+wantName+";";
					needWrite = true;
				}
				break;
			}
		}
	}

	if(needWrite) {
		//console.log("package name: ", pkgName, "  修改为：", wantName);
		fs.writeFileSync(filepath, data.join('\n'), 'utf8');
	} else {
		//console.log("package name: "+pkgName, "filename: "+wantName);
	}

	return dependMuds;
}

function fixClientOutput(filepath, dependMuds) {
	const data = fs.readFileSync(filepath, 'utf8').split('\n');

	for(var i=0, len=data.length; i<len; i++) {
		if(data[i].match("require") && data[i].match("protobuf")){
			data.splice(i, 1, 'var $protobuf = protobuf;');
			break;
		}
	}

	if(dependMuds && dependMuds != "") {
		var flag = false;
		for(var from=0, len=data.length; from<len; from++) {
			if(data[from].match("root."+dependMuds+" = ") && data[from].match("function")){
				//console.log("第"+from+"行发现累赘import: "+data[from]);
				data.splice(from);
				flag = true;
				break;
			}
		}
		if(flag) {
			data.push("module.exports = $root;");
		}
	}

	fs.writeFileSync(filepath, data.join('\n'), 'utf8');
}

//检查import关系
function getImportMud(data) {
	if(data===null || data===undefined) {
		return null;
	}
	var tmp = null;
	var len = data.length;
	if(len > 10) { len = 10 }
	for(var n=0; n<len; n++) {
		if(data[n].match("import") && !data[n].match("//")) {
			tmp = data[n].slice(0, data[n].indexOf('\.'));
			tmp = tmp.slice(tmp.indexOf("\"")+1, tmp.length);
			break;
		}
	}
	return tmp;
}

function getPackageName(data) {
	for(var i=0, len=data.length; i<len; i++) {
		var str = data[i];
		str = str.replace(';', ' ');  //去分号
		var arr = str.split(/\s+/); 	//去空白字符
		if(arr[0].match("//") || (arr[0]=="" && arr[1] && arr[1].match("//"))){
			//是一个注释
		}
		else {
			//解析 "message Request"
			if( arr[0]=="package" && arr[1].indexOf("\.")<0 ) {
				return arr[1];
			}
			else if( arr[0]=="" && arr[1]=="package" && arr[2].indexOf("\.")<0 ){
				return arr[2];
			}
		}
	}
	return null;
}




module.exports = {
	write2file: write2file,
	fixPackageName: fixPackageName,
	getFileName: getFileName,
	isProtoFile: isProtoFile,
	fixClientOutput: fixClientOutput,
	createDir: createDir,
	getImportMud: getImportMud,
	getPackageName: getPackageName,
	createDirectory: createDirectory,
	deleteDirectory: deleteDirectory,
	readDirectory: readDirectory,
	copyDirectory: copyDirectory,
	getSufix: getSufix,
	renameFile: renameFile,
	delFilesByFilter: delFilesByFilter,
	deleteFile: deleteFile,
	isNil: isNil,
	firstToUpper: firstToUpper,
	travelDirectory: travelDirectory,
}