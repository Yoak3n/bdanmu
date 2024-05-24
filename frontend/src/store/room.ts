import { defineStore } from 'pinia'
import {SyncAuth} from '../../wailsjs/go/app/App'

// 第一个参数是应用程序中 store 的唯一 id
export const useRoomsStore = defineStore('room', {
  state:()=> {
      return {
          room_id: 0,
          room_title: '',
          cooke: '',
          user: '',
      }
  },
  actions: {
      setRoomId(id: number) {
          this.room_id = id
          localStorage.setItem('room_id', id.toString())
      },
      setRoomTitle(title: string){
        this.room_title = title
        localStorage.setItem('room_title', title)
      },
      setCookie(cookie: string) {
          this.cooke = cookie
          localStorage.setItem('cookie', cookie)
      },
      setToken(token: string){
          this.user = token
          localStorage.setItem('token', token)
      },
      async syncAuth (){
        const auth = await SyncAuth()
        if (auth[0] != '' || auth[1] != '') {
          this.setCookie(auth[0])
          this.setToken(auth[1])
        }
      }
  },
  getters: {
      getRoom: (state) => state.room_id
  }
})