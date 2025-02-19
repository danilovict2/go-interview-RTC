<template>
    <AppLoading v-if="isLoading" />
    <Dialog :open="isOpen" @update:open="isOpen = !isOpen">
        <DialogTrigger asChild>
            <Button variant="secondary" class="w-full">
                <MessageSquare class="h-4 w-4 mr-2" />
                Add Comment
            </Button>
        </DialogTrigger>

        <DialogContent class="sm:max-w-[600px]">
            <DialogHeader>
                <DialogTitle>Interview Comment</DialogTitle>
            </DialogHeader>

            <div class="space-y-6">
                <div class="space-y-4" v-if="comments.length > 0">
                    <div class="flex items-center justify-between">
                        <h4 class="text-sm font-medium">Previous Comments</h4>
                        <Badge variant="outline">
                            {{ comments.length }} Comment{{ comments.length !== 1 ? 's' : '' }}
                        </Badge>
                    </div>

                    <ScrollArea class="h-[240px]">
                        <div class="space-y-4">
                            <div
                                v-for="comment in comments"
                                :key="comment.id"
                                class="rounded-lg border p-4 space-y-3"
                            >
                                <div class="flex items-center justify-between">
                                    <div class="flex items-center gap-2">
                                        <Avatar class="h-8 w-8">
                                            <AvatarImage
                                                :src="`https://api.dicebear.com/9.x/initials/svg/seed=${comment.created_by.first_name}-${comment.created_by.last_name}`"
                                            />
                                            <AvatarFallback>
                                                <UserCircle class="h-6 w-6" />
                                            </AvatarFallback>
                                        </Avatar>
                                        <div>
                                            <p class="text-sm font-medium">
                                                {{
                                                    comment.created_by.first_name +
                                                    ' ' +
                                                    comment.created_by.last_name
                                                }}
                                            </p>
                                            <p class="text-xs text-muted-foreground">
                                                {{
                                                    format(
                                                        comment.CreatedAt,
                                                        'MMM d, yyyy • h:mm a',
                                                    )
                                                }}
                                            </p>
                                        </div>
                                    </div>
                                    <StarRating :value="comment.rating" />
                                </div>
                                <p class="text-sm text-muted-foreground">{{ comment.content }}</p>
                            </div>
                        </div>
                    </ScrollArea>
                </div>

                <div class="space-y-4">
                    <div class="space-y-2">
                        <Label>Rating</Label>
                        <Select v-model="rating">
                            <SelectTrigger>
                                <SelectValue placeholder="Select rating" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem v-for="index in 5" :key="index" :value="index">
                                    <div class="flex items-center gap-2">
                                        <StarRating :value="index" />
                                    </div>
                                </SelectItem>
                            </SelectContent>
                        </Select>
                    </div>

                    <div class="space-y-2">
                        <Label>Your Comment</Label>
                        <Textarea
                            v-model="comment"
                            placeholder="Share your detailed comment about the candidate..."
                            class="h-32"
                        />
                    </div>
                </div>
            </div>

            <DialogFooter>
                <Button variant="outline" @click="isOpen = false"> Cancel </Button>
                <Button @click="handleSubmit">Submit</Button>
            </DialogFooter>
        </DialogContent>
    </Dialog>
</template>

<script setup>
import { MessageSquare } from 'lucide-vue-next';
import Button from './ui/button/Button.vue';
import Dialog from './ui/dialog/Dialog.vue';
import DialogTrigger from './ui/dialog/DialogTrigger.vue';
import DialogContent from './ui/dialog/DialogContent.vue';
import DialogHeader from './ui/dialog/DialogHeader.vue';
import DialogTitle from './ui/dialog/DialogTitle.vue';
import { ref } from 'vue';
import Badge from './ui/badge/Badge.vue';
import ScrollArea from './ui/scroll-area/ScrollArea.vue';
import Avatar from './ui/avatar/Avatar.vue';
import AvatarImage from './ui/avatar/AvatarImage.vue';
import AvatarFallback from './ui/avatar/AvatarFallback.vue';
import Label from './ui/label/Label.vue';
import Select from './ui/select/Select.vue';
import SelectTrigger from './ui/select/SelectTrigger.vue';
import SelectValue from './ui/select/SelectValue.vue';
import SelectContent from './ui/select/SelectContent.vue';
import SelectItem from './ui/select/SelectItem.vue';
import Textarea from './ui/textarea/Textarea.vue';
import { DialogFooter } from './ui/dialog';
import { format } from 'date-fns';
import axios from 'axios';
import { toast } from 'vue3-toastify';
import Cookies from 'js-cookie';
import StarRating from './StarRating.vue';
import AppLoading from './AppLoading.vue';

const { interviewID } = defineProps({
    interviewID: String,
});

const isOpen = ref(false);
const comments = ref([]);
const rating = ref(1);
const comment = ref('');
const isLoading = ref(false);

const handleSubmit = () => {
    isLoading.value = true;
    axios
        .post(
            '/comments/store',
            {
                content: comment.value,
                rating: rating.value,
                interviewID: interviewID,
            },
            {
                headers: {
                    Authorization: `Bearer ${Cookies.get('jwt')}`,
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
            },
        )
        .then((resp) => {
            rating.value = 1;
            comment.value = '';
            comments.value.push(resp.data.comment);
        })
        .catch((err) => {
            console.error('Failed to add comment:', err);
            toast.error('Failed to add comment!');
        })
        .finally(() => (isLoading.value = false));
};

const getComments = (interviewID) => {
    isLoading.value = true;
    axios
        .get(`/interviews/${interviewID}/comments`, {
            headers: {
                Authorization: `Bearer ${Cookies.get('jwt')}`,
            },
        })
        .then((resp) => {
            comments.value = resp.data.comments;
        })
        .catch((err) => {
            console.error("Couldn't load comments:", err);
            toast.error("Couldn't load comments!");
        })
        .finally(() => (isLoading.value = false));
};

getComments(interviewID);
</script>
