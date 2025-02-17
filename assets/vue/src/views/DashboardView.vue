<template>
    <div class="container mx-auto py-10">
        <div class="flex items-center mb-8">
            <RouterLink to="/schedule">
                <Button>Schedule New Interview</Button>
            </RouterLink>
        </div>

        <div class="space-y-8">
            <section v-for="category in categories" :key="category.id">
                <div class="flex items-center gap-2 mb-4">
                    <h2 class="text-xl font-semibold">{{ category.title }}</h2>
                    <Badge :variant="category.variant">{{
                        groupedInterviews[category.id].length
                    }}</Badge>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    <Card
                        class="hover:shadow-md transition-all"
                        v-for="interview in groupedInterviews[category.id]"
                        :key="interview.id"
                    >
                        <CardHeader class="p-4">
                            <div class="flex items-center gap-3">
                                <Avatar class="h-10 w-10">
                                    <AvatarImage
                                        :src="`https://api.dicebear.com/9.x/initials/svg/seed=${findCandidate(interview).first_name}-${findCandidate(interview).last_name}`"
                                    />
                                    <AvatarFallback>
                                        <UserCircle class="h-8 w-8" />
                                    </AvatarFallback>
                                </Avatar>
                                <div>
                                    <CardTitle class="text-base">{{
                                        findCandidate(interview).first_name +
                                        ' ' +
                                        findCandidate(interview).last_name
                                    }}</CardTitle>
                                    <p class="text-sm text-muted-foreground">
                                        {{ interview.title }}
                                    </p>
                                </div>
                            </div>
                        </CardHeader>

                        <CardContent class="p-4">
                            <div class="flex items-center gap-4 text-sm text-muted-foreground">
                                <div class="flex items-center gap-1">
                                    <CalendarIcon class="h-4 w-4" />
                                    {{ format(interview.start_time, 'MMM dd') }}
                                </div>
                                <div class="flex items-center gap-1">
                                    <Clock class="h-4 w-4" />
                                    {{ format(interview.start_time, 'hh:mm a') }}
                                </div>
                            </div>
                        </CardContent>

                        <CardFooter class="p-4 pt-0 flex flex-col gap-3">
                            <div
                                class="flex gap-2 w-full"
                                v-if="interview.decision === 'undecided'"
                            >
                                <Button
                                    variant="success"
                                    class="flex-1"
                                    @click="updateDecision(interview, 'pass')"
                                >
                                    <CheckCircle2 class="h-4 w-4 mr-2" />
                                    Pass
                                </Button>
                                <Button
                                    variant="destructive"
                                    class="flex-1"
                                    @click="updateDecision(interview, 'fail')"
                                >
                                    <XCircle class="h-4 w-4 mr-2" />
                                    Fail
                                </Button>
                            </div>
                            <CommentDialog :interviewID="interview.stream_call_id" />
                        </CardFooter>
                    </Card>
                </div>
            </section>
        </div>
    </div>
</template>

<script setup>
import CommentDialog from '@/components/CommentDialog.vue';
import Avatar from '@/components/ui/avatar/Avatar.vue';
import AvatarFallback from '@/components/ui/avatar/AvatarFallback.vue';
import AvatarImage from '@/components/ui/avatar/AvatarImage.vue';
import Badge from '@/components/ui/badge/Badge.vue';
import Button from '@/components/ui/button/Button.vue';
import Card from '@/components/ui/card/Card.vue';
import CardContent from '@/components/ui/card/CardContent.vue';
import CardFooter from '@/components/ui/card/CardFooter.vue';
import CardHeader from '@/components/ui/card/CardHeader.vue';
import CardTitle from '@/components/ui/card/CardTitle.vue';
import { useInterview } from '@/composables/interview';
import router from '@/router';
import { useAuthStore } from '@/stores/auth';
import axios from 'axios';
import { format } from 'date-fns';
import Cookies from 'js-cookie';
import { CalendarIcon, CheckCircle2, Clock, UserCircle, XCircle } from 'lucide-vue-next';
import { computed, ref } from 'vue';
import { RouterLink } from 'vue-router';
import { toast } from 'vue3-toastify';

const authUser = useAuthStore().authUser;
if (authUser.role !== 'interviewer') {
    router.push({ name: 'Home' });
}

const interviews = ref([]);
useInterview().getInterviews(interviews);

const categories = computed(() => {
    return [
        { id: 'pass', title: 'Pass', variant: 'default' },
        { id: 'fail', title: 'Fail', variant: 'destructive' },
        { id: 'undecided', title: 'Undecided', variant: 'outline' },
    ];
});

const groupedInterviews = computed(() => {
    const initial = {
        pass: [],
        fail: [],
        undecided: [],
    };

    if (interviews.value.length === 0) return initial;

    return interviews.value.reduce((acc, i) => {
        switch (i.decision) {
            case 'pass':
                acc.pass = [...(acc.pass || []), i];
                break;
            case 'fail':
                acc.fail = [...(acc.fail || []), i];
                break;
            default:
                acc.undecided = [...(acc.undecided || []), i];
                break;
        }

        return acc;
    }, initial);
});

const findCandidate = (interview) => {
    return interview.attendees.find((at) => at.role === 'candidate');
};

const updateDecision = (interview, decision) => {
    axios
        .patch(
            `/interviews/${interview.stream_call_id}/change_decision`,
            {
                decision: decision,
            },
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    Authorization: `Bearer ${Cookies.get('jwt')}`,
                },
            },
        )
        .then(() => {
            useInterview().getInterviews(interviews);
        })
        .catch((err) => toast.error(err.response.data.message));
};
</script>
