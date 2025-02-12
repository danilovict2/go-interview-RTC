import { defineStore } from 'pinia'
import { StreamVideoClient } from '@stream-io/video-client'
import { useAuthStore } from './auth'
import axios from 'axios'
import Cookies from 'js-cookie'

export const useStreamStore = defineStore('stream', {
    state: () => {
        const apiKey = import.meta.env.VITE_STREAM_API_KEY
        const authUser = useAuthStore().authUser
        const user = {
            id: authUser.uuid,
            name: authUser.first_name + ' ' + authUser.last_name,
            image: `https://api.dicebear.com/9.x/initials/svg/seed=${authUser.first_name}-${authUser.last_name}`,
        }

        const tokenProvider = async () => {
            const { data } = await axios.get('/stream/token', {
                headers: {
                    Authorization: `Bearer ${Cookies.get('jwt')}`,
                },
            })

            return data.token
        }

        return {
            client: new StreamVideoClient({
                apiKey, 
                tokenProvider, 
                user, 
            }),
        }
    },
})
