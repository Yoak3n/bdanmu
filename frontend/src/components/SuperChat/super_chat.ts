import type {User} from "../types";


export interface SuperChat{
    user:User
    content:string
    room_id:number
    message_id:string
    timestamp:number
    end_time:number
    price:number
}

function computeSuperChatBackground(price :number):string{
    switch (price){
        case 30:
            return "bg-[#F8D766]"
        case 50:
            return "bg-[#F8D766]"
    }
    return "aa"
}