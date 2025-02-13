<template>
    <ResizablePanelGroup direction="vertical" class="min-h-[calc-100vh-4rem-1px]">
        <ResizablePanel>
            <ScrollArea class="h-full">
                <div class="p-6">
                    <div class="max-w-4xl mx-auto space-y-6">
                        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
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
                                <Select v-model="selectedQuestion.titleSlug">
                                    <SelectTrigger class="w-[180px]">
                                        <SelectValue placeholder="Select question" />
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectItem v-for="q in questions" :key="q.titleSlug" :value="q.titleSlug">
                                            {{ q.title }}
                                        </SelectItem>
                                    </SelectContent>
                                </Select>

                                <Select v-model="language">
                                    <SelectTrigger class="w-[150px]">
                                        <SelectValue>
                                            <SelectValue placeholder="Select a language" />
                                        </SelectValue>
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectItem v-for="lang in supportedLanguages" :key="lang[0]" :value="lang[0]">
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
                <VueMonacoEditor height="100%" :default-language="language" :language="language" theme="vs-dark"
                    v-model:value="code" :options="{
                        minimap: { enabled: false },
                        fontSize: 18,
                        lineNumbers: 'on',
                        scrollBeyondLastLine: true,
                        automaticLayout: true,
                        padding: { top: 16, bottom: 16 },
                        wordWrap: 'on',
                        wrappingIndent: 'indent',
                    }">
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
import { ref } from 'vue';
import ResizableHandle from './ui/resizable/ResizableHandle.vue';
import SelectContent from './ui/select/SelectContent.vue';
import SelectItem from './ui/select/SelectItem.vue';
import ScrollBar from './ui/scroll-area/ScrollBar.vue';
import ScrollArea from './ui/scroll-area/ScrollArea.vue';
import { VueMonacoEditor } from '@guolao/vue-monaco-editor';

// Placeholder for dynamic data fetching from an API
const questions = [
    {
        titleSlug: 'two-sum',
        title: 'Two Sum',
        questionHTML:
            '<p>Given an array of integers <code>nums</code>&nbsp;and an integer <code>target</code>, return <em>indices of the two numbers such that they add up to <code>target</code></em>.</p>\n\n<p>You may assume that each input would have <strong><em>exactly</em> one solution</strong>, and you may not use the <em>same</em> element twice.</p>\n\n<p>You can return the answer in any order.</p>\n\n<p>&nbsp;</p>\n<p><strong class="example">Example 1:</strong></p>\n\n<pre>\n<strong>Input:</strong> nums = [2,7,11,15], target = 9\n<strong>Output:</strong> [0,1]\n<strong>Explanation:</strong> Because nums[0] + nums[1] == 9, we return [0, 1].\n</pre>\n\n<p><strong class="example">Example 2:</strong></p>\n\n<pre>\n<strong>Input:</strong> nums = [3,2,4], target = 6\n<strong>Output:</strong> [1,2]\n</pre>\n\n<p><strong class="example">Example 3:</strong></p>\n\n<pre>\n<strong>Input:</strong> nums = [3,3], target = 6\n<strong>Output:</strong> [0,1]\n</pre>\n\n<p>&nbsp;</p>\n<p><strong>Constraints:</strong></p>\n\n<ul>\n\t<li><code>2 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>\n\t<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>\n\t<li><code>-10<sup>9</sup> &lt;= target &lt;= 10<sup>9</sup></code></li>\n\t<li><strong>Only one valid answer exists.</strong></li>\n</ul>\n\n<p>&nbsp;</p>\n<strong>Follow-up:&nbsp;</strong>Can you come up with an algorithm that is less than <code>O(n<sup>2</sup>)</code><font face="monospace">&nbsp;</font>time complexity?',
    },
];

const supportedLanguages = [['go', 'Go'], ['python', 'Python'], ['cpp', 'C++'], ['javascript', 'Javascript']];

const selectedQuestion = ref(questions[0]);
const language = ref(supportedLanguages[0][0]);
const code = ref('');
</script>
