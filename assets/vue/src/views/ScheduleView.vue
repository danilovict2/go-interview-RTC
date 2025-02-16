<template>
    <div class="container max-w-7xl mx-auto p-6 space-y-8">
        <div class="flex items-center justify-between">
            <div>
                <h1 class="text-3xl font-bold">Interviews</h1>
                <p class="text-muted-foreground mt-1">Schedule and manage interviews</p>
            </div>

            <Dialog :open="open" @update:open="open = !open">
                <DialogTrigger asChild>
                    <Button size="lg">Schedule an Interview</Button>
                </DialogTrigger>

                <DialogContent class="sm:max-w-[500px] h-[calc(100vh-200px)] overflow-auto">
                    <DialogHeader>
                        <DialogTitle>Schedule an Interview</DialogTitle>
                    </DialogHeader>
                    <div class="space-y-4 py-4">
                        <div class="space-y-2">
                            <label class="text-sm font-medium">Title</label>
                            <Input placeholder="Interview title" v-model="formData.title" />
                        </div>

                        <div class="space-y-2">
                            <label class="text-sm font-medium">Description</label>
                            <Textarea
                                placeholder="Interview description"
                                v-model="formData.desription"
                                :rows="3"
                            />
                        </div>

                        <div class="space-y-2">
                            <label class="text-sm font-medium">Candidate</label>
                            <Select v-model="formData.candidateUUID">
                                <SelectTrigger>
                                    <SelectValue placeholder="Select candidate" />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem
                                        v-for="candidate in candidates"
                                        :key="candidate.uuid"
                                        :value="candidate.uuid"
                                    >
                                        <UserInfo :user="candidate" />
                                    </SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <div class="space-y-2">
                            <label class="text-sm font-medium">Interviewers</label>
                            <div class="flex flex-wrap gap-2 mb-2">
                                <div
                                    v-for="interviewer in selectedInterviewers"
                                    :key="interviewer.uuid"
                                    class="inline-flex items-center gap-2 bg-secondary px-2 py-1 rounded-md text-sm"
                                >
                                    <UserInfo :user="interviewer" />
                                    <button
                                        @click="removeInterviewer(interviewer.uuid)"
                                        class="hover:text-destructive transition-colors"
                                        v-show="interviewer.uuid !== authUser.uuid"
                                    >
                                        <X class="h-4 w-4" />
                                    </button>
                                </div>
                            </div>

                            <Select
                                @update:modelValue="($event) => addInterviewer($event)"
                                v-show="availableInterviewers.length > 0"
                            >
                                <SelectTrigger>
                                    <SelectValue placeholder="Add interviewer" />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem
                                        v-for="interviewer in availableInterviewers"
                                        :key="interviewer.uuid"
                                        :value="interviewer.uuid"
                                    >
                                        <UserInfo :user="interviewer" />
                                    </SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <div class="flex gap-4">
                            <div class="space-y-2">
                                <label class="text-sm font-medium">Date</label>
                                <Calendar
                                    :multiple="false"
                                    v-model="formData.date"
                                    :disabled="formData.date < new Date()"
                                    class="rounded-md border"
                                />
                            </div>

                            <div class="space-y-2">
                                <label class="text-sm font-medium">Time</label>
                                <Select v-model="formData.time">
                                    <SelectTrigger>
                                        <SelectValue placeholder="Select time" />
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectItem
                                            v-for="time in TIME_SLOTS"
                                            :key="time"
                                            :value="time"
                                        >
                                            {{ time }}
                                        </SelectItem>
                                    </SelectContent>
                                </Select>
                            </div>
                        </div>

                        <div class="flex justify-end gap-3 pt-4">
                            <Button variant="outline" @click="open = false"> Cancel </Button>
                            <Button @click="schedule"> Schedule </Button>
                        </div>
                    </div>
                </DialogContent>
            </Dialog>
        </div>

        <div class="spacey-4" v-if="interviews.length > 0">
            <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
                <MeetingCard
                    v-for="interview in interviews"
                    :key="interview.id"
                    :interview="interview"
                />
            </div>
        </div>
        <div class="text-center py-12 text-muted-foreground" v-else>No interviews scheduled</div>
    </div>
</template>

