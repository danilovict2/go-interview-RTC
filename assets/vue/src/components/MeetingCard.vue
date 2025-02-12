<template>
    <Card>
        <CardHeader class="space-y-2">
            <div class="flex items-center justify-between">
                <div class="flex items-center gap-2 text-sm text-muted-foreground">
                    <CalendarIcon class="h-4 w-4" />
                    {{ interview.startDate }}
                </div>

                <Badge
                    :variant="
                        status === 'live'
                            ? 'default'
                            : status === 'upcoming'
                              ? 'secondary'
                              : 'outline'
                    "
                >
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

            <CardDescription v-show="interview.description !== ''" class="line-clamp-2"
                >{{ interview.description }}
            </CardDescription>
        </CardHeader>

        <CardContent>
            <Button v-show="status === 'live'" class="w-full" @click="console.log('join meeting')">
                Join Meeting
            </Button>

            <Button variant="outline" class="w-full" disabled v-show="status === 'upcoming'">
                Waiting to Start
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

const status = 'live';

defineProps({
    interview: Object,
});
</script>
