import type {User} from "../user";


export interface SuperChat{
    user:User
    content:string
    room_id:number
    message_id:string
    timestamp:number
    end_time:number
}