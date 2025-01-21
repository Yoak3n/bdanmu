export interface User{
    uid:number
    name:string
    sex:number
    guard:boolean
    avatar:string
    fans_count?:number
    medal?:Medal
}

interface Medal {
    name:string,
    owner_id:number,
    level:number,
    target_id:number
}

export interface Room{
    short_id:number,
    user?:User,
    title:string,
    cover:string,
    long_id:number,
    follower_count:number,
}

export enum MessageType {
    SuperChat = 0,
    UserEntry = 1,
    UserInfo = 2,
    Danmu = 3,
}

import type {Danmu} from '../../src/components/Danmu/danmu'
import type {SuperChat} from '../../src/components/SuperChat/super_chat'
export interface Message {
    type: MessageType,
    data: Danmu|SuperChat,
}