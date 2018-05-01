const fs = require('fs');
const Transform = require('stream').Transform;

class CSVParser extends Transform {
    constructor(options){
        super(options);
        this.value = '';
        this.headers = [];
        this.values = [];
        this.line = 0;
    }
    _transform(chunk,encoding,done){
        let c;
        let i;
        chunk = chunk.toString();
        console.log(chunk)
        for(let v of chunk){
            if(v === ','){
                this.addValue();
            }else if(v === 'n'){
                this.addValue();
                if(this.line>0){
                    this.push(JSON.stringify(this.toObject()));
                }
                this.values = [];
                this.lines++;
            }else{
                this.value += v;
            }
        }
        done();
    }
    toObject(){
        let i;
        let obj = {};
        for(let [key,header] of this.headers.entries()){
            obj[header] = this.values[key]
        }
        return obj;
    }
    addValue(){
        if(this.line === 0){
            this.headers.push(this.value);
        }else{
            this.values.push(this.value);
        }
        this.values = ''; 
    }
}


var parser = new CSVParser();
fs.createReadStream(__dirname +'/sample.csv').pipe(parser).pipe(process.stdout);