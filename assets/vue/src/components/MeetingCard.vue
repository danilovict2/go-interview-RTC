<template>
    <Card>
        <CardHeader class="space-y-2">
            <div class="flex items-center justify-between">
                <div class="flex items-center gap-2 text-sm text-muted-foreground">
                    <CalendarIcon class="h-4 w-4" />
                    {{ formattedStartTime }}
                </div>

                <Badge
                    :variant="
                        interview.status === 'live'
                            ? 'default'
                            : interview.status === 'upcoming'
                              ? 'secondary'
                              : 'outline'
                    "
                >
                    {{
                        interview.status === 'live'
                            ? 'Live Now'
                            : interview.status === 'upcoming'
                              ? 'Upcoming'
                              : 'Completed'
                    }}
                </Badge>
            </div>

            <CardTitle>{{ interview.title }}</CardTitle>

            <CardDescription v-show="interview.description !== ''" class="line-clamp-2"
                >{{ interview.description }}
            </CardDescription>
        </CardHeader>

        <CardContent>
            <Button v-if="interview.status === 'live'" class="w-full" @click="console.log('join meeting')">
                Join Meeting
            </Button>

            <Button v-else-if="interview.status === 'upcoming'" variant="outline" class="w-full" disabled> Waiting to Start </Button>
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


const { interview } = defineProps({
    interview: Object,
});

const formattedStartTime = format(new Date(interview.start_time), 'MMM d, yyyy, hh:mm a');
</script>
