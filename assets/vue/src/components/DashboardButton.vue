<template>
    <RouterLink to="/dashboard">
        <Button class="gap-2 font-medium" size="sm">
            <SparkleIcon class="size-4" />
            Dashboard
        </Button>
    </RouterLink>
</template>

<script setup>
import { SparkleIcon } from 'lucide-vue-next';
import { RouterLink } from 'vue-router';
import Button from './ui/button/Button.vue';
</script>