<script setup>
import MeetingCard from '@/components/MeetingCard.vue';
import Button from '@/components/ui/button/Button.vue';
import Calendar from '@/components/ui/calendar/Calendar.vue';
import Dialog from '@/components/ui/dialog/Dialog.vue';
import DialogContent from '@/components/ui/dialog/DialogContent.vue';
import DialogHeader from '@/components/ui/dialog/DialogHeader.vue';
import DialogTitle from '@/components/ui/dialog/DialogTitle.vue';
import DialogTrigger from '@/components/ui/dialog/DialogTrigger.vue';
import Input from '@/components/ui/input/Input.vue';
import { Select } from '@/components/ui/select';
import SelectContent from '@/components/ui/select/SelectContent.vue';
import SelectItem from '@/components/ui/select/SelectItem.vue';
import SelectTrigger from '@/components/ui/select/SelectTrigger.vue';
import SelectValue from '@/components/ui/select/SelectValue.vue';
import Textarea from '@/components/ui/textarea/Textarea.vue';
import UserInfo from '@/components/UserInfo.vue';
import router from '@/router';
import { useAuthStore } from '@/stores/auth';
import { getLocalTimeZone, today } from '@internationalized/date';
import axios from 'axios';
import Cookies from 'js-cookie';
import { X } from 'lucide-vue-next';
import { computed, reactive, ref } from 'vue';
import { toast } from 'vue3-toastify';

const authUser = useAuthStore().authUser;
if (authUser.role !== 'interviewer') {
    router.push({ name: 'home' });
}

const TIME_SLOTS = [
    '09:00',
    '09:30',
    '10:00',
    '10:30',
    '11:00',
    '11:30',
    '12:00',
    '12:30',
    '13:00',
    '13:30',
    '14:00',
    '14:30',
    '15:00',
    '15:30',
    '16:00',
    '16:30',
    '17:00',
];

const open = ref(false);
const candidates = ref([]);
const interviewers = ref([]);
const interviews = ref([]);

const getInterviews = () => {
    axios
        .get('/interviews', {
            headers: {
                Authorization: `Bearer ${Cookies.get('jwt')}`,
            },
        })
        .then((resp) => {
            interviews.value = resp.data.interviews;
        })
        .catch((err) => toast.error(err));
};

const getUsers = () => {
    axios
        .get('/users', {
            headers: {
                Authorization: `Bearer ${Cookies.get('jwt')}`,
            },
        })
        .then((resp) => {
            const users = resp.data.users;
            candidates.value = users.filter((u) => u.role === 'candidate');
            interviewers.value = users.filter((u) => u.role === 'interviewer');
        })
        .catch((err) => toast.error(err));
};

getInterviews();
getUsers();

const selectedInterviewers = computed(() => {
    return interviewers.value.filter((i) => formData.interviewerUUIDs.includes(i.uuid));
});

const availableInterviewers = computed(() => {
    return interviewers.value.filter((i) => !formData.interviewerUUIDs.includes(i.uuid));
});

const formData = reactive({
    title: '',
    desription: '',
    date: today(getLocalTimeZone()),
    time: '12:00',
    candidateUUID: '',
    interviewerUUIDs: [authUser.uuid],
});

const addInterviewer = (interviewerUUID) => {
    formData.interviewerUUIDs.push(interviewerUUID);
};

const removeInterviewer = (interviewerUUID) => {
    if (interviewerUUID === authUser.uuid) {
        toast.error("You can't remove yourself!");
        return;
    }

    formData.interviewerUUIDs = formData.interviewerUUIDs.filter(
        (uuid) => uuid !== interviewerUUID,
    );
};

const schedule = () => {
    if (!formData.candidateUUID || formData.interviewerUUIDs.length === 0) {
        toast.error('Please select both candidate and at least one interviewer');
        return;
    }

    if (!formData.title || !formData.desription) {
        toast.error('Please write a description and a title');
        return;
    }

    const [hours, minutes] = formData.time.split(':');
    const startTime = new Date(formData.date.toDate());
    startTime.setHours(parseInt(hours), parseInt(minutes), 0);

    if (startTime < new Date()) {
        toast.error(
            'The interview cannot be scheduled in the past. Please select a future date and time.',
        );
        return;
    }

    const data = {
        title: formData.title,
        description: formData.desription,
        attendeeUUIDs: JSON.stringify([formData.candidateUUID, ...formData.interviewerUUIDs]),
        startTime: startTime.toUTCString(),
    };

    axios
        .post('/interviews/store', data, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                Authorization: `Bearer ${Cookies.get('jwt')}`,
            },
        })
        .then(() => {
            open.value = false;
            Object.assign(formData, {
                title: '',
                desription: '',
                date: today(getLocalTimeZone()),
                time: '12:00',
                candidateUUID: '',
                interviewerUUIDs: [authUser.uuid],
            });
            getInterviews();
            toast.success('Meeting scheduled successfully!');
        })
        .catch((err) => console.log(err));
};
</script>
