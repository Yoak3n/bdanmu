import {OpenWindow} from '../../../wailsjs/go/app/App'
import {Quit,WindowMinimise} from '../../../wailsjs/runtime'
export const JumpToLiveRoom =(e:MouseEvent)=> {
    e.preventDefault()
    const room_id:string = localStorage.getItem('room_id')!
    OpenWindow('https://live.bilibili.com/' + room_id).then(()  =>{})
}
export const AppQuit = ()=>{
    Quit()
}

export const HideWindow =(e:MouseEvent)=> {
    e.preventDefault()
    WindowMinimise()
}

