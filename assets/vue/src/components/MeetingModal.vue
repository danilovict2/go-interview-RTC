<template>
    <Dialog :open="isOpen" @update:open="$emit('close')">
        <DialogContent class="sm:max-w-[425px]">
            <DialogHeader>
                <DialogTitle>{{ title }}</DialogTitle>
            </DialogHeader>

            <div class="space-y-4 pt-4">
                <Input v-if="isJoinMeeting" placeholder="Paste meeting link here..." value=""
                    @change="console.log('new meeting url: ', e.target.value)" />


                <div class="flex justify-end gap-3">
                    <Button variant="outline" @click="$emit('close')">
                        Cancel
                    </Button>
                    <Button @click="console.log('start meeting')" :disabled="isJoinMeeting">
                        {{ isJoinMeeting ? "Join Meeting" : "Start Meeting" }}
                    </Button>
                </div>
            </div>
        </DialogContent>
    </Dialog>
</template>

<script setup>
import Button from './ui/button/Button.vue';
import Dialog from './ui/dialog/Dialog.vue';
import DialogContent from './ui/dialog/DialogContent.vue';
import DialogHeader from './ui/dialog/DialogHeader.vue';
import { Input } from './ui/input';

defineEmits(['close'])

defineProps({
    isOpen: Boolean,
    title: String,
    isJoinMeeting: Boolean,
})
</script>