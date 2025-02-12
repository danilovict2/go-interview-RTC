<template>
    <div class="relative mx-auto shadow rounded-lg p-6">
        <Button @click="$emit('close')" variant="ghost" class="absolute top-4 right-4">
            <X />
        </Button>
        <h6 class="font-bold mb-4">Participants</h6>
        <ul class="divide-y">
            <li class="flex items-center py-4" v-for="(p, index) in participants" :key="index">
                <img class="h-12 w-12 rounded-full object-cover" :src="p.image" />
                <div class="ml-4">
                    <p class="font-semibold">{{ p.name }}</p>
                </div>
            </li>
        </ul>
    </div>
</template>

<script setup>
import { X } from 'lucide-vue-next';
import Button from './ui/button/Button.vue';

defineEmits(['close']);

defineProps({
    participants: Array,
});
</script>
