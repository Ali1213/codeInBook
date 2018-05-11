const Event = require('events')
class EventTracker extends Event {
    constructor(){
        super();
    }
}

const eventTracker = new EventTracker();

eventTracker.on('newListener',(name,listener)=>{
    console.log('Event name has been add:',name);
    console.log(listener(111))
})

eventTracker.on('a listener',(aaa)=>aaa)