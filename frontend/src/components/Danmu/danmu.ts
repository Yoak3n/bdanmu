import type{User} from '../types'
export interface Danmu{
    user:User
    content:string
    type:danmuType
    room_id:number
    message_id:string
}

export enum danmuType{
    UserEntry = 1,
    Danmu = 2,
    EmoticonDanmu = 3
}


export function selectColor(level:number):danmuLevelColor{
    if (level >0 &&level < 5) {
        return danmuLevelColor.green
    } else if (level < 9) {
        return danmuLevelColor.blue
    } else if (level < 13) {
        return danmuLevelColor.purple
    }else if (level < 17) {
        return danmuLevelColor.pink
    } else {
        return danmuLevelColor.yellow
    }
}

export enum danmuLevelColor{
    green = "#42b983",
    blue = "#5d7c9b",
    purple = "#8d7aaf",
    pink = "#bc6786",
    yellow= "#c89d24"
}