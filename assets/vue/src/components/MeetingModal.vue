<template>
    <Dialog :open="isOpen" @update:open="$emit('close')">
        <DialogContent class="sm:max-w-[425px]">
            <DialogHeader>
                <DialogTitle>{{ title }}</DialogTitle>
            </DialogHeader>

            <div class="space-y-4 pt-4">
                <Input
                    v-if="isJoinMeeting"
                    placeholder="Paste meeting link here..."
                    v-model="meetingURL"
                />

                <div class="flex justify-end gap-3">
                    <Button variant="outline" @click="$emit('close')"> Cancel </Button>
                    <Button @click="startMeeting" :disabled="isJoinMeeting && !meetingURL.trim()">
                        {{ isJoinMeeting ? 'Join Meeting' : 'Start Meeting' }}
                    </Button>
                </div>
            </div>
        </DialogContent>
    </Dialog>
</template>

<script setup>
import { ref } from 'vue';
import Button from './ui/button/Button.vue';
import Dialog from './ui/dialog/Dialog.vue';
import DialogContent from './ui/dialog/DialogContent.vue';
import DialogHeader from './ui/dialog/DialogHeader.vue';
import { Input } from './ui/input';
import router from '@/router';

const meetingURL = ref('');

const startMeeting = () => {
    if (isJoinMeeting) {
        const meetingID = meetingURL.value.split('/').pop();
        router.push(`/meeting/${meetingID}`);
    } else {
        console.log('create meeting');
    }
};

defineEmits(['close']);

const { isJoinMeeting } = defineProps({
    isOpen: Boolean,
    title: String,
    isJoinMeeting: Boolean,
});
</script>
