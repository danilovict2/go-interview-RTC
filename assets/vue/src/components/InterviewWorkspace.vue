<template>
    <AppLoading v-if="isLoading" />
    <ResizablePanelGroup direction="vertical" class="min-h-[calc-100vh-4rem-1px]" v-else>
        <ResizablePanel>
            <ScrollArea class="h-full">
                <div class="p-6">
                    <div class="max-w-4xl mx-auto space-y-6">
                        <div
                            class="flex flex-col sm:flex-row sm:items-center justify-between gap-4"
                        >
                            <div class="space-y-1">
                                <div class="flex items-center gap-2">
                                    <h2 class="text-2xl font-semibold tracking-tight">
                                        {{ selectedQuestion.title }}
                                    </h2>
                                    <p class="text-sm text-muted-foreground">
                                        Choose your language and solve the problem
                                    </p>
                                </div>
                            </div>
                            <div class="flex items-center gap-3">
                                <Select v-model="language">
                                    <SelectTrigger class="w-[150px]">
                                        <SelectValue>
                                            <SelectValue placeholder="Select a language" />
                                        </SelectValue>
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectItem
                                            v-for="lang in supportedLanguages"
                                            :key="lang[0]"
                                            :value="lang[0]"
                                        >
                                            <div class="flex items-center gap-2">
                                                {{ lang[1] }}
                                            </div>
                                        </SelectItem>
                                    </SelectContent>
                                </Select>
                            </div>
                        </div>
                        <span v-html="selectedQuestion.questionHTML"></span>
                    </div>
                </div>
                <ScrollBar />
            </ScrollArea>
        </ResizablePanel>

        <ResizableHandle withHandle />

        <ResizablePanel :default-size="60" :max-size="100">
            <div class="h-full relative">
                <VueMonacoEditor
                    height="100%"
                    :default-language="language"
                    :language="language"
                    theme="vs-dark"
                    v-model:value="code"
                    :options="{
                        minimap: { enabled: false },
                        fontSize: 18,
                        lineNumbers: 'on',
                        scrollBeyondLastLine: true,
                        automaticLayout: true,
                        padding: { top: 16, bottom: 16 },
                        wordWrap: 'on',
                        wrappingIndent: 'indent',
                    }"
                >
                </VueMonacoEditor>
            </div>
        </ResizablePanel>
    </ResizablePanelGroup>
</template>

<script setup>
import { ResizablePanel } from './ui/resizable';
import ResizablePanelGroup from './ui/resizable/ResizablePanelGroup.vue';
import Select from './ui/select/Select.vue';
import SelectTrigger from './ui/select/SelectTrigger.vue';
import SelectValue from './ui/select/SelectValue.vue';
import { onUnmounted, ref, watch } from 'vue';
import ResizableHandle from './ui/resizable/ResizableHandle.vue';
import SelectContent from './ui/select/SelectContent.vue';
import SelectItem from './ui/select/SelectItem.vue';
import ScrollBar from './ui/scroll-area/ScrollBar.vue';
import ScrollArea from './ui/scroll-area/ScrollArea.vue';
import { VueMonacoEditor } from '@guolao/vue-monaco-editor';
import axios from 'axios';
import { toast } from 'vue3-toastify';
import AppLoading from './AppLoading.vue';
import { useAuthStore } from '@/stores/auth';
import { Peer } from 'peerjs';

const { call } = defineProps({
    call: Object,
});

const supportedLanguages = [
    ['go', 'Go'],
    ['python', 'Python'],
    ['cpp', 'C++'],
    ['javascript', 'Javascript'],
];

const isLoading = ref(true);
const authUser = useAuthStore().authUser;
const selectedQuestion = ref({});
const language = ref(supportedLanguages[0][0]);
const code = ref('');

const peer = new Peer(authUser.uuid);
let connections = [];

const subscription = call.state.participants$.subscribe((p) => {
    connections = [];

    for (const participant of p) {
        if (participant.userId === authUser.uuid) {
            continue;
        }

        const conn = peer.connect(participant.userId);
        connections.push(conn);
    }
});

peer.on('connection', (conn) => {
    conn.on('data', (data) => {
        code.value = data;
    });
});

watch(code, (newCode) => {
    // Broadcast the updated code to all participants except the sender
    for (const conn of connections) {
        conn.send(newCode);
    }
});

onUnmounted(() => {
    if (subscription) {
        subscription.unsubscribe();
    }
});

axios
    .get(import.meta.env.VITE_LEET_CODE_API_URL + '/dailyQuestion')
    .then((resp) => {
        selectedQuestion.value = {
            title: resp.data.data.activeDailyCodingChallengeQuestion.question.title,
            questionHTML: resp.data.data.activeDailyCodingChallengeQuestion.question.content,
        };
    })
    .catch((error) => {
        console.error('Error loading the problem:', error);
        toast.error('An error occurred while loading the problem');
    })
    .finally(() => (isLoading.value = false));
</script>
