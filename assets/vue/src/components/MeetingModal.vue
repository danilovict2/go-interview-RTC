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
import axios from 'axios';
import { toast } from 'vue3-toastify';
import Cookies from 'js-cookie';
import { useAuthStore } from '@/stores/auth';

defineEmits(['close']);

const { isJoinMeeting } = defineProps({
    isOpen: Boolean,
    title: String,
    isJoinMeeting: Boolean,
});

const meetingURL = ref('');

const startMeeting = () => {
    if (isJoinMeeting) {
        const meetingID = meetingURL.value.split('/').pop();
        router.push({ name: 'Meeting', params: { id: meetingID } });
    } else {
        axios
            .post(
                '/interviews/store',
                {
                    title: 'Instant Meeting',
                    description: 'Instant Meeting',
                    startTime: new Date().toUTCString(),
                    attendeeUUIDs: JSON.stringify([useAuthStore().authUser.uuid]),
                },
                {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                        Authorization: `Bearer ${Cookies.get('jwt')}`,
                    },
                },
            )
            .then((resp) => {
                router.push({ name: 'Meeting', params: { id: resp.data.stream_call_id } });
            })
            .catch((err) => toast.error(err.response.data.message));
    }
};
</script>
