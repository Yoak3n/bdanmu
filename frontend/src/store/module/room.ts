import { defineStore } from 'pinia'
import {SyncRoom} from '../../../wailsjs/go/app/App'
import type{ Room } from '@/components/types'

// 第一个参数是应用程序中 store 的唯一 id
export const useRoomStore = defineStore('room', {
  state:()=> {
      return {
          room_id: 0 as number || null,
          room_title: '',
      }
  },
  actions: {
      // 初始化
      async syncRoomId(){
        const id =  await SyncRoom()
        if (id != 0 || id != null) {
          this.room_id = id
        }else{
          const room = localStorage.getItem('room')
          if (room) {
            this.room_id = JSON.parse(room).short_id
            this.room_title = JSON.parse(room).title
          }else{
            this.room_id = null
            this.room_title = ''
          }
        }
      },
      setRoomId(id: number) {
          this.room_id = id
          localStorage.setItem('room_id', id.toString())
      },
      setRoomTitle(title: string){
        this.room_title = title
        localStorage.setItem('room_title', title)
      },
      setRoom(room:Room){
        localStorage.setItem('room', room.toString())
        this.room_id = room.short_id
        this.room_title = room.title
      }
  },
  getters: {
      getRoom: (state) => state.room_id
  }
})