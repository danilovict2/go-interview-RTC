<template>
    <AppLoading v-if="isLoading"></AppLoading>
    <div class="container max-w-7xl mx-auto p-6" v-else>
        <div class="rounded-lg bg-card p-6 border shadow-sm mb-10">
            <h1
                class="text-4xl font-bold bg-gradient-to-r from-emerald-600 to-teal-500 bg-clip-text text-transparent"
            >
                Welcome back!
            </h1>
            <p class="text-muted-foreground mt-2">
                {{
                    isInterviewer
                        ? 'Manage your interviews and review candidates effectively'
                        : 'Access your upcoming interviews and preparations'
                }}
            </p>
        </div>

        <div v-if="isInterviewer">
            <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
                <ActionCard
                    v-for="(action, index) in actions"
                    :key="index"
                    :action="action"
                    @click="handleAction(action.title)"
                ></ActionCard>
            </div>

            <MeetingModal
                :isOpen="showModal"
                @close="showModal = false"
                :title="modalType === 'join' ? 'Join Meeting' : 'Start Meeting'"
                :isJoinMeeting="modalType === 'join'"
            />
        </div>
        <div v-else>
            <div>
                <h1 class="text-3xl font-bold">Your Interviews</h1>
                <p class="text-muted-foreground mt-1">View and join your scheduled interviews</p>
            </div>

            <div class="mt-8">
                <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3" v-if="interviews.length > 0">
                    <MeetingCard
                        v-for="interview in interviews"
                        :key="interview.id"
                        :interview="interview"
                    />
                </div>

                <div class="text-center py-12 text-muted-foreground" v-else>
                    You have no scheduled interviews at the moment
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import ActionCard from '@/components/ActionCard.vue';
import AppLoading from '@/components/AppLoading.vue';
import MeetingCard from '@/components/MeetingCard.vue';
import MeetingModal from '@/components/MeetingModal.vue';
import { useInterview } from '@/composables/interview';
import router from '@/router';
import { useAuthStore } from '@/stores/auth';
import { Calendar, Clock, Code2, Users } from 'lucide-vue-next';
import { computed, ref } from 'vue';

const interviews = ref([]);
const isLoading = ref(false);
useInterview().getInterviews(interviews, isLoading);

const authStore = useAuthStore();
const isInterviewer = computed(() => {
    return authStore.authUser?.role === 'interviewer';
});

const showModal = ref(false);
const modalType = ref('');

const actions = [
    {
        icon: Code2,
        title: 'New Call',
        description: 'Start an instant call',
        color: 'primary',
        gradient: 'from-primary/10 via-primary/5 to-transparent',
    },
    {
        icon: Users,
        title: 'Join Interview',
        description: 'Enter via invitation link',
        color: 'purple-500',
        gradient: 'from-purple-500/10 via-purple-500/5 to-transparent',
    },
    {
        icon: Calendar,
        title: 'Schedule',
        description: 'Plan upcoming interviews',
        color: 'blue-500',
        gradient: 'from-blue-500/10 via-blue-500/5 to-transparent',
    },
    {
        icon: Clock,
        title: 'Recordings',
        description: 'Access past interviews',
        color: 'orange-500',
        gradient: 'from-orange-500/10 via-orange-500/5 to-transparent',
    },
];

const handleAction = (actionTitle) => {
    switch (actionTitle) {
        case 'New Call':
            modalType.value = 'start';
            showModal.value = true;
            break;
        case 'Join Interview':
            modalType.value = 'join';
            showModal.value = true;
            break;
        default:
            router.push(actionTitle.toLowerCase());
            break;
    }
};
</script>
