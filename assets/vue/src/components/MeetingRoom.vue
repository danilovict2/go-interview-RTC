<template>
    <div class="h-[calc(100vh-4rem-1px)]">
        <ResizablePanelGroup direction="horizontal">
            <ResizablePanel :default-size="35" :min-size="25" :max-size="100" class="relative">
                <div class="absolute inset-0">
                    <PaginatedGridLayout
                        v-if="layout === 'grid'"
                        :call="call"
                        :participants="participants"
                    />
                    <SpeakerLayout :call="call" :participants="participants" v-else />

                    <div
                        v-show="showParticipants"
                        class="absolute right-0 top-0 h-full w-[300px] bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 border border-border"
                    >
                        <CallParticipantsList
                            :participants="participants"
                            @close="showParticipants = false"
                        />
                    </div>
                </div>

                <div class="absolute bottom-4 left-0 right-0">
                    <div class="flex flex-col items-center gap-4">
                        <div class="flex items-center gap-2 flex-wrap justify-center px-4">
                            <CallControlls
                                :call="call"
                                @leave="
                                    call.leave();
                                    router.push({ name: 'Home' });
                                "
                            />

                            <div class="flex items-center gap-2">
                                <DropdownMenu>
                                    <DropdownMenuTrigger asChild>
                                        <Button variant="outline" size="icon" class="size-10">
                                            <LayoutList class="size-4" />
                                        </Button>
                                    </DropdownMenuTrigger>
                                    <DropdownMenuContent>
                                        <DropdownMenuItem @click="layout = 'grid'">
                                            Grid View
                                        </DropdownMenuItem>
                                        <DropdownMenuItem @click="layout = 'speaker'">
                                            Speaker View
                                        </DropdownMenuItem>
                                    </DropdownMenuContent>
                                </DropdownMenu>

                                <Button
                                    variant="outline"
                                    size="icon"
                                    class="size-10"
                                    @click="showParticipants = !showParticipants"
                                >
                                    <Users class="size-4" />
                                </Button>
                                <EndCallButton
                                    :call="call"
                                    v-if="authUser.role === 'interviewer'"
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </ResizablePanel>

            <ResizableHandle with-handle />

            <ResizablePanel :default-size="65" :min-size="25">
                <InterviewWorkspace />
            </ResizablePanel>
        </ResizablePanelGroup>
    </div>
</template>

<script setup>
import { onUnmounted, ref } from 'vue';
import { ResizablePanel } from './ui/resizable';
import ResizablePanelGroup from './ui/resizable/ResizablePanelGroup.vue';
import { DropdownMenu, DropdownMenuContent } from './ui/dropdown-menu';
import { DropdownMenuTrigger } from 'radix-vue';
import Button from './ui/button/Button.vue';
import DropdownMenuItem from './ui/dropdown-menu/DropdownMenuItem.vue';
import CallParticipantsList from './CallParticipantsList.vue';
import CallControlls from './CallControlls.vue';
import router from '@/router';
import EndCallButton from './EndCallButton.vue';
import PaginatedGridLayout from './PaginatedGridLayout.vue';
import SpeakerLayout from './SpeakerLayout.vue';
import { LayoutList, Users } from 'lucide-vue-next';
import ResizableHandle from './ui/resizable/ResizableHandle.vue';
import { publishingVideo, speakerLayoutSortPreset } from '@stream-io/video-client';
import InterviewWorkspace from './InterviewWorkspace.vue';
import { useAuthStore } from '@/stores/auth';

const { call } = defineProps({
    call: Object,
});

const showParticipants = ref(false);
const layout = ref('grid');
const participants = ref([]);
const authUser = useAuthStore().authUser;

const subscription = call.state.participants$.subscribe((p) => {
    if (layout.value === 'grid') {
        participants.value = p.sort(publishingVideo);
    } else {
        participants.value = p.sort(speakerLayoutSortPreset);
    }
});

onUnmounted(() => {
    if (subscription) {
        subscription.unsubscribe();
    }
});
</script>
