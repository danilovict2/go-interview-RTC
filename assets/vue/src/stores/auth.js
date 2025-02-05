import axios from "axios";
import Cookies from "js-cookie";
import { defineStore } from "pinia";
import { toast } from "vue3-toastify";

export const useAuthStore = defineStore('auth', {
    state: () => {
        return {
            authUser: null,
        }
    },

    actions: {
        loadAuthUser() {
            axios.get("/users/me", {
                headers: {
                    "Authorization": `Bearer ${Cookies.get('jwt')}`
                }
            })
                .then(resp => {
                    this.authUser = resp.data
                })
                .catch(err => {
                    toast.error(err)
                })
        }
    }
});