<template>
    <div class="min-h-screen flex items-center justify-center p-6 bg-background/95">
        <div class="w-full max-w-[1200px] mx-auto">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <Card class="md:col-span-1 p-6 flex flex-col">
                    <div>
                        <h1 class="text-xl font-semibold mb-1">Camera Preview</h1>
                        <p class="text-sm text-muted-foreground">Make sure you look good!</p>
                    </div>

                    <div class="mt-4 flex-1 min-h-[400px] rounded-xl overflow-hidden bg-muted/50 border relative">
                        <div class="absolute inset-0">
                            <VideoPreview :camera="call.camera" :is-camera-enabled="isCameraEnabled"
                                :microphone="call.microphone" class="h-full w-full" />
                        </div>
                    </div>
                </Card>

                <Card class="md:col-span-1 p-6">
                    <div class="h-full flex flex-col">
                        <div>
                            <h2 class="text-xl font-semibold mb-1">Meeting Details</h2>
                            <p class="text-sm text-muted-foreground break-all">{{ call.id }}</p>
                        </div>

                        <div class="flex-1 flex flex-col justify-between">
                            <div class="spacey-6 mt-8">
                                <div class="flex items-center justify-between">
                                    <div class="flex items-center gap-3">
                                        <div
                                            class="h-10 w-10 rounded-full bg-primary/10 flex items-center justify-center">
                                            <Camera class="h-5 w-5 text-primary" />
                                        </div>
                                        <div>
                                            <p class="font-medium">Camera</p>
                                            <p class="text-sm text-muted-foreground">
                                                {{ isCameraEnabled ? 'On' : 'Off' }}
                                            </p>
                                        </div>
                                    </div>
                                    <Switch :checked="isCameraEnabled"
                                        @update:checked="isCameraEnabled = !isCameraEnabled" />
                                </div>

                                <div class="flex items-center justify-between">
                                    <div class="flex items-center gap-3">
                                        <div
                                            class="h-10 w-10 rounded-full bg-primary/10 flex items-center justify-center">
                                            <Mic class="h-5 w-5 text-primary" />
                                        </div>
                                        <div>
                                            <p class="font-medium">Microphone</p>
                                            <p class="text-sm text-muted-foreground">
                                                {{ isMicEnabled ? 'On' : 'Off' }}
                                            </p>
                                        </div>
                                    </div>
                                    <Switch :checked="isMicEnabled" @update:checked="isMicEnabled = !isMicEnabled" />
                                </div>
                                <div class="flex items-center justify-between">
                                    <div class="flex items-center gap-3">
                                        <div
                                            class="h-10 w-10 rounded-full bg-primary/10 flex items-center justify-center">
                                            <Settings class="h-5 w-5 text-primary" />
                                        </div>
                                        <div>
                                            <p class="font-medium">Settings</p>
                                            <p class="text-sm text-muted-foreground">Configure devices</p>
                                        </div>
                                    </div>
                                    <DeviceSettings :call="call" />
                                </div>
                            </div>
                            <div class="space-y-3 mt-8">
                                <Button class="w-full" size="lg" @click.once="handleJoin">Join Meeting</Button>
                                <p class="text-xs text-center text-muted-foreground">
                                    Do not worry, our team is super friendly! We want you to succeed. 🎉
                                </p>
                            </div>
                        </div>
                    </div>
                </Card>
            </div>
        </div>
    </div>
</template>

<script setup>
import { Camera, Mic, Settings } from 'lucide-vue-next'
import Card from './ui/card/Card.vue'
import Switch from './ui/switch/Switch.vue'
import Button from './ui/button/Button.vue'
import VideoPreview from './VideoPreview.vue'
import DeviceSettings from './DeviceSettings.vue'
import { useCallControlls } from '@/composables/callControls'

const emit = defineEmits(['setup-complete']);

const { call } = defineProps({
    call: Object,
});

const { isCameraEnabled, isMicEnabled } = useCallControlls(call);

const handleJoin = async () => {
    await call.join()
    emit('setup-complete')
}
</script>
