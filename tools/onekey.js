const fs = require("fs");
const path = require("path");
const { exec } = require('child_process');
const helputil = require("./helputil");

// 定义要搜索的目录
const directory = '../internal/protocol/gofile';
// 定义要替换的文本
const searchFor = 'proto "github.com/golang/protobuf/proto"';
const searchFor1 = 'const _ = proto.ProtoPackageIsVersion4';
const replaceWith = '';


//--------------------------------------------------------------
var protoDir = "../internal/protocol"
//server begin
var outServerMsg =  "../nodes/game/msg/msg.go";
var outRouter = "../nodes/leaf/jettengame/gate/router.go";

var all_protos = {};
var all_message = {};
var g_CMDID = -1;
//要生成的pb文件
var pbfiles = []
//--------------------------------------------------------------

function getCmdId() {
	g_CMDID++;
	return g_CMDID;
}

function parseProtoFile(protoPath, fileName) {
    if (!protoPath.endsWith(".proto")) {
        return;
    }

    let ptoName = "proto_" + fileName;
    all_protos[fileName] = [];
    console.log("----------------", protoPath, fileName);

    if(fileName.indexOf("chat") != -1){
        pbfiles.push({
            name:fileName,
            router:"chat",
            channel:"game"
        })
    }else if(fileName.indexOf("login") != -1){
        pbfiles.push({
            name:fileName,
            router:"login",
            channel:"game"
        })
    }else{
        pbfiles.push({
            name:fileName,
            router:"game",
            channel:"game"
        })
    }

  

    let protoFileText = fs.readFileSync(protoPath, { encoding: "utf8" });
    protoFileText = protoFileText.trim();
    let lines = protoFileText.split("\r\n");
    
    let isNote = false
    lines.forEach((line) => {
        if(line.startsWith("/*")) {
            isNote= true
        }
        if(line.endsWith("*/")) {
            isNote= false
        }
        if(isNote) {
            return
         }
        if (line.startsWith("message")) {
            let msgLines = line.split(" ");
            let msgName = msgLines[1];
            
            
            if(msgName.endsWith("{")) {
                msgName = msgName.substring(0, msgName.length-1);
            }

            console.log(fileName, msgName);

            all_message[msgName] = fileName;
            all_protos[fileName].push(msgName);
        }
    });
}

function genCode() {
    //创建导出目录
	// helputil.createDirectory(cfgData.clientOutDir);
    //遍历协议目录
    helputil.travelDirectory(protoDir, (item_path, item) => {
        parseProtoFile(item_path, helputil.getFileName(item))
    });

    //生成代码
    genServer();
}

function genServer() {
	var msgStr = "";
	var routerStr = "";

    msgStr = "//---------------------------------\n";
	msgStr += "//该文件自动生成，请勿手动更改\n";
	msgStr += "//---------------------------------\n";
    msgStr += "package msg\n\n"
    msgStr += "import (\n"
    msgStr += '    "google.golang.org/protobuf/proto"\n'
    msgStr += '    "superman/internal/utils"\n'
    msgStr += '    "superman/nodes/game/msg/process"\n'
    msgStr += '    protoMsg "superman/internal/protocol/gofile"\n'
    msgStr += ')\n\n'
	msgStr += "// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）\n"
	msgStr += "// var ProcessorJson = json.NewProcessor()\n"
	msgStr += "//var ProcessorProto = protobuf.NewProcessor()\n\n"
    msgStr += " var ProcessorProto = process.NewProcessor(false)\n\n"
    msgStr += " var ServerChanRPC = process.NewServer(10000)\n\n"

	msgStr += "//对外接口 【这里的注册函数并非线程安全】\n"
    msgStr += "var msgMap = make(map[uint16]string, 0)\n"
	msgStr += "func RegisterMessage(message proto.Message) {\n"
	msgStr += "    id, name := ProcessorProto.Register(message)\n"
    msgStr += "    msgMap[id] = name\n"
    msgStr += "    ProcessorProto.SetRouter(message, ServerChanRPC)\n"
	msgStr += "}\n\n"
	msgStr += "func init() {"


    routerStr = "//---------------------------------\n";
	routerStr += "//该文件自动生成，请勿手动更改\n";
	routerStr += "//---------------------------------\n";
	routerStr += "package gate\n\n";
	routerStr += 'import (\n';
	routerStr += '    "superman/nodes/leaf/jettengame/game"\n';
	// routerStr += '    "superman/nodes/leaf/jettengame/login"\n';
	routerStr += '    "superman/nodes/leaf/jettengame/msg"\n';
	// routerStr += '    "server/robot"\n';
	routerStr += '    protoMsg "superman/internal/protocol/gofile"\n';
	routerStr += ')\n\n';
	routerStr += '//路由模块分发消息【模块间使用 ChanRPC 通讯，消息路由也不例外】\n';
	routerStr += '//注:需要解析的结构体才进行路由分派，即用客户端主动发起的)\n';
	routerStr += "func init() {";


    for(let i=0; i<pbfiles.length; i++) {
        let info = pbfiles[i];
        let fileName = info.name;
        console.log("生成代码----- ", info, " filename: "+info.name, " router: "+info.router, " channel: "+info.channel);

        msgStr += "\n    //" + fileName + "文件生成的代码\n";
		routerStr += "\n    //" + fileName + "文件生成的代码\n";

        let msgList = all_protos[fileName];
        for(let msgName of msgList) {
            let msgId = getCmdId();
            console.log(msgId, msgName);
            msgStr += "    RegisterMessage(&protoMsg." + msgName + "{})\n";
            // if(msgName.substring(msgName.length-3)==="Req" && info.router !== "login"){
            //     routerStr += "    msg.ProcessorProto.SetRouter(&protoMsg."+ msgName +"{}, "  +"robot.ChanRPC)\n";
            // }else{
                routerStr += "    msg.ProcessorProto.SetRouter(&protoMsg."+ msgName +"{}, "+ info.router +".ChanRPC)\n";
            // }
        }
    }

    routerStr += "}\n\n"
	msgStr += '  \n\tutils.ToJsonFile("./../../config/leafconf/message_id.json", msgMap, "", "\\t") \n'
	msgStr += "\tmsgMap = nil\n"
	msgStr += "}\n"

    
	
	
	helputil.write2file(outServerMsg, msgStr);
	helputil.write2file(outRouter, routerStr);
}



// 遍历目录
fs.readdir(directory, (err, files) => {
    if (err) {
        console.error('Error reading directory:', err);
        return;
    }

    files.forEach(file => {
        // 检查文件扩展名是否为 .go
        if (path.extname(file) === '.go' && file.includes('pb.go') ) {
            // 构建完整的文件路径
            const filePath = path.join(directory, file);
            // 读取文件内容
            fs.readFile(filePath, 'utf8', (err, content) => {
                if (err) {
                    console.error('Error reading file:', err);
                    return;
                }

                // 替换文件内容
                let newContent = content.replace(new RegExp(searchFor, 'g'), replaceWith);
                newContent = newContent.replace(new RegExp(searchFor1, 'g'), replaceWith);
                // 写回文件
                fs.writeFile(filePath, newContent, 'utf8', err => {
                    if (err) {
                        console.error('Error writing file:', err);
                        return;
                    }
                    console.log(`File ${file} has been updated.`);
                });
            });
        }
    });
});

genCode();


