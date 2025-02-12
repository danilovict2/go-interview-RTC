<template>
    <CallParticipantVideo v-if="hasVideo(participant)" :call="call" :participant="participant" />
    <div class="grid place-items-center border border-border h-[300px]" v-else>Video is disabled</div>
    <div
        class="absolute border border-border bottom-0 left-0 right-0 bg-background bg-opacity-50 text-sm p-1 text-center rounded-b-lg">
        {{ participant.name }}
    </div>

    <audio ref="audio" :muted="participant.isLocalParticipant || !hasAudio(participant)"></audio>
</template>

<script setup>
import { hasAudio, hasVideo } from '@stream-io/video-client';
import CallParticipantVideo from './CallParticipantVideo.vue';
import { onMounted, onUnmounted, useTemplateRef } from 'vue';

const { call, participant } = defineProps({
    call: Object,
    participant: Object,
});

const audio = useTemplateRef('audio');
let unbind = null;

onMounted(() => {
    unbind = (call.bindAudioElement(audio.value, participant.sessionId));
});

onUnmounted(() => {
    if (unbind) {
        unbind();
    }
});
</script>