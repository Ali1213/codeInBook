const CountStream = require('./1.2.1.js');
const countStream = new CountStream('book');
const http = require('http');


http.get('http://www.manning.com',(res)=>{
    res.pipe(countStream);
})


countStream.on('total',(count)=>{
    console.log('total matches:', count);
})