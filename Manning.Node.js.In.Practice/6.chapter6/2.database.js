const fs = require('fs')
const EventEmitter = require('events').EventEmitter
let p,private
p = private = {
    records: Symbol("records"),
    writeStream: Symbol('write stream'),
    load:Symbol('load')
}


class Database extends EventEmitter {
    constructor(filepath){
        super()
        this.path = filepath
        this[p.records] = Object.create(null)
        this[p.writeStream] = fs.createWriteStream(this.path,{
            encoding:'utf8',
            flags:'a',
        })
        this[p.load]()
    }
    [p.load](){
        const stream = fs.createReadStream(this.path,{encoding:'utf8'})

        let data = ''

        stream.on('readable',()=>{
            data += stream.read()
            let records = data.split('\n')
            data = records.pop()
            records.forEach(item => {
                let record
                try {
                    record = JSON.parse(item)
                    if( record.value == null ){
                        delete this[p.records][record.key]
                    }else{
                        this[p.records][record.key] = record.value
                    }
                }catch(e){
                    this.emit("error",'found invalid records:', item)
                }
            })

        })

        stream.on("end",()=>{
            this.emit('load')
        })
    }
}


// const db = new Database('./111')

// console.log(db[p.records])