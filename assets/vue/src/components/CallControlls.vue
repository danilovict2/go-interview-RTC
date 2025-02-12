<template>
    <div class="flex items-center gap-2 flex-wrap justify-center px-4">
        <Button @click="isMicEnabled = !isMicEnabled" :variant="isMicEnabled ? 'outline' : 'destructive'">
            <Mic v-if="isMicEnabled" />
            <MicOff v-else />
        </Button>
        <Button @click="isCameraEnabled = !isCameraEnabled" :variant="isCameraEnabled ? 'outline' : 'destructive'">
            <Camera v-if="isCameraEnabled" />
            <CameraOff v-else />
        </Button>
        <Button variant="destructive" @click="$emit('leave')">
            <PhoneForwarded />
        </Button>
    </div>
</template>

<script setup>
import { useCallControlls } from '@/composables/callControls';
import { Camera, CameraOff, Mic, MicOff, PhoneForwarded } from 'lucide-vue-next';
import Button from './ui/button/Button.vue';

defineEmits(['leave']);

const { call } = defineProps({
    call: Object
});

const { isCameraEnabled, isMicEnabled } = useCallControlls(call);
</script>