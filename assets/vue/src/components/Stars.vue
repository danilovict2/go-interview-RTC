<template>
    <div class="flex gap-0.5">
        <Star v-for="index in 5" :key="index" class="h-4 w-4"
            :class="index <= value ? 'fill-primary text-primary' : 'text-muted-foreground'" />
    </div>
</template>

<script setup>
import { Star } from 'lucide-vue-next';


const { value } = defineProps({
    value: Number,
});
</script>