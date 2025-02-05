<template>
    <TheNavbar></TheNavbar>
    <RouterView />
</template>

<script setup>
import { RouterView } from 'vue-router';
import TheNavbar from './components/TheNavbar.vue';
import { useAuthStore } from './stores/auth';
import { onMounted } from 'vue';

const authStore = useAuthStore();
onMounted(() => {
    authStore.loadAuthUser()
})
</script>
