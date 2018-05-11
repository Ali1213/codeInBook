const fs = require('fs');
const http = require('http');
const path = require('path');

http.createServer((req,res)=>{
    let a = fs.createReadStream(path.join(__dirname,'./index.html'));
    a.on('error',(err)=>res.end('Not Found'))
    a.pipe(res);
}).listen(3001)