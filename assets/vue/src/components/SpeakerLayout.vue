<template>
    <main class="flex-1 p-4 overflow-auto flex flex-col space-y-4" ref="viewport">
        <!-- Speaker -->
        <div class="rounded-lg relative">
            <CallParticipant :call="call" :participant="participants[0]" />
        </div>

        <div class="grid grid-cols-2 gap-4 md:grid-cols-4">
            <div class="rounded-lg relative" v-for="(p, index) in participants.slice(1)" :key="index">
                <CallParticipant :participant="p" :call="call" />
            </div>
        </div>
    </main>
</template>

<script setup>
import { onMounted, useTemplateRef } from 'vue';
import CallParticipant from './CallParticipant.vue';

const { call, participants } = defineProps({
    call: Object,
    participants: Array
});

const viewport = useTemplateRef('viewport');

onMounted(() => {
    call.setViewport(viewport.value);
});
</script>