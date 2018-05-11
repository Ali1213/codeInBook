const fs = require('fs')
const path = require('path')

function findSync(filenameRe, startPath){
    if( !(filenameRe instanceof RegExp) ) throw Error('filenameRe must be a regexp')
    if( !(startPath instanceof string)) throw Error('startPath must be a string')
    const result = []

    function finder(filepath){
        const files = fs.readdirSync(filepath)

        for(let file of files){
            const fPath = path.join(filepath,file)
            const stat = fs.statSync(fPath)

            if(stat.isDirectory()) finder(fPath)
            if(stat.isFile() && filenameRe.test(fPath)) result.push(fPath)
        }
    }
    const stat = fs.statSync(startPath)
    if(stat.isDirectory()) finder(startPath)
    if(stat.isFile() && filenameRe.test(startPath)) result.push(startPath)
    return result
}


function find(filenameRe, startPath, cb){
    if( !(filenameRe instanceof RegExp) ) return cb(Error('filenameRe must be a regexp'))
    if( !(startPath instanceof string)) return cb(Error('startPath must be a string'))
    const results = []
    let asyncOps = 0

    function finder(filepath){
        asyncOps++;
        fs.readdir(filepath,(err,files)=>{
            if(err) return cb(err)

            files.forEach(file=>{
                const fPath = path.join(filepath, file)
                asyncOps++;
                fs.stat(fPath,(err, stat)=>{
                    if(err) return cb(err)
                    if(stat.isDirectory()) return finder(fPath)
                    if(stat.isFile() && filenameRe.test(fPath)) results.push(fPath)
                    asyncOps--;
                    if(asyncOps === 0)cb(null,results)
                })
            })
            asyncOps--;
            if(asyncOps === 0) cb(null,results)
        })
    }

}

exports.findSync = findSync;

exports.find = find;