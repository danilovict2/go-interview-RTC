<template>
    <Button variant="destructive" @click="endCall">End Meeting</Button>
</template>

<script setup>
import router from '@/router';
import Button from './ui/button/Button.vue';
import { toast } from 'vue3-toastify';
import axios from 'axios';
import Cookies from 'js-cookie';

const { call } = defineProps({
    call: Object,
});

const endCall = async () => {
    try {
        await axios.patch(
            `/interviews/${call.id}/end`,
            {
                endTime: new Date().toUTCString(),
            },
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    Authorization: `Bearer ${Cookies.get('jwt')}`,
                },
            },
        );
        await call.endCall();

        router.push({ name: 'Home' });
        toast.success('The meeting ended for everyone');
    } catch (error) {
        console.log(error);
        toast.error('Failed to end the meeting');
    }
};
</script>
