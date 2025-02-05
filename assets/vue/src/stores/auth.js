import axios from "axios";
import Cookies from "js-cookie";
import { defineStore } from "pinia";

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
                    this.authUser = resp.data;
                })
                .catch(err => {
                    Cookies.remove('jwt');
                    this.authUser = null;
                })
        }
    }
});