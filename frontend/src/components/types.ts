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