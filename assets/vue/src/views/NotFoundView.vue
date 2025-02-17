<template>
    <div class="py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6">
        <div class="mx-auto max-w-screen-sm text-center">
            <h1
                class="mb-4 text-7xl tracking-tight font-extrabold lg:text-9xl text-primary-600 dark:text-primary-500"
            >
                404
            </h1>
            <p class="mb-4 text-lg font-light text-primary-500 dark:text-primary-400">
                Sorry, we can't find that page.
            </p>
            <RouterLink to="/">
                <Button>Back to Homepage</Button>
            </RouterLink>
        </div>
    </div>
</template>

<script setup>
import Button from '@/components/ui/button/Button.vue';
import { RouterLink } from 'vue-router';
</script>
