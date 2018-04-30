const Readable = require('stream').Readable;
const express = require('express');
const app = express();
const util = require('util');



class StatStream extends Readable {
    constructor(limit){
        super()
        this.limit = limit
    }
    _read(size){
        if(this.limit === 0){
            this.push(null);
            return 
        }
        this.push(util.inspect(process.memoryUsage()));
        this.push('\n');
        this.limit--;
    }
}

app.get('/',(req,res)=>{
    const statStream = new StatStream(10);
    statStream.pipe(res)
}).listen(3003)