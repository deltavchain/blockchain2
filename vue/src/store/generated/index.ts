// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import HackerBlog from './hacker.blog'
import HackerCreate from './hacker.create'
import HackerHack from './hacker.hack'


export default { 
  HackerBlog: load(HackerBlog, 'hacker.blog'),
  HackerCreate: load(HackerCreate, 'hacker.create'),
  HackerHack: load(HackerHack, 'hacker.hack'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}