<template>
    <DropdownMenu>
        <div
            class="h-10 w-10 rounded-full bg-primary/10 flex items-center justify-center cursor-pointer hover:bg-accent"
        >
            <DropdownMenuTrigger as-child>
                <Settings class="h-5 w-5 text-primary" />
            </DropdownMenuTrigger>
        </div>
        <DropdownMenuContent class="w-56">
            <DropdownMenuLabel>Select a Camera</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuRadioGroup v-model="selectedCameraID">
                <DropdownMenuRadioItem
                    v-for="(cam, index) in cameras"
                    :key="index"
                    :value="cam.deviceId"
                >
                    {{ cam.label }}
                </DropdownMenuRadioItem>
            </DropdownMenuRadioGroup>

            <DropdownMenuLabel>Select a Mic</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuRadioGroup v-model="selectedMicID">
                <DropdownMenuRadioItem
                    v-for="(mic, index) in microphones"
                    :key="index"
                    :value="mic.deviceId"
                >
                    {{ mic.label }}
                </DropdownMenuRadioItem>
            </DropdownMenuRadioGroup>

            <DropdownMenuLabel>Select a Speaker</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuRadioGroup v-model="selectedSpeakerID">
                <DropdownMenuRadioItem
                    v-for="(speaker, index) in speakers"
                    :key="index"
                    :value="speaker.deviceId"
                >
                    {{ speaker.label }}
                </DropdownMenuRadioItem>
            </DropdownMenuRadioGroup>
        </DropdownMenuContent>
    </DropdownMenu>
</template>

<script setup>
import { Settings } from 'lucide-vue-next';
import DropdownMenu from './ui/dropdown-menu/DropdownMenu.vue';
import DropdownMenuContent from './ui/dropdown-menu/DropdownMenuContent.vue';
import DropdownMenuLabel from './ui/dropdown-menu/DropdownMenuLabel.vue';
import DropdownMenuTrigger from './ui/dropdown-menu/DropdownMenuTrigger.vue';
import { ref, watch } from 'vue';
import DropdownMenuSeparator from './ui/dropdown-menu/DropdownMenuSeparator.vue';
import DropdownMenuRadioGroup from './ui/dropdown-menu/DropdownMenuRadioGroup.vue';
import DropdownMenuRadioItem from './ui/dropdown-menu/DropdownMenuRadioItem.vue';

const { call } = defineProps({
    call: Object,
});

const cameras = ref([]);
const microphones = ref([]);
const speakers = ref([]);

const selectedCameraID = ref(null);
const selectedMicID = ref(null);
const selectedSpeakerID = ref(null);

watch(selectedCameraID, (newID) => {
    call.camera.select(newID);
});

watch(selectedMicID, (newID) => {
    call.microphone.select(newID);
});

watch(selectedSpeakerID, (newID) => {
    console.log(newID);
    call.speaker.select(newID);
});

call.camera.listDevices().subscribe({
    next: (devices) => (cameras.value = devices),
    error: (err) => console.error(err),
});

call.speaker.listDevices().subscribe({
    next: (devices) => (microphones.value = devices),
    error: (err) => console.error(err),
});

call.speaker.listDevices().subscribe({
    next: (devices) => (speakers.value = devices),
    error: (err) => console.error(err),
});
</script>
