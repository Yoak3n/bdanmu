import {HideToTray} from '../../../wailsjs/go/app/App'
import {Quit,BrowserOpenURL,WindowSetAlwaysOnTop} from '../../../wailsjs/runtime'

export const JumpToLiveRoom =(e:MouseEvent)=> {
    e.preventDefault()
    const room_id:string = localStorage.getItem('room_id')!
    BrowserOpenURL('https://live.bilibili.com/' + room_id)
    // OpenWindow('https://live.bilibili.com/' + room_id).then(()  =>{})
}


export const AppQuit = ()=>{
    Quit()
}

export const HideWindow =(e:MouseEvent)=> {
    e.preventDefault()
    HideToTray()
}

export const TopWindow = (e:MouseEvent,flag:boolean)=> {
    e.preventDefault();
    WindowSetAlwaysOnTop(!flag);
    console.log(!flag);
    
}
