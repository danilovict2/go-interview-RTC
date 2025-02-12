import axios from 'axios';
import Cookies from 'js-cookie';
import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => {
        return {
            authUser: null,
        };
    },

    actions: {
        async loadAuthUser() {
            try {
                const { data } = await axios.get('/users/me', {
                    headers: {
                        Authorization: `Bearer ${Cookies.get('jwt')}`,
                    },
                });

                this.authUser = data;
            } catch (err) {
                Cookies.remove('jwt');
                this.authUser = null;
            }
        },
    },
});
