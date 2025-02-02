<template>
    <nav class="border-b">
        <div class="flex h-16 items-center px-4 container mx-auto">
            <RouterLink to="/"
                class="flex items-center gap-2 font-semibold text-2xl mr-6 font-mono hover:opacity-80 transition-opacity">
                <CodeIcon class="size-8 text-emerald-500" />
                <span class="bg-gradient-to-r from-emerald-600 to-teal-500 bg-clip-text text-transparent">
                    Interview Platform
                </span>
            </RouterLink>


            <div class="flex items-center space-x-4 ml-auto">
                <DashboardButton />
                <ModeToggle />
                <UserIcon />
            </div>
        </div>
    </nav>
</template>

<script setup>
import { CodeIcon, UserIcon } from 'lucide-vue-next';
import { RouterLink } from 'vue-router';
import DashboardButton from './DashboardButton.vue';
import ModeToggle from './ModeToggle.vue';


</script>
