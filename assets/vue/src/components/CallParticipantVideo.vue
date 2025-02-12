<template>
    <video class="object-contain h-[300px]" ref="video"></video>
</template>

<script setup>
import { onMounted, onUnmounted, useTemplateRef } from 'vue';

const { call, participant } = defineProps({
    call: Object,
    participant: Object,
});

const video = useTemplateRef('video');
let untrack = null;
let unbind = null;

onMounted(() => {
    untrack = call.trackElementVisibility(video.value, participant.sessionId, 'videoTrack');

    unbind = call.bindVideoElement(video.value, participant.sessionId, 'videoTrack');
});

onUnmounted(() => {
    if (untrack) {
        untrack();
    }

    if (unbind) {
        unbind();
    }
});
</script>
