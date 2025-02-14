<template>
    <Card class="group hover:shadow-md transition-all">
        <CardHeader class="space-y-1">
            <div class="space-y-2">
                <div class="flex flex-col gap-1.5">
                    <div class="flex items-center text-sm text-muted-foreground gap-2">
                        <Calendar class="h-3.5 w-3.5" />
                        <span>{{ formattedStartTime }}</span>
                    </div>
                    <div class="flex items-center text-sm text-muted-foreground gap-2">
                        <Clock class="h-3.5 w-3.5" />
                        <span>{{ duration }}</span>
                    </div>
                </div>
            </div>
        </CardHeader>

        <CardContent>
            <div class="w-full aspect-video bg-muted/50 rounded-lg flex items-center justify-center cursor-pointer group"
                @click="playRecording">
                <div
                    class="size-12 rounded-full bg-background/90 flex items-center justify-center group-hover:bg-primary transition-colors">
                    <Play class="size-6 text-muted-foreground group-hover:text-primary-foreground transition-colors" />
                </div>
            </div>
        </CardContent>
        <CardFooter class="gap-2">
            <Button class="flex-1" @click="playRecording">
                <PlayIcon class="size-4 mr-2" />
                Play Recording
            </Button>
            <Button variant="secondary" @click="copyLink">
                <CopyIcon class="size-4" />
            </Button>
        </CardFooter>
    </Card>
</template>

<script setup>
import { format, intervalToDuration } from 'date-fns';
import Card from './ui/card/Card.vue';
import CardHeader from './ui/card/CardHeader.vue';
import { CardContent } from './ui/card';
import { Calendar, Clock, Play } from 'lucide-vue-next';
import CardFooter from './ui/card/CardFooter.vue';
import { toast } from 'vue3-toastify';


const { recording } = defineProps({
    recording: Object,
});

const calculateRecordingDuration = (startTime, endTime) => {
    const start = new Date(startTime);
    const end = new Date(endTime);

    const duration = intervalToDuration({ start, end });

    if (duration.hours && duration.hours > 0) {
        return `${duration.hours}:${String(duration.minutes).padStart(2, "0")}:${String(
            duration.seconds
        ).padStart(2, "0")}`;
    }

    if (duration.minutes && duration.minutes > 0) {
        return `${duration.minutes}:${String(duration.seconds).padStart(2, "0")}`;
    }

    return `${duration.seconds} seconds`;
};

const duration =
    recording.start_time && recording.end_time
        ? calculateRecordingDuration(recording.start_time, recording.end_time)
        : "Unknown duration";

const formattedStartTime = recording.start_time ? format(new Date(recording.start_time), "MMM d, yyyy, hh:mm a") : "Unknown"

const copyLink = async () => {
    try {
        await navigator.clipboard.writeText(recording.url);
        toast.success("Recording link copied to clipboard");
    } catch (error) {
        console.log(error);
        toast.error("Failed to copy link to clipboard");
    }
}

const playRecording = () => {
    window.open(recording.url, '_blank')
}
</script>