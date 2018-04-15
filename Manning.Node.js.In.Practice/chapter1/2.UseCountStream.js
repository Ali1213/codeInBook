const CountStream = require('./1.CountStream.js');
const countStream = new CountStream('book');
const https = require('https');


https.get('https://www.manning.com',(res)=>{
    res.pipe(countStream);
})


countStream.on('total',(count)=>{
    console.log('total matches:', count);
})