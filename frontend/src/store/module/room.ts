import { defineStore } from 'pinia'
import {SyncRoom} from '../../../wailsjs/go/app/App'

// 第一个参数是应用程序中 store 的唯一 id
export const useRoomStore = defineStore('room', {
  state:()=> {
      return {
          room_id: 0 as number || null,
          room_title: '',

      }
  },
  actions: {
      async syncRoomId(){
        const id =  await SyncRoom()
        if (id != 0 || id != null) {
          this.room_id = id
        }else{
          this.room_id = null
        }
      },
      setRoomId(id: number) {
          this.room_id = id
          localStorage.setItem('room_id', id.toString())
      },
      setRoomTitle(title: string){
        this.room_title = title
        localStorage.setItem('room_title', title)
      }
  },
  getters: {
      getRoom: (state) => state.room_id
  }
})