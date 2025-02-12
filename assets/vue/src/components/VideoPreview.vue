<template>
    <video playsinline muted autoplay ref="preview" v-if="isCameraEnabled"></video>
    <div class="grid place-items-center" v-else>Video is disabled</div>
</template>

<script setup>
import { useTemplateRef } from 'vue';

const { camera, isCameraEnabled } = defineProps({
    camera: Object,
    isCameraEnabled: Boolean,
});

const preview = useTemplateRef('preview');

camera.state.mediaStream$.subscribe((mediaStream) => {
    if (preview.value) {
        preview.value.srcObject = mediaStream;
    }
});
</script>
