import type {User} from "../types";


export interface SuperChat{
    user:User
    content:string
    room_id:number
    message_id:string
    timestamp:number
    end_time:number
}