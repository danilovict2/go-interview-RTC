<template>
    <Card>
        <CardHeader class="space-y-2">
            <div class="flex items-center justify-between">
                <div class="flex items-center gap-2 text-sm text-muted-foreground">
                    <CalendarIcon class="h-4 w-4" />
                    {{ formattedStartTime }}
                </div>

                <Badge :variant="status === 'live'
                    ? 'default'
                    : status === 'upcoming'
                        ? 'secondary'
                        : 'outline'
                    ">
                    {{
                        status === 'live'
                            ? 'Live Now'
                            : status === 'upcoming'
                                ? 'Upcoming'
                                : 'Completed'
                    }}
                </Badge>
            </div>

            <CardTitle>{{ interview.title }}</CardTitle>

            <CardDescription v-show="interview.description !== ''" class="line-clamp-2">{{ interview.description }}
            </CardDescription>
        </CardHeader>

        <CardContent>
            <Button v-if="status === 'live'" class="w-full"
                @click="router.push({ name: 'Meeting', params: { id: interview.stream_call_id } });">
                Join Meeting
            </Button>

            <Button v-else-if="status === 'upcoming'" :variant="authUser.role === 'candidate' ? 'outline' : 'default'"
                class="w-full" :disabled="authUser.role === 'candidate'" @click="startInterview">
                {{ authUser.role === 'candidate' ? 'Waiting to Start' : 'Start Interview' }}
            </Button>
        </CardContent>
    </Card>
</template>

<script setup>
import { CalendarIcon } from 'lucide-vue-next';
import Card from './ui/card/Card.vue';
import CardHeader from './ui/card/CardHeader.vue';
import Badge from './ui/badge/Badge.vue';
import CardTitle from './ui/card/CardTitle.vue';
import CardDescription from './ui/card/CardDescription.vue';
import CardContent from './ui/card/CardContent.vue';
import Button from './ui/button/Button.vue';
import { format } from 'date-fns';
import { ref } from 'vue';
import axios from 'axios';
import { toast } from 'vue3-toastify';
import Cookies from 'js-cookie';
import router from '@/router';
import { useAuthStore } from '@/stores/auth';


const authUser = useAuthStore().authUser;

const { interview } = defineProps({
    interview: Object,
});

const startTime = new Date(interview.start_time);
const status = ref(interview.status);
const formattedStartTime = format(startTime, 'MMM d, yyyy, hh:mm a');


const startInterview = () => {
    status.value = 'live';
    axios.patch(
        `/interviews/${interview.stream_call_id}/mark_live`,
        null,
        {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                Authorization: `Bearer ${Cookies.get('jwt')}`,
            },
        },
    )
        .then(() => {
            router.push({ name: 'Meeting', params: { id: interview.stream_call_id } })
        })
        .catch(err => toast.error(err));
};
</script>
