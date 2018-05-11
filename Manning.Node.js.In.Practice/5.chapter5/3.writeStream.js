const Writeable = require('stream').Writable;


class GreenStream extends Writeable {
    constructor(options){
        super(options);
    }
    _write(chunk,encoding,callback){
        process.stdout.write('\[33m\]'+chunk);
        callback
    }
}

process.stdin.pipe(new GreenStream());