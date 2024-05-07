import {OpenWindow} from '../../../wailsjs/go/app/App'
export const JumpToLiveRoom =(e:MouseEvent)=> {
    e.preventDefault()
    const room_id:string = localStorage.getItem('room_id')!
    OpenWindow('https://live.bilibili.com/'+room_id)
}