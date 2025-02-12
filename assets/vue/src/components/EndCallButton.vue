<template>
    <Button variant="destructive" @click="endCall">End Meeting</Button>
</template>

<script setup>
import router from '@/router'
import Button from './ui/button/Button.vue'
import { toast } from 'vue3-toastify'

defineProps({
    call: Object,
})

const endCall = async () => {
    try {
        await call.endCall()

        router.push({ name: 'home' })
        toast.success('The meeting ended for everyone')
    } catch (error) {
        toast.error('Failed to end the meeting')
    }
}
</script>
