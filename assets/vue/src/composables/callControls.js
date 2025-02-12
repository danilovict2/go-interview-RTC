import { ref, watch } from "vue";

export function useCallControlls(call) {
    const isCameraEnabled = ref(call.camera.state.status === 'enabled');
    const isMicEnabled = ref(call.microphone.state.status === 'enabled');

    if (isCameraEnabled.value) {
        call.camera.enable();
    } else {
        call.camera.disable();
    }

    if (isMicEnabled.value) {
        call.microphone.enable();
    } else {
        call.microphone.disable();
    }

    watch(isCameraEnabled, () => {
        call.camera.toggle();
    })

    watch(isMicEnabled, () => {
        call.microphone.toggle();
    })

    return { isCameraEnabled, isMicEnabled };
}